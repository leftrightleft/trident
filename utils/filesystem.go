// Copyright 2022 NetApp, Inc. All Rights Reserved.

package utils

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"
	"go.uber.org/multierr"

	. "github.com/netapp/trident/logger"
)

const (
	// Filesystem types
	fsXfs  = "xfs"
	fsExt3 = "ext3"
	fsExt4 = "ext4"
	fsRaw  = "raw"
)

// DFInfo data structure for wrapping the parsed output from the 'df' command
type DFInfo struct {
	Target string
	Source string
}

// GetDFOutput returns parsed DF output
func GetDFOutput(ctx context.Context) ([]DFInfo, error) {
	Logc(ctx).Debug(">>>> filesystem.GetDFOutput")
	defer Logc(ctx).Debug("<<<< filesystem.GetDFOutput")

	var result []DFInfo
	out, err := execCommand(ctx, "df", "--output=target,source")
	if err != nil {
		// df returns an error if there's a stale file handle that we can
		// safely ignore. There may be other reasons. Consider it a warning if
		// it printed anything to stdout.
		if len(out) == 0 {
			Logc(ctx).Error("Error encountered gathering df output.")
			return nil, err
		}
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	for _, l := range lines {

		a := strings.Fields(l)
		if len(a) > 1 {
			result = append(result, DFInfo{
				Target: a[0],
				Source: a[1],
			})
		}
	}
	if len(result) > 1 {
		return result[1:], nil
	}
	return result, nil
}

// formatVolume creates a filesystem for the supplied device of the supplied type.
func formatVolume(ctx context.Context, device, fstype string) error {
	logFields := log.Fields{"device": device, "fsType": fstype}
	Logc(ctx).WithFields(logFields).Debug(">>>> filesystem.formatVolume")
	defer Logc(ctx).WithFields(logFields).Debug("<<<< filesystem.formatVolume")

	maxDuration := 30 * time.Second

	formatVolume := func() error {
		var err error

		switch fstype {
		case fsXfs:
			_, err = execCommand(ctx, "mkfs.xfs", "-f", device)
		case fsExt3:
			_, err = execCommand(ctx, "mkfs.ext3", "-F", device)
		case fsExt4:
			_, err = execCommand(ctx, "mkfs.ext4", "-F", device)
		default:
			return fmt.Errorf("unsupported file system type: %s", fstype)
		}

		return err
	}

	formatNotify := func(err error, duration time.Duration) {
		Logc(ctx).WithField("increment", duration).Debug("Format failed, retrying.")
	}

	formatBackoff := backoff.NewExponentialBackOff()
	formatBackoff.InitialInterval = 2 * time.Second
	formatBackoff.Multiplier = 2
	formatBackoff.RandomizationFactor = 0.1
	formatBackoff.MaxElapsedTime = maxDuration

	// Run the check/scan using an exponential backoff
	if err := backoff.RetryNotify(formatVolume, formatBackoff, formatNotify); err != nil {
		Logc(ctx).Warnf("Could not format device after %3.2f seconds.", maxDuration.Seconds())
		return err
	}

	Logc(ctx).WithFields(logFields).Info("Device formatted.")
	return nil
}

// repairVolume runs fsck on a volume.
func repairVolume(ctx context.Context, device, fstype string) (err error) {
	logFields := log.Fields{"device": device, "fsType": fstype}
	Logc(ctx).WithFields(logFields).Debug(">>>> filesystem.repairVolume")
	defer Logc(ctx).WithFields(logFields).Debug("<<<< filesystem.repairVolume")

	switch fstype {
	case "xfs":
		break // fsck.xfs does nothing
	case "ext3":
		_, err = execCommand(ctx, "fsck.ext3", "-p", device)
	case "ext4":
		_, err = execCommand(ctx, "fsck.ext4", "-p", device)
	default:
		return fmt.Errorf("unsupported file system type: %s", fstype)
	}

	return
}

// ExpandFilesystemOnNode will expand the filesystem of an already expanded volume.
func ExpandFilesystemOnNode(
	ctx context.Context, publishInfo *VolumePublishInfo, stagedTargetPath, fsType, mountOptions string,
) (int64, error) {
	var err error
	devicePath := publishInfo.DevicePath
	expansionMountPoint := publishInfo.StagingMountpoint

	logFields := log.Fields{
		"devicePath":        devicePath,
		"stagedTargetPath":  stagedTargetPath,
		"mountOptions":      mountOptions,
		"filesystemType":    fsType,
		"stagingMountpoint": expansionMountPoint,
	}
	Logc(ctx).WithFields(logFields).Debug(">>>> filesystem.ExpandFilesystemOnNode")
	defer Logc(ctx).WithFields(logFields).Debug("<<<< filesystem.ExpandFilesystemOnNode")

	if expansionMountPoint == "" {
		expansionMountPoint, err = mountFilesystemForResize(ctx, devicePath, stagedTargetPath, mountOptions)
		if err != nil {
			return 0, err
		}
		defer func() {
			err = multierr.Append(err, RemoveMountPoint(ctx, expansionMountPoint))
		}()
	}

	// Don't need to verify the filesystem type as the resize utilities will throw an error if the filesystem
	// is not the correct type.
	var size int64
	switch fsType {
	case "xfs":
		size, err = expandFilesystem(ctx, "xfs_growfs", expansionMountPoint, expansionMountPoint)
	case "ext3", "ext4":
		size, err = expandFilesystem(ctx, "resize2fs", devicePath, expansionMountPoint)
	default:
		err = fmt.Errorf("unsupported file system type: %s", fsType)
	}
	if err != nil {
		return 0, err
	}
	return size, err
}

func expandFilesystem(ctx context.Context, cmd, cmdArguments, tmpMountPoint string) (int64, error) {
	logFields := log.Fields{
		"cmd":           cmd,
		"cmdArguments":  cmdArguments,
		"tmpMountPoint": tmpMountPoint,
	}
	Logc(ctx).WithFields(logFields).Debug(">>>> filesystem.expandFilesystem")
	defer Logc(ctx).WithFields(logFields).Debug("<<<< filesystem.expandFilesystem")

	preExpandSize, err := getFilesystemSize(ctx, tmpMountPoint)
	if err != nil {
		return 0, err
	}
	_, err = execCommand(ctx, cmd, cmdArguments)
	if err != nil {
		Logc(ctx).Errorf("Expanding filesystem failed; %s", err)
		return 0, err
	}

	postExpandSize, err := getFilesystemSize(ctx, tmpMountPoint)
	if err != nil {
		return 0, err
	}

	if postExpandSize == preExpandSize {
		Logc(ctx).Warnf("Failed to expand filesystem; size=%d", postExpandSize)
	}

	return postExpandSize, nil
}
