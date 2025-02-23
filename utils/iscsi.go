// Copyright 2022 NetApp, Inc. All Rights Reserved.

package utils

import (
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"

	. "github.com/netapp/trident/logger"
)

const (
	iSCSIErrNoObjsFound                 = 21
	iSCSIDeviceDiscoveryTimeoutSecs     = 90
	multipathDeviceDiscoveryTimeoutSecs = 90
	temporaryMountDir                   = "/tmp_mnt"
	volumeMountDir                      = "/vol_mnt"
	unknownFstype                       = "<unknown>"
)

// AttachISCSIVolume attaches the volume to the local host.  This method must be able to accomplish its task using only the data passed in.
// It may be assumed that this method always runs on the host to which the volume will be attached.  If the mountpoint
// parameter is specified, the volume will be mounted.  The device path is set on the in-out publishInfo parameter
// so that it may be mounted later instead.
func AttachISCSIVolume(ctx context.Context, name, mountpoint string, publishInfo *VolumePublishInfo) error {
	Logc(ctx).Debug(">>>> iscsi.AttachISCSIVolume")
	defer Logc(ctx).Debug("<<<< iscsi.AttachISCSIVolume")

	var err error
	lunID := int(publishInfo.IscsiLunNumber)

	var bkportal []string
	var portalIps []string
	bkportal = append(bkportal, ensureHostportFormatted(publishInfo.IscsiTargetPortal))
	portalIps = append(portalIps, getHostportIP(publishInfo.IscsiTargetPortal))

	for _, p := range publishInfo.IscsiPortals {
		bkportal = append(bkportal, ensureHostportFormatted(p))
		portalIps = append(portalIps, getHostportIP(p))
	}

	advertisedPortalCount := 1 + len(publishInfo.IscsiPortals)

	if publishInfo.IscsiInterface == "" {
		publishInfo.IscsiInterface = "default"
	}

	Logc(ctx).WithFields(log.Fields{
		"volume":         name,
		"mountpoint":     mountpoint,
		"lunID":          lunID,
		"targetPortals":  bkportal,
		"targetIQN":      publishInfo.IscsiTargetIQN,
		"iscsiInterface": publishInfo.IscsiInterface,
		"fstype":         publishInfo.FilesystemType,
	}).Debug("Attaching iSCSI volume.")

	if !ISCSISupported(ctx) {
		err := errors.New("unable to attach: open-iscsi tools not found on host")
		Logc(ctx).Errorf("Unable to attach volume: open-iscsi utils not found")
		return err
	}

	// Ensure we are logged into correct portals
	if publishInfo.UseCHAP {
		bkPortalsToLogin, loggedIn, err := portalsToLogin(ctx, publishInfo.IscsiTargetIQN, bkportal)
		if err != nil {
			return err
		}

		loggedIn = logInToPortals(ctx, bkPortalsToLogin, publishInfo, loggedIn)

		if !loggedIn {
			return AuthError("iSCSI login failed using CHAP")
		}
	} else {
		portalIpsToLogin, loggedIn, err := portalsIpsToLogin(ctx, publishInfo.IscsiTargetIQN, portalIps)
		if err != nil {
			return err
		}

		newLogin := EnsureISCSISessions(ctx, publishInfo.IscsiTargetIQN, publishInfo.IscsiInterface, portalIpsToLogin)

		if !loggedIn && !newLogin {
			return fmt.Errorf("iSCSI login failed")
		}
	}

	// First attempt to fix invalid serials by rescanning them
	err = handleInvalidSerials(ctx, lunID, publishInfo.IscsiTargetIQN, publishInfo.IscsiLunSerial, rescanOneLun)
	if err != nil {
		return err
	}

	// Then attempt to fix invalid serials by purging them (to be scanned
	// again later)
	err = handleInvalidSerials(ctx, lunID, publishInfo.IscsiTargetIQN, publishInfo.IscsiLunSerial, purgeOneLun)
	if err != nil {
		return err
	}

	// If LUN isn't present, scan the target and wait for the device(s) to appear
	// if not attached need to scan
	shouldScan := !IsAlreadyAttached(ctx, lunID, publishInfo.IscsiTargetIQN)
	err = waitForDeviceScanIfNeeded(ctx, lunID, publishInfo.IscsiTargetIQN, shouldScan)
	if err != nil {
		Logc(ctx).Errorf("Could not find iSCSI device: %+v", err)
		return err
	}

	// At this point if the serials are still invalid, give up so the
	// caller can retry (invoking the remediation steps above in the
	// process, if they haven't already been run).
	failHandler := func(ctx context.Context, path string) error {
		Logc(ctx).Error("Detected LUN serial number mismatch, attaching volume would risk data corruption, giving up")
		return fmt.Errorf("LUN serial number mismatch, kernel has stale cached data")
	}
	err = handleInvalidSerials(ctx, lunID, publishInfo.IscsiTargetIQN, publishInfo.IscsiLunSerial, failHandler)
	if err != nil {
		return err
	}

	err = waitForMultipathDeviceForLUN(ctx, lunID, advertisedPortalCount, publishInfo.IscsiTargetIQN)
	if err != nil {
		return err
	}

	// Lookup all the SCSI device information, and include filesystem type only if not raw block volume
	needFSType := publishInfo.FilesystemType != fsRaw

	deviceInfo, err := getDeviceInfoForLUN(ctx, lunID, publishInfo.IscsiTargetIQN, needFSType, false)
	if err != nil {
		return fmt.Errorf("error getting iSCSI device information: %v", err)
	} else if deviceInfo == nil {
		return fmt.Errorf("could not get iSCSI device information for LUN %d", lunID)
	}

	Logc(ctx).WithFields(log.Fields{
		"scsiLun":         deviceInfo.LUN,
		"multipathDevice": deviceInfo.MultipathDevice,
		"devices":         deviceInfo.Devices,
		"fsType":          deviceInfo.Filesystem,
		"iqn":             deviceInfo.IQN,
	}).Debug("Found device.")

	// Make sure we use the proper device (multipath if in use)
	deviceToUse := deviceInfo.Devices[0]
	if deviceInfo.MultipathDevice != "" {
		deviceToUse = deviceInfo.MultipathDevice
	}
	if deviceToUse == "" {
		return fmt.Errorf("could not determine device to use for %v", name)
	}
	devicePath := "/dev/" + deviceToUse
	if err := waitForDevice(ctx, devicePath); err != nil {
		return fmt.Errorf("could not find device %v; %s", devicePath, err)
	}

	// Return the device in the publish info in case the mount will be done later
	publishInfo.DevicePath = devicePath

	if publishInfo.FilesystemType == fsRaw {
		return nil
	}

	existingFstype := deviceInfo.Filesystem
	if existingFstype == "" {
		Logc(ctx).WithFields(log.Fields{"volume": name, "fstype": publishInfo.FilesystemType}).Debug("Formatting LUN.")
		err := formatVolume(ctx, devicePath, publishInfo.FilesystemType)
		if err != nil {
			return fmt.Errorf("error formatting LUN %s, device %s: %v", name, deviceToUse, err)
		}
	} else if existingFstype != unknownFstype && existingFstype != publishInfo.FilesystemType {
		Logc(ctx).WithFields(log.Fields{
			"volume":          name,
			"existingFstype":  existingFstype,
			"requestedFstype": publishInfo.FilesystemType,
		}).Error("LUN already formatted with a different file system type.")
		return fmt.Errorf("LUN %s, device %s already formatted with other filesystem: %s",
			name, deviceToUse, existingFstype)
	} else {
		Logc(ctx).WithFields(log.Fields{
			"volume": name,
			"fstype": deviceInfo.Filesystem,
		}).Debug("LUN already formatted.")
	}

	// Attempt to resolve any filesystem inconsistencies that might be due to dirty node shutdowns, cloning
	// in-use volumes, or creating volumes from snapshots taken from in-use volumes.  This is only safe to do
	// if a device is not mounted.  The fsck command returns a non-zero exit code if filesystem errors are found,
	// even if they are completely and automatically fixed, so we don't return any error here.
	mounted, err := IsMounted(ctx, devicePath, "")
	if err != nil {
		return err
	}
	if !mounted {
		_ = repairVolume(ctx, devicePath, publishInfo.FilesystemType)
	}

	// Optionally mount the device
	if mountpoint != "" {
		if err := MountDevice(ctx, devicePath, mountpoint, publishInfo.MountOptions, false); err != nil {
			return fmt.Errorf("error mounting LUN %v, device %v, mountpoint %v; %s",
				name, deviceToUse, mountpoint, err)
		}
	}

	return nil
}

func logInToPortals(
	ctx context.Context, bkPortalsToLogin []string, publishInfo *VolumePublishInfo, loggedIn bool,
) bool {
	for _, portal := range bkPortalsToLogin {
		// Set scanning to manual
		// Swallow this error, someone is running an old version of Debian/Ubuntu
		_ = configureISCSITarget(ctx, publishInfo.IscsiTargetIQN, portal, "node.session.scan", "manual")

		err := loginWithChap(ctx, publishInfo.IscsiTargetIQN, portal, publishInfo.IscsiUsername,
			publishInfo.IscsiInitiatorSecret, publishInfo.IscsiTargetUsername, publishInfo.IscsiTargetSecret,
			publishInfo.IscsiInterface)
		if err != nil {
			Logc(ctx).WithFields(log.Fields{
				"err":    err,
				"portal": portal,
			}).Errorf("Failed to login to portal using CHAP.")

			continue
		}

		loggedIn = true
	}
	return loggedIn
}

// GetInitiatorIqns returns parsed contents of /etc/iscsi/initiatorname.iscsi
func GetInitiatorIqns(ctx context.Context) ([]string, error) {
	Logc(ctx).Debug(">>>> iscsi.GetInitiatorIqns")
	defer Logc(ctx).Debug("<<<< iscsi.GetInitiatorIqns")

	out, err := execCommand(ctx, "cat", "/etc/iscsi/initiatorname.iscsi")
	if err != nil {
		Logc(ctx).WithField("Error", err).Warn("Could not read initiatorname.iscsi; perhaps iSCSI is not installed?")
		return nil, err
	}

	return parseInitiatorIQNs(ctx, string(out)), nil
}

// parseInitiatorIQNs accepts the contents of /etc/iscsi/initiatorname.iscsi and returns the IQN(s).
func parseInitiatorIQNs(ctx context.Context, contents string) []string {
	iqns := make([]string, 0)

	lines := strings.Split(contents, "\n")
	for _, line := range lines {

		match := iqnRegex.FindStringSubmatch(line)

		if match == nil {
			continue
		}

		paramsMap := make(map[string]string)
		for i, name := range iqnRegex.SubexpNames() {
			if i > 0 && i <= len(match) {
				paramsMap[name] = match[i]
			}
		}

		if iqn, ok := paramsMap["iqn"]; ok {
			iqns = append(iqns, iqn)
		}
	}

	return iqns
}

// getSysfsBlockDirsForLUN returns the list of directories in sysfs where the block devices should appear
// after the scan is successful. One directory is returned for each path in the host session map.
func getSysfsBlockDirsForLUN(lunID int, hostSessionMap map[int]int) []string {
	paths := make([]string, 0)
	for hostNumber, sessionNumber := range hostSessionMap {
		p := fmt.Sprintf(
			chrootPathPrefix+"/sys/class/scsi_host/host%d/device/session%d/iscsi_session/session%d/device/target%d:0:0/%d:0:0:%d",
			hostNumber, sessionNumber, sessionNumber, hostNumber, hostNumber, lunID)
		paths = append(paths, p)
	}
	return paths
}

// getDevicesForLUN find the /dev/sd* device names for an iSCSI LUN.
func getDevicesForLUN(paths []string) ([]string, error) {
	devices := make([]string, 0)
	for _, p := range paths {
		dirname := p + "/block"
		if !PathExists(dirname) {
			continue
		}
		dirFd, err := os.Open(dirname)
		if err != nil {
			return nil, err
		}
		list, err := dirFd.Readdir(1)
		dirFd.Close()
		if err != nil {
			return nil, err
		}
		if 0 == len(list) {
			continue
		}
		devices = append(devices, list[0].Name())
	}
	return devices, nil
}

// waitForDeviceScanIfNeeded scans all paths to a specific LUN and waits until all
// SCSI disk-by-path devices for that LUN are present on the host.
func waitForDeviceScanIfNeeded(ctx context.Context, lunID int, iSCSINodeName string, shouldScan bool) error {
	fields := log.Fields{
		"lunID":         lunID,
		"iSCSINodeName": iSCSINodeName,
	}
	Logc(ctx).WithFields(fields).Debug(">>>> iscsi.waitForDeviceScanIfNeeded")
	defer Logc(ctx).WithFields(fields).Debug("<<<< iscsi.waitForDeviceScanIfNeeded")

	hostSessionMap := GetISCSIHostSessionMapForTarget(ctx, iSCSINodeName)
	if len(hostSessionMap) == 0 {
		return fmt.Errorf("no iSCSI hosts found for target %s", iSCSINodeName)
	}

	Logc(ctx).WithField("hostSessionMap", hostSessionMap).Debug("Built iSCSI host/session map.")
	hosts := make([]int, 0)
	for hostNumber := range hostSessionMap {
		hosts = append(hosts, hostNumber)
	}

	if shouldScan {
		if err := iSCSIScanTargetLUN(ctx, lunID, hosts); err != nil {
			Logc(ctx).WithField("scanError", err).Error("Could not scan for new LUN.")
		}
	}

	paths := getSysfsBlockDirsForLUN(lunID, hostSessionMap)
	Logc(ctx).Debugf("Scanning paths: %v", paths)
	found := make([]string, 0)

	checkAllDevicesExist := func() error {
		found := make([]string, 0)
		// Check if any paths present, and return nil (success) if so
		for _, p := range paths {
			dirname := p + "/block"
			if !PathExists(dirname) {
				return errors.New("device not present yet")
			}
			found = append(found, dirname)
		}
		return nil
	}

	devicesNotify := func(err error, duration time.Duration) {
		Logc(ctx).WithField("increment", duration).Debug("All devices not yet present, waiting.")
	}

	deviceBackoff := backoff.NewExponentialBackOff()
	deviceBackoff.InitialInterval = 1 * time.Second
	deviceBackoff.Multiplier = 1.414 // approx sqrt(2)
	deviceBackoff.RandomizationFactor = 0.1
	deviceBackoff.MaxElapsedTime = 5 * time.Second

	if err := backoff.RetryNotify(checkAllDevicesExist, deviceBackoff, devicesNotify); err == nil {
		Logc(ctx).Debugf("Paths found: %v", found)
		return nil
	}

	Logc(ctx).Debugf("Paths found so far: %v", found)

	checkAnyDeviceExists := func() error {
		found := make([]string, 0)
		// Check if any paths present, and return nil (success) if so
		for _, p := range paths {
			dirname := p + "/block"
			if PathExists(dirname) {
				found = append(found, dirname)
			}
		}
		if 0 == len(found) {
			return errors.New("no devices present yet")
		}
		return nil
	}

	devicesNotify = func(err error, duration time.Duration) {
		Logc(ctx).WithField("increment", duration).Debug("No devices present yet, waiting.")
	}

	deviceBackoff = backoff.NewExponentialBackOff()
	deviceBackoff.InitialInterval = 1 * time.Second
	deviceBackoff.Multiplier = 1.414 // approx sqrt(2)
	deviceBackoff.RandomizationFactor = 0.1
	deviceBackoff.MaxElapsedTime = (iSCSIDeviceDiscoveryTimeoutSecs - 5) * time.Second

	// Run the check/scan using an exponential backoff
	if err := backoff.RetryNotify(checkAnyDeviceExists, deviceBackoff, devicesNotify); err != nil {
		Logc(ctx).Warnf("Could not find all devices after %d seconds.", iSCSIDeviceDiscoveryTimeoutSecs)

		// In the case of a failure, log info about what devices are present
		if _, err := execCommand(ctx, "ls", "-al", "/dev"); err != nil {
			Logc(ctx).Warnf("Could not run ls -al /dev: %v", err)
		}
		if _, err := execCommand(ctx, "ls", "-al", "/dev/mapper"); err != nil {
			Logc(ctx).Warnf("Could not run ls -al /dev/mapper: %v", err)
		}
		if _, err := execCommand(ctx, "ls", "-al", "/dev/disk/by-path"); err != nil {
			Logc(ctx).Warnf("Could not run ls -al /dev/disk/by-path: %v", err)
		}
		if _, err := execCommand(ctx, "lsscsi"); err != nil {
			Logc(ctx).Warnf("Could not run lsscsi: %v", err)
		}
		if _, err := execCommand(ctx, "lsscsi", "-t"); err != nil {
			Logc(ctx).Warnf("Could not run lsscsi -t: %v", err)
		}
		if _, err := execCommand(ctx, "free"); err != nil {
			Logc(ctx).Warnf("Could not run free: %v", err)
		}
		return err
	}

	Logc(ctx).Debugf("Paths found: %v", found)
	return nil
}

// ISCSISupported returns true if iscsiadm is installed and in the PATH.
func ISCSISupported(ctx context.Context) bool {
	Logc(ctx).Debug(">>>> iscsi.ISCSISupported")
	defer Logc(ctx).Debug("<<<< iscsi.ISCSISupported")

	_, err := execIscsiadmCommand(ctx, "-V")
	if err != nil {
		Logc(ctx).Debug("iscsiadm tools not found on this host.")
		return false
	}
	return true
}

// ISCSIDiscoveryInfo contains information about discovered iSCSI targets.
type ISCSIDiscoveryInfo struct {
	Portal     string
	PortalIP   string
	TargetName string
}

// iSCSIDiscovery uses the 'iscsiadm' command to perform discovery.
func iSCSIDiscovery(ctx context.Context, portal string) ([]ISCSIDiscoveryInfo, error) {
	Logc(ctx).WithField("portal", portal).Debug(">>>> iscsi.iSCSIDiscovery")
	defer Logc(ctx).Debug("<<<< iscsi.iSCSIDiscovery")

	out, err := execIscsiadmCommand(ctx, "-m", "discovery", "-t", "sendtargets", "-p", portal)
	if err != nil {
		return nil, err
	}

	/*
	   iscsiadm -m discovery -t st -p 10.63.152.249:3260

	   10.63.152.249:3260,1 iqn.1992-08.com.netapp:2752.600a0980006074c20000000056b32c4d
	   10.63.152.250:3260,2 iqn.1992-08.com.netapp:2752.600a0980006074c20000000056b32c4d

	   a[0]==10.63.152.249:3260,1
	   a[1]==iqn.1992-08.com.netapp:2752.600a0980006074c20000000056b32c4d

	   For IPv6
	   [fd20:8b1e:b258:2000:f816:3eff:feec:2]:3260,1038 iqn.1992-08.com.netapp:sn.7894d7af053711ea88b100a0b886136a

	   a[0]==[fd20:8b1e:b258:2000:f816:3eff:feec:2]:3260,1038
	   a[1]==iqn.1992-08.com.netapp:sn.7894d7af053711ea88b100a0b886136a
	*/

	var discoveryInfo []ISCSIDiscoveryInfo

	lines := strings.Split(string(out), "\n")
	for _, l := range lines {
		a := strings.Fields(l)
		if len(a) >= 2 {

			portalIP := ""
			if IPv6Check(a[0]) {
				// This is an IPv6 address
				portalIP = strings.Split(a[0], "]")[0]
				portalIP += "]"
			} else {
				portalIP = strings.Split(a[0], ":")[0]
			}

			discoveryInfo = append(discoveryInfo, ISCSIDiscoveryInfo{
				Portal:     a[0],
				PortalIP:   portalIP,
				TargetName: a[1],
			})

			Logc(ctx).WithFields(log.Fields{
				"Portal":     a[0],
				"PortalIP":   portalIP,
				"TargetName": a[1],
			}).Debug("Adding iSCSI discovery info.")
		}
	}
	return discoveryInfo, nil
}

// ISCSISessionInfo contains information about iSCSI sessions.
type ISCSISessionInfo struct {
	SID        string
	Portal     string
	PortalIP   string
	TargetName string
}

// getISCSISessionInfo parses output from 'iscsiadm -m session' and returns the parsed output.
func getISCSISessionInfo(ctx context.Context) ([]ISCSISessionInfo, error) {
	Logc(ctx).Debug(">>>> iscsi.getISCSISessionInfo")
	defer Logc(ctx).Debug("<<<< iscsi.getISCSISessionInfo")

	out, err := execIscsiadmCommand(ctx, "-m", "session")
	if err != nil {
		exitErr, ok := err.(*exec.ExitError)
		if ok && exitErr.ProcessState.Sys().(syscall.WaitStatus).ExitStatus() == iSCSIErrNoObjsFound {
			Logc(ctx).Debug("No iSCSI session found.")
			return []ISCSISessionInfo{}, nil
		} else {
			Logc(ctx).WithField("error", err).Error("Problem checking iSCSI sessions.")
			return nil, err
		}
	}

	/*
	   # iscsiadm -m session

	   tcp: [3] 10.0.207.7:3260,1028 iqn.1992-08.com.netapp:sn.afbb1784f77411e582f8080027e22798:vs.3 (non-flash)
	   tcp: [4] 10.0.207.9:3260,1029 iqn.1992-08.com.netapp:sn.afbb1784f77411e582f8080027e22798:vs.3 (non-flash)

	   a[0]==tcp:
	   a[1]==[4]
	   a[2]==10.0.207.9:3260,1029
	   a[3]==iqn.1992-08.com.netapp:sn.afbb1784f77411e582f8080027e22798:vs.3
	   a[4]==(non-flash)
	*/

	var sessionInfo []ISCSISessionInfo

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, l := range lines {

		a := strings.Fields(l)
		if len(a) > 3 {
			sid := a[1]
			sid = sid[1 : len(sid)-1]

			portalIP := ""
			if IPv6Check(a[2]) {
				// This is an IPv6 address
				portalIP = strings.Split(a[2], "]")[0]
				portalIP += "]"
			} else {
				portalIP = strings.Split(a[2], ":")[0]
			}

			sessionInfo = append(sessionInfo, ISCSISessionInfo{
				SID:        sid,
				Portal:     a[2],
				PortalIP:   portalIP,
				TargetName: a[3],
			})

			Logc(ctx).WithFields(log.Fields{
				"SID":        sid,
				"Portal":     a[2],
				"PortalIP":   portalIP,
				"TargetName": a[3],
			}).Debug("Adding iSCSI session info.")
		}
	}

	return sessionInfo, nil
}

// ISCSILogout logs out from the supplied target
func ISCSILogout(ctx context.Context, targetIQN, targetPortal string) error {
	logFields := log.Fields{
		"targetIQN":    targetIQN,
		"targetPortal": targetPortal,
	}
	Logc(ctx).WithFields(logFields).Debug(">>>> iscsi.ISCSILogout")
	defer Logc(ctx).WithFields(logFields).Debug("<<<< iscsi.ISCSILogout")

	defer listAllISCSIDevices(ctx)
	if _, err := execIscsiadmCommand(ctx, "-m", "node", "-T", targetIQN, "--portal", targetPortal, "-u"); err != nil {
		Logc(ctx).WithField("error", err).Debug("Error during iSCSI logout.")
	}

	// We used to delete the iscsi "node" at this point but that could interfere with
	// another iSCSI client (such as kubelet with and "iscsi" PV) attempting to use
	// the same node.

	listAllISCSIDevices(ctx)
	return nil
}

// iSCSISessionExists checks to see if a session exists to the specified portal.
func iSCSISessionExists(ctx context.Context, portal string) (bool, error) {
	Logc(ctx).Debug(">>>> iscsi.iSCSISessionExists")
	defer Logc(ctx).Debug("<<<< iscsi.iSCSISessionExists")

	sessionInfo, err := getISCSISessionInfo(ctx)
	if err != nil {
		Logc(ctx).WithField("error", err).Error("Problem checking iSCSI sessions.")
		return false, err
	}

	for _, e := range sessionInfo {
		if strings.Contains(e.PortalIP, portal) {
			return true, nil
		}
	}

	return false, nil
}

// iSCSISessionExistsToTargetIQN checks to see if a session exists to the specified target.
func iSCSISessionExistsToTargetIQN(ctx context.Context, targetIQN string) (bool, error) {
	Logc(ctx).Debug(">>>> iscsi.iSCSISessionExistsToTargetIQN")
	defer Logc(ctx).Debug("<<<< iscsi.iSCSISessionExistsToTargetIQN")

	sessionInfo, err := getISCSISessionInfo(ctx)
	if err != nil {
		Logc(ctx).WithField("error", err).Error("Problem checking iSCSI sessions.")
		return false, err
	}

	for _, e := range sessionInfo {
		if e.TargetName == targetIQN {
			return true, nil
		}
	}

	return false, nil
}

// portalsToLogin checks to see if session to for all the specified portals exist to the specified
// target. If a session does not exist for a give portal it is added to list of portals that Trident
// needs to login to.
func portalsToLogin(ctx context.Context, targetIQN string, portals []string) ([]string, bool, error) {
	logFields := log.Fields{
		"targetIQN": targetIQN,
		"portals":   portals,
	}

	Logc(ctx).WithFields(logFields).Debug(">>>> iscsi.portalsToLogin")
	defer Logc(ctx).Debug("<<<< iscsi.portalsToLogin")

	portalsNotLoggedIn := make([]string, len(portals))
	copy(portalsNotLoggedIn, portals)

	sessionInfo, err := getISCSISessionInfo(ctx)
	if err != nil {
		Logc(ctx).WithField("error", err).Error("Problem checking iSCSI sessions.")
		return portalsNotLoggedIn, false, err
	}

	for _, e := range sessionInfo {
		if e.TargetName == targetIQN {

			// Portals (portalsNotLoggedIn) may/may not contain anything after ":", so instead of matching complete
			// portal value (with e.Portal), check if e.Portal's IP address matches portal's IP address
			matchFunc := func(main, val string) bool {
				mainIpAddress := getHostportIP(main)
				valIpAddress := getHostportIP(val)

				return mainIpAddress == valIpAddress
			}

			portalsNotLoggedIn = RemoveStringFromSliceConditionally(portalsNotLoggedIn, e.Portal, matchFunc)
		}
	}

	loggedIn := len(portals) != len(portalsNotLoggedIn)
	return portalsNotLoggedIn, loggedIn, nil
}

// portalsIpsToLogin checks to see if session to for all the specified portal IPs exist to the specified
// target. If a session does not exist for a give portal IP it is added to list of portals IPs that Trident
// needs to login to.
func portalsIpsToLogin(ctx context.Context, targetIQN string, portalsIps []string) ([]string, bool, error) {
	logFields := log.Fields{
		"targetIQN":  targetIQN,
		"portalsIps": portalsIps,
	}

	Logc(ctx).WithFields(logFields).Debug(">>>> iscsi.portalsIpsToLogin")
	defer Logc(ctx).Debug("<<<< iscsi.portalsIpsToLogin")

	portalIpsNotLoggedIn := make([]string, len(portalsIps))
	copy(portalIpsNotLoggedIn, portalsIps)

	sessionInfo, err := getISCSISessionInfo(ctx)
	if err != nil {
		Logc(ctx).WithField("error", err).Error("Problem checking iSCSI sessions.")
		return portalIpsNotLoggedIn, false, err
	}

	for _, e := range sessionInfo {
		if e.TargetName == targetIQN {
			portalIpsNotLoggedIn = RemoveStringFromSlice(portalIpsNotLoggedIn, e.PortalIP)
		}
	}

	loggedIn := len(portalsIps) != len(portalIpsNotLoggedIn)
	return portalIpsNotLoggedIn, loggedIn, nil
}

// formatPortal returns the iSCSI portal string, appending a port number if one isn't
// already present, and also appending a target portal group tag if one is not present
func formatPortal(portal string) string {
	if portalPortPattern.MatchString(portal) {
		return portal
	} else {
		return portal + ":3260"
	}
}

// iSCSIScanTargetLUN scans a single LUN on an iSCSI target to discover it.
func iSCSIScanTargetLUN(ctx context.Context, lunID int, hosts []int) error {
	fields := log.Fields{"hosts": hosts, "lunID": lunID}
	Logc(ctx).WithFields(fields).Debug(">>>> iscsi.iSCSIScanTargetLUN")
	defer Logc(ctx).WithFields(fields).Debug("<<<< iscsi.iSCSIScanTargetLUN")

	var (
		f   *os.File
		err error
	)

	listAllISCSIDevices(ctx)
	for _, hostNumber := range hosts {

		filename := fmt.Sprintf(chrootPathPrefix+"/sys/class/scsi_host/host%d/scan", hostNumber)
		if f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0o200); err != nil {
			Logc(ctx).WithField("file", filename).Warning("Could not open file for writing.")
			return err
		}

		scanCmd := fmt.Sprintf("0 0 %d", lunID)
		if written, err := f.WriteString(scanCmd); err != nil {
			Logc(ctx).WithFields(log.Fields{"file": filename, "error": err}).Warning("Could not write to file.")
			f.Close()
			return err
		} else if written == 0 {
			Logc(ctx).WithField("file", filename).Warning("No data written to file.")
			f.Close()
			return fmt.Errorf("no data written to %s", filename)
		}

		f.Close()

		listAllISCSIDevices(ctx)
		Logc(ctx).WithFields(log.Fields{
			"scanCmd":  scanCmd,
			"scanFile": filename,
		}).Debug("Invoked single-LUN scan.")
	}

	return nil
}

// IsAlreadyAttached checks if there is already an established iSCSI session to the specified LUN.
func IsAlreadyAttached(ctx context.Context, lunID int, targetIqn string) bool {
	hostSessionMap := GetISCSIHostSessionMapForTarget(ctx, targetIqn)
	if len(hostSessionMap) == 0 {
		return false
	}

	paths := getSysfsBlockDirsForLUN(lunID, hostSessionMap)

	devices, err := getDevicesForLUN(paths)
	if nil != err {
		return false
	}

	return 0 < len(devices)
}

// getLunSerial get Linux's idea of what the LUN serial number is
func getLunSerial(ctx context.Context, path string) (string, error) {
	Logc(ctx).WithField("path", path).Debug("Get LUN Serial")
	// We're going to read the SCSI VPD page 80 serial number
	// information. Linux helpfully provides this through sysfs
	// so we don't need to open the device and send the ioctl
	// ourselves.
	filename := path + "/vpd_pg80"
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if 4 > len(b) || 0x80 != b[1] {
		Logc(ctx).WithFields(log.Fields{
			"data": b,
		}).Error("VPD page 80 format check failed")
		return "", fmt.Errorf("malformed VPD page 80 data")
	}
	length := int(binary.BigEndian.Uint16(b[2:4]))
	if len(b) != length+4 {
		Logc(ctx).WithFields(log.Fields{
			"actual":   len(b),
			"expected": length + 4,
		}).Error("VPD page 80 length check failed")
		return "", fmt.Errorf("incorrect length for VPD page 80 serial number")
	}
	return string(b[4:]), nil
}

// purgeOneLun issues a delete for one LUN, based on the sysfs path
func purgeOneLun(ctx context.Context, path string) error {
	Logc(ctx).WithField("path", path).Debug("Purging one LUN")
	filename := path + "/delete"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0o200)
	if err != nil {
		Logc(ctx).WithField("file", filename).Warning("Could not open file for writing.")
		return err
	}
	defer f.Close()

	// Deleting a LUN is achieved by writing the string "1" to the "delete" file
	written, err := f.WriteString("1")
	if err != nil {
		Logc(ctx).WithFields(log.Fields{"file": filename, "error": err}).Warning("Could not write to file.")
		return err
	}
	if written == 0 {
		Logc(ctx).WithField("file", filename).Warning("No data written to file.")
		return errors.New("too few bytes written to sysfs file")
	}

	return nil
}

// rescanOneLun issues a rescan for one LUN, based on the sysfs path
func rescanOneLun(ctx context.Context, path string) error {
	Logc(ctx).WithField("path", path).Debug("Rescaning one LUN")
	filename := path + "/rescan"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0o200)
	if err != nil {
		Logc(ctx).WithField("file", filename).Warning("Could not open file for writing.")
		return err
	}
	defer f.Close()

	written, err := f.WriteString("1")
	if err != nil {
		Logc(ctx).WithFields(log.Fields{"file": filename, "error": err}).Warning("Could not write to file.")
		return err
	}
	if written == 0 {
		Logc(ctx).WithField("file", filename).Warning("No data written to file.")
		return errors.New("too few bytes written to sysfs file")
	}

	return nil
}

// handleInvalidSerials checks the LUN serial number for each path of a given LUN, and
// if it doesn't match the expected value, runs a handler function.
func handleInvalidSerials(
	ctx context.Context, lunID int, targetIqn, expectedSerial string,
	handler func(ctx context.Context, path string) error,
) error {
	if "" == expectedSerial {
		// Empty string means don't care
		return nil
	}

	hostSessionMap := GetISCSIHostSessionMapForTarget(ctx, targetIqn)
	paths := getSysfsBlockDirsForLUN(lunID, hostSessionMap)
	for _, path := range paths {
		serial, err := getLunSerial(ctx, path)
		if err != nil {
			if os.IsNotExist(err) {
				// LUN either isn't scanned yet, or this kernel
				// doesn't support VPD page 80 in sysfs. Assume
				// correctness and move on
				Logc(ctx).WithFields(log.Fields{
					"lun":    lunID,
					"target": targetIqn,
					"path":   path,
				}).Debug("LUN serial check skipped")
				continue
			}
			return err
		}

		if serial != expectedSerial {
			Logc(ctx).WithFields(log.Fields{
				"expected": expectedSerial,
				"actual":   serial,
				"lun":      lunID,
				"target":   targetIqn,
				"path":     path,
			}).Warn("LUN serial check failed")
			err = handler(ctx, path)
			if err != nil {
				return err
			}
		} else {
			Logc(ctx).WithFields(log.Fields{
				"serial": serial,
				"lun":    lunID,
				"target": targetIqn,
				"path":   path,
			}).Debug("LUN serial check passed")
		}
	}

	return nil
}

// GetISCSIHostSessionMapForTarget returns a map of iSCSI host numbers to iSCSI session numbers
// for a given iSCSI target.
func GetISCSIHostSessionMapForTarget(ctx context.Context, iSCSINodeName string) map[int]int {
	fields := log.Fields{"iSCSINodeName": iSCSINodeName}
	Logc(ctx).WithFields(fields).Debug(">>>> iscsi.GetISCSIHostSessionMapForTarget")
	defer Logc(ctx).WithFields(fields).Debug("<<<< iscsi.GetISCSIHostSessionMapForTarget")

	var (
		hostNumber    int
		sessionNumber int
	)

	hostSessionMap := make(map[int]int)

	sysPath := chrootPathPrefix + "/sys/class/iscsi_host/"
	if hostDirs, err := ioutil.ReadDir(sysPath); err != nil {
		Logc(ctx).WithField("error", err).Errorf("Could not read %s", sysPath)
		return hostSessionMap
	} else {
		for _, hostDir := range hostDirs {

			hostName := hostDir.Name()
			if !strings.HasPrefix(hostName, "host") {
				continue
			} else if hostNumber, err = strconv.Atoi(strings.TrimPrefix(hostName, "host")); err != nil {
				Logc(ctx).WithField("host", hostName).Error("Could not parse host number")
				continue
			}

			devicePath := sysPath + hostName + "/device/"
			if deviceDirs, err := ioutil.ReadDir(devicePath); err != nil {
				Logc(ctx).WithFields(log.Fields{
					"error":      err,
					"devicePath": devicePath,
				}).Error("Could not read device path.")
				return hostSessionMap
			} else {
				for _, deviceDir := range deviceDirs {

					sessionName := deviceDir.Name()
					if !strings.HasPrefix(sessionName, "session") {
						continue
					} else if sessionNumber, err = strconv.Atoi(strings.TrimPrefix(sessionName,
						"session")); err != nil {
						Logc(ctx).WithField("session", sessionName).Error("Could not parse session number")
						continue
					}

					targetNamePath := devicePath + sessionName + "/iscsi_session/" + sessionName + "/targetname"
					if targetName, err := ioutil.ReadFile(targetNamePath); err != nil {
						Logc(ctx).WithFields(log.Fields{
							"path":  targetNamePath,
							"error": err,
						}).Error("Could not read targetname file")
					} else if strings.TrimSpace(string(targetName)) == iSCSINodeName {

						Logc(ctx).WithFields(log.Fields{
							"hostNumber":    hostNumber,
							"sessionNumber": sessionNumber,
						}).Debug("Found iSCSI host/session.")

						hostSessionMap[hostNumber] = sessionNumber
					}
				}
			}
		}
	}

	return hostSessionMap
}

// ISCSITargetHasMountedDevice returns true if this host has any mounted devices on the specified target.
func ISCSITargetHasMountedDevice(ctx context.Context, targetIQN string) (bool, error) {
	mountedISCSIDevices, err := GetMountedISCSIDevices(ctx)
	if err != nil {
		return false, err
	}

	for _, device := range mountedISCSIDevices {
		if device.IQN == targetIQN {
			return true, nil
		}
	}

	return false, nil
}

// multipathdIsRunning returns true if the multipath daemon is running.
func multipathdIsRunning(ctx context.Context) bool {
	Logc(ctx).Debug(">>>> iscsi.multipathdIsRunning")
	defer Logc(ctx).Debug("<<<< iscsi.multipathdIsRunning")

	out, err := execCommand(ctx, "pgrep", "multipathd")
	if err == nil {
		pid := strings.TrimSpace(string(out))
		if pidRegex.MatchString(pid) {
			Logc(ctx).WithField("pid", pid).Debug("multipathd is running")
			return true
		}
	} else {
		Logc(ctx).Error(err)
	}

	out, err = execCommand(ctx, "multipathd", "show", "daemon")
	if err == nil {
		if pidRunningOrIdleRegex.MatchString(string(out)) {
			Logc(ctx).Debug("multipathd is running")
			return true
		}
	} else {
		Logc(ctx).Error(err)
	}

	return false
}

// filterTargets parses the output of iscsiadm -m node or -m discoverydb -t st -D
// and returns the target IQNs for a given portal
func filterTargets(output, tp string) []string {
	regex := regexp.MustCompile(`^([^,]+),(-?\d+)\s+(.+)$`)
	targets := make([]string, 0)
	for _, line := range strings.Split(output, "\n") {
		if 0 == len(line) {
			continue
		}
		matches := regex.FindStringSubmatch(line)
		if 4 == len(matches) {
			portal := matches[1]
			iqn := matches[3]
			if tp == portal {
				targets = append(targets, iqn)
			}
		}
	}
	return targets
}

// getTargets gets a list of discovered iSCSI targets
func getTargets(ctx context.Context, tp string) ([]string, error) {
	Logc(ctx).WithFields(log.Fields{
		"Portal": tp,
	}).Debug(">>>> iscsi.getTargets")
	defer Logc(ctx).Debug("<<<< iscsi.getTargets")

	output, err := execCommand(ctx, "iscsiadm", "-m", "node")
	if nil != err {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if status, ok := exitErr.Sys().(syscall.WaitStatus); ok {
				if iSCSIErrNoObjsFound == status.ExitStatus() {
					Logc(ctx).Debug("No iSCSI nodes found.")
					// No records
					return nil, nil
				}
			}
		}
		Logc(ctx).WithFields(log.Fields{
			"error":  err,
			"output": string(output),
		}).Error("Failed to list nodes")
		return nil, fmt.Errorf("failed to list nodes: %v", err)
	}
	return filterTargets(string(output), tp), nil
}

func updateDiscoveryDb(ctx context.Context, tp, iface, key, value string) error {
	Logc(ctx).WithFields(log.Fields{
		"Key":       key,
		"Value":     value,
		"Portal":    tp,
		"Interface": iface,
	}).Debug(">>>> iscsi.updateDiscoveryDb")
	defer Logc(ctx).Debug("<<<< iscsi.updateDiscoveryDb")

	output, err := execCommand(ctx, "iscsiadm", "-m", "discoverydb",
		"-t", "st", "-p", tp, "-I", iface, "-o", "update", "-n", key, "-v", value)
	if err != nil {
		Logc(ctx).WithFields(log.Fields{
			"portal": tp,
			"error":  err,
			"key":    key,
			"value":  value,
			"output": string(output),
		}).Error("Failed to update discovery DB.")
		return fmt.Errorf("failed to update discovery db: %v", err)
	}

	return nil
}

// ensureIscsiTarget creates the iSCSI target if we haven't done so already
// This function first checks if the target is already known, and if not,
// uses sendtargets to try to discover it. Because sendtargets will find
// all of the targets given just 1 portal, it will be very common to hit
// the case where the target is already known.
// Note: Adding iSCSI targets using sendtargets rather than static discover
// ensures that targets are added with the correct target group portal tags.
func ensureIscsiTarget(
	ctx context.Context, tp, targetIqn, username, password, targetUsername, targetInitiatorSecret, iface string,
) error {
	Logc(ctx).WithFields(log.Fields{
		"IQN":       targetIqn,
		"Portal":    tp,
		"Interface": iface,
	}).Debug(">>>> iscsi.ensureIscsiTarget")
	defer Logc(ctx).Debug("<<<< iscsi.ensureIscsiTarget")

	targets, err := getTargets(ctx, tp)
	if err != nil {
		// Already logged
		return err
	}
	for _, iqn := range targets {
		if targetIqn == iqn {
			Logc(ctx).WithField("Target", iqn).Info("Target exists already")
			return nil
		}
	}

	if "" != username && "" != password {
		// To do discovery on a CHAP-enabled target, we need to set the CHAP
		// secrets on the discoverydb object before making the sendtargets
		// call.

		// Ignore result
		_, _ = execCommand(ctx, "iscsiadm", "-m", "discoverydb", "-t", "st", "-p", tp, "-I", iface, "-o", "new")

		err = updateDiscoveryDb(ctx, tp, iface, "discovery.sendtargets.auth.authmethod", "CHAP")
		if err != nil {
			// Already logged
			return err
		}

		err = updateDiscoveryDb(ctx, tp, iface, "discovery.sendtargets.auth.username", username)
		if err != nil {
			// Already logged
			return err
		}

		err = updateDiscoveryDb(ctx, tp, iface, "discovery.sendtargets.auth.password", password)
		if err != nil {
			// Already logged
			return err
		}

		if targetUsername != "" && targetInitiatorSecret != "" {
			// Bidirectional CHAP case

			err = updateDiscoveryDb(ctx, tp, iface, "discovery.sendtargets.auth.username_in", targetUsername)
			if err != nil {
				// Already logged
				return err
			}

			err = updateDiscoveryDb(ctx, tp, iface, "discovery.sendtargets.auth.password_in", targetInitiatorSecret)
			if err != nil {
				// Already logged
				return err
			}
		}
	}

	// Discovery is here. This will populate the iscsiadm database with the
	// ALL of the nodes known to the given portal.
	output, err := execCommand(ctx, "iscsiadm", "-m", "discoverydb",
		"-t", "st", "-p", tp, "-I", iface, "-D")
	if err != nil {
		Logc(ctx).WithFields(log.Fields{
			"portal": tp,
			"error":  err,
			"output": string(output),
		}).Error("Failed to discover targets")
		return fmt.Errorf("failed to discover targets: %v", err)
	}

	targets = filterTargets(string(output), tp)
	for _, iqn := range targets {
		if targetIqn == iqn {
			Logc(ctx).WithField("Target", iqn).Info("Target discovered successfully")
			// Discovered successfully
			return nil
		}
	}

	Logc(ctx).WithFields(log.Fields{
		"portal": tp,
		"iqn":    targetIqn,
	}).Warning("Target not discovered")
	return fmt.Errorf("target not discovered")
}

// loginISCSITarget logs in to an iSCSI target.
func configureISCSITarget(ctx context.Context, iqn, portal, name, value string) error {
	Logc(ctx).WithFields(log.Fields{
		"IQN":    iqn,
		"Portal": portal,
		"Name":   name,
		"Value":  value,
	}).Debug(">>>> iscsi.configureISCSITarget")
	defer Logc(ctx).Debug("<<<< iscsi.configureISCSITarget")

	args := []string{"-m", "node", "-T", iqn, "-p", formatPortal(portal), "-o", "update", "-n", name, "-v", value}
	if _, err := execIscsiadmCommand(ctx, args...); err != nil {
		Logc(ctx).WithField("error", err).Warn("Error configuring iSCSI target.")
		return err
	}
	return nil
}

// loginISCSITarget logs in to an iSCSI target.
func loginISCSITarget(ctx context.Context, iqn, portal string) error {
	Logc(ctx).WithFields(log.Fields{
		"IQN":    iqn,
		"Portal": portal,
	}).Debug(">>>> iscsi.loginISCSITarget")
	defer Logc(ctx).Debug("<<<< iscsi.loginISCSITarget")

	args := []string{"-m", "node", "-T", iqn, "-l", "-p", formatPortal(portal)}
	listAllISCSIDevices(ctx)
	if _, err := execIscsiadmCommandWithTimeout(ctx, 10, args...); err != nil {
		Logc(ctx).WithField("error", err).Error("Error logging in to iSCSI target.")
		return err
	}
	listAllISCSIDevices(ctx)
	return nil
}

// loginWithChap will login to the iSCSI target with the supplied credentials.
func loginWithChap(
	ctx context.Context, tiqn, portal, username, password, targetUsername, targetInitiatorSecret, iface string,
) error {
	logFields := log.Fields{
		"IQN":                   tiqn,
		"portal":                portal,
		"username":              REDACTED,
		"password":              REDACTED,
		"targetUsername":        REDACTED,
		"targetInitiatorSecret": REDACTED,
		"iface":                 iface,
	}
	Logc(ctx).WithFields(logFields).Debug(">>>> iscsi.loginWithChap")
	defer Logc(ctx).Debug("<<<< iscsi.loginWithChap")

	args := []string{"-m", "node", "-T", tiqn, "-p", formatPortal(portal)}

	secretsToRedact := map[string]string{
		"--value=" + username:              "--value=" + REDACTED,
		"--value=" + password:              "--value=" + REDACTED,
		"--value=" + targetUsername:        "--value=" + REDACTED,
		"--value=" + targetInitiatorSecret: "--value=" + REDACTED,
	}

	listAllISCSIDevices(ctx)
	if err := ensureIscsiTarget(ctx, formatPortal(portal), tiqn, username, password, targetUsername,
		targetInitiatorSecret, iface); err != nil {
		Logc(ctx).Error("Error running iscsiadm node create.")
		return err
	}

	authMethodArgs := append(args, []string{"--op=update", "--name", "node.session.auth.authmethod", "--value=CHAP"}...)
	if _, err := execIscsiadmCommand(ctx, authMethodArgs...); err != nil {
		Logc(ctx).Error("Error running iscsiadm set authmethod.")
		return err
	}

	authUserArgs := append(args,
		[]string{"--op=update", "--name", "node.session.auth.username", "--value=" + username}...)
	if _, err := execIscsiadmCommandRedacted(ctx, authUserArgs, secretsToRedact); err != nil {
		Logc(ctx).Error("Error running iscsiadm set authuser.")
		return err
	}

	authPasswordArgs := append(args,
		[]string{"--op=update", "--name", "node.session.auth.password", "--value=" + password}...)
	if _, err := execIscsiadmCommandRedacted(ctx, authPasswordArgs, secretsToRedact); err != nil {
		Logc(ctx).Error("Error running iscsiadm set authpassword.")
		return err
	}

	if targetUsername != "" && targetInitiatorSecret != "" {
		targetAuthUserArgs := append(args,
			[]string{"--op=update", "--name", "node.session.auth.username_in", "--value=" + targetUsername}...)
		if _, err := execIscsiadmCommandRedacted(ctx, targetAuthUserArgs, secretsToRedact); err != nil {
			Logc(ctx).Error("Error running iscsiadm set authuser_in.")
			return err
		}

		targetAuthPasswordArgs := append(args,
			[]string{"--op=update", "--name", "node.session.auth.password_in", "--value=" + targetInitiatorSecret}...)
		if _, err := execIscsiadmCommandRedacted(ctx, targetAuthPasswordArgs, secretsToRedact); err != nil {
			Logc(ctx).Error("Error running iscsiadm set authpassword_in.")
			return err
		}
	}

	loginArgs := append(args, []string{"--login"}...)
	if _, err := execIscsiadmCommandWithTimeout(ctx, 10, loginArgs...); err != nil {
		msg := "Error running iscsiadm login."
		Logc(ctx).Error(msg)
		return AuthError(msg)
	}
	listAllISCSIDevices(ctx)
	return nil
}

func EnsureISCSISessions(ctx context.Context, targetIQN, iface string, portalsIps []string) bool {
	logFields := log.Fields{
		"targetIQN":  targetIQN,
		"portalsIps": portalsIps,
	}

	Logc(ctx).WithFields(logFields).Debug(">>>> iscsi.EnsureISCSISessions")
	defer Logc(ctx).Debug("<<<< iscsi.EnsureISCSISessions")

	loggedInPortals := make([]string, 0)

	for _, portalIp := range portalsIps {
		listAllISCSIDevices(ctx)

		formattedPortal := formatPortal(portalIp)
		if err := ensureIscsiTarget(ctx, formattedPortal, targetIQN, "", "", "", "",
			iface); nil != err {
			Logc(ctx).WithFields(log.Fields{
				"tp":        formattedPortal,
				"targetIqn": targetIQN,
				"iface":     iface,
				"err":       err,
			}).Errorf("unable to ensure iSCSI target exists: %v", err)
			continue
		}

		// Set scanning to manual
		// Swallow this error, someone is running an old version of Debian/Ubuntu
		_ = configureISCSITarget(ctx, targetIQN, portalIp, "node.session.scan", "manual")

		// Update replacement timeout
		timeout_param := "node.session.timeo.replacement_timeout"
		if err := configureISCSITarget(ctx, targetIQN, portalIp, timeout_param, "5"); err != nil {
			Logc(ctx).WithFields(log.Fields{
				"iqn":    targetIQN,
				"portal": portalIp,
				"name":   timeout_param,
				"value":  "5",
				"err":    err,
			}).Errorf("set replacement timeout failed: %v", err)
			continue
		}

		// Log in to target
		if err := loginISCSITarget(ctx, targetIQN, portalIp); err != nil {
			Logc(ctx).WithFields(log.Fields{
				"err":      err,
				"portalIP": portalIp,
			}).Error("Login to iSCSI target failed.")
			continue
		}

		loggedInPortals = append(loggedInPortals, portalIp)
	}

	var successfulLogin bool

	for _, portalIp := range loggedInPortals {
		// Recheck to ensure a session is now open
		sessionExists, err := iSCSISessionExists(ctx, portalIp)
		if err != nil {
			Logc(ctx).WithFields(log.Fields{
				"err":      err,
				"portalIP": portalIp,
			}).Error("Could not recheck for iSCSI session.")
			continue
		}

		if !sessionExists {
			Logc(ctx).Errorf("Expected iSCSI session %v NOT found, please login to the iSCSI portal", portalIp)
			continue
		}

		successfulLogin = true

		Logc(ctx).WithField("portalIp", portalIp).Debug("Session established with iSCSI portal.")
	}

	return successfulLogin
}

func EnsureISCSISessionsWithPortalDiscovery(ctx context.Context, hostDataIPs []string) error {
	for _, ip := range hostDataIPs {
		if err := EnsureISCSISessionWithPortalDiscovery(ctx, ip); nil != err {
			return err
		}
	}
	return nil
}

func EnsureISCSISessionWithPortalDiscovery(ctx context.Context, hostDataIP string) error {
	Logc(ctx).WithField("hostDataIP", hostDataIP).Debug(">>>> iscsi.EnsureISCSISessionWithPortalDiscovery")
	defer Logc(ctx).Debug("<<<< iscsi.EnsureISCSISessionWithPortalDiscovery")

	// Ensure iSCSI is supported on system
	if !ISCSISupported(ctx) {
		return errors.New("iSCSI support not detected")
	}

	// Ensure iSCSI session exists for the specified iSCSI portal
	sessionExists, err := iSCSISessionExists(ctx, hostDataIP)
	if err != nil {
		return fmt.Errorf("could not check for iSCSI session: %v", err)
	}
	if !sessionExists {

		// Run discovery in case we haven't seen this target from this host
		targets, err := iSCSIDiscovery(ctx, hostDataIP)
		if err != nil {
			return fmt.Errorf("could not run iSCSI discovery: %v", err)
		}
		if len(targets) == 0 {
			return errors.New("iSCSI discovery found no targets")
		}

		Logc(ctx).WithFields(log.Fields{
			"Targets": targets,
		}).Debug("Found matching iSCSI targets.")

		// Determine which target matches the portal we requested
		targetIndex := -1
		for i, target := range targets {
			if strings.Contains(target.PortalIP, hostDataIP) {
				targetIndex = i
				break
			}
		}

		if targetIndex == -1 {
			return fmt.Errorf("iSCSI discovery found no targets with portal %s", hostDataIP)
		}

		// To enable multipath, log in to each discovered target with the same IQN (target name)
		targetName := targets[targetIndex].TargetName
		for _, target := range targets {
			if target.TargetName == targetName {
				// Set scan to manual
				// Swallow this error, someone is running an old version of Debian/Ubuntu
				_ = configureISCSITarget(ctx, target.TargetName, target.PortalIP, "node.session.scan", "manual")

				// Update replacement timeout
				err = configureISCSITarget(
					ctx, target.TargetName, target.PortalIP, "node.session.timeo.replacement_timeout", "5")
				if err != nil {
					return fmt.Errorf("set replacement timeout failed: %v", err)
				}
				// Log in to target
				err = loginISCSITarget(ctx, target.TargetName, target.PortalIP)
				if err != nil {
					return fmt.Errorf("login to iSCSI target failed: %v", err)
				}
			}
		}

		// Recheck to ensure a session is now open
		sessionExists, err = iSCSISessionExists(ctx, hostDataIP)
		if err != nil {
			return fmt.Errorf("could not recheck for iSCSI session: %v", err)
		}
		if !sessionExists {
			return fmt.Errorf("Expected iSCSI session %v NOT found, please login to the iSCSI portal", hostDataIP)
		}
	}

	Logc(ctx).WithField("hostDataIP", hostDataIP).Debug("Found session to iSCSI portal.")

	return nil
}

// execIscsiadmCommand uses the 'iscsiadm' command to perform operations without logging specified secrets
func execIscsiadmCommandRedacted(ctx context.Context, args []string, secretsToRedact map[string]string) ([]byte, error) {
	return execCommandRedacted(ctx, "iscsiadm", args, secretsToRedact)
}

// execIscsiadmCommand uses the 'iscsiadm' command to perform operations
func execIscsiadmCommand(ctx context.Context, args ...string) ([]byte, error) {
	return execCommand(ctx, "iscsiadm", args...)
}

// execIscsiadmCommandWithTimeout uses the 'iscsiadm' command to perform operations with timeout
func execIscsiadmCommandWithTimeout(ctx context.Context, timeout time.Duration, args ...string) ([]byte, error) {
	return execCommandWithTimeout(ctx, "iscsiadm", timeout, true, args...)
}

// SafeToLogOut looks for remaining block devices on a given iSCSI host, and returns
// true if there are none, indicating that logging out would be safe.
func SafeToLogOut(ctx context.Context, hostNumber, sessionNumber int) bool {
	Logc(ctx).Debug(">>>> iscsi.SafeToLogOut")
	defer Logc(ctx).Debug("<<<< iscsi.SafeToLogOut")

	devicePath := fmt.Sprintf("/sys/class/iscsi_host/host%d/device", hostNumber)

	// The list of block devices on the scsi bus will be in a
	// directory called "target%d:%d:%d".
	// See drivers/scsi/scsi_scan.c in Linux
	// We assume the channel/bus and device/controller are always zero for iSCSI
	targetPath := devicePath + fmt.Sprintf("/session%d/target%d:0:0", sessionNumber, hostNumber)
	dirs, err := ioutil.ReadDir(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		Logc(ctx).WithFields(log.Fields{
			"path":  targetPath,
			"error": err,
		}).Warn("Failed to read dir")
		return true
	}

	// The existence of any directories here indicate devices that
	// still exist, so report unsafe
	if 0 < len(dirs) {
		return false
	}

	return true
}

// identifyFindMultipathsValue reads /etc/multipath.conf and identifies find_multipaths value (if set)
func identifyFindMultipathsValue(ctx context.Context) (string, error) {
	if output, err := execCommandWithTimeout(ctx, "multipathd", 5, false, "show", "config"); err != nil {
		Logc(ctx).WithFields(log.Fields{
			"error": err,
		}).Error("Could not read multipathd configuration")

		return "", fmt.Errorf("could not read multipathd configuration: %v", err)
	} else {
		findMultipathsValue := GetFindMultipathValue(string(output))

		Logc(ctx).WithField("findMultipathsValue", findMultipathsValue).Debug("Multipath find_multipaths value found.")

		return findMultipathsValue, nil
	}
}
