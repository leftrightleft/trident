// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// NetworkIPInterfaceModifyReader is a Reader for the NetworkIPInterfaceModify structure.
type NetworkIPInterfaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPInterfaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPInterfaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPInterfaceModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPInterfaceModifyOK creates a NetworkIPInterfaceModifyOK with default headers values
func NewNetworkIPInterfaceModifyOK() *NetworkIPInterfaceModifyOK {
	return &NetworkIPInterfaceModifyOK{}
}

/* NetworkIPInterfaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPInterfaceModifyOK struct {
}

func (o *NetworkIPInterfaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/interfaces/{uuid}][%d] networkIpInterfaceModifyOK ", 200)
}

func (o *NetworkIPInterfaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkIPInterfaceModifyDefault creates a NetworkIPInterfaceModifyDefault with default headers values
func NewNetworkIPInterfaceModifyDefault(code int) *NetworkIPInterfaceModifyDefault {
	return &NetworkIPInterfaceModifyDefault{
		_statusCode: code,
	}
}

/* NetworkIPInterfaceModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 1376663 | Cannot add interface to DNS zone because all interfaces from a single DNS zone must be in the same SVM. |
| 1376963 | Duplicate IP address. |
| 1376997 | Interface failed to migrate because the node hosting the port is not healthy. |
| 1376998 | The specified location.node does not own any ports in the same broadcast domain as the home port of the interface. |
| 1376999 | Interface failed to migrate because port is in the down admin state. |
| 1377607 | The specified location.port is not in the same broadcast domain as the home port of the interface. |
| 1966138 | The same IP address may not be used for both a mgmt interface and a gateway address. |
| 1966141 | Invalid DNS zone name. |
| 1966142 | Only data LIFs can be assigned a DNS zone. |
| 1966197 | Migration of cluster interfaces must be done from the local node. |
| 1966267 | IPv6 addresses must have a prefix length between 1 and 127. |
| 1966269 | IPv4 addresses must have a prefix length between 1 and 32. |
| 1966476 | DNS Update is supported only on data interfaces. |
| 1966477 | DNS Update is supported only on interfaces configured with the NFS or CIFS protocol. |
| 1967106 | The specified location.home_port.name does not match the specified port name of location.home_port.uuid. |
| 1967107 | The specified location.home_port.uuid is not valid. |
| 1967111 | A home node must be specified by at least one location.home_node, location.home_port, or location.broadcast_domain field. |
| 1967113 | The specified location.port.name does not match the port name of location.port.uuid. |
| 1967114 | The specified location.port.uuid is not valid. |
| 1967115 | The specified location.node.name does not match the node name of location.node.uuid. |
| 1967116 | The specified location.port.node.name does not match the node name of location.node.uuid. |
| 1967117 | The specified location.port.node.name does not match location.node.name. |
| 1967118 | A node must be specified by at least one location.node or location.port field. |
| 1967119 | The specified location.node.name does not match the node name of location.port.uuid. |
| 1967120 | The specified service_policy.name does not match the specified service policy name of service_policy.uuid. |
| 1967121 | The specified service_policy.uuid is not valid. |
| 1967125 | You cannot patch the "location.node" or "location.port" fields to migrate interfaces using the iSCSI data protocol. Instead perform the following PATCH operations on the interface: set the "enabled" field to "false"; change one or more "location.home_port" fields to migrate the interface; and then set the "enabled" field to "true". |
| 1967129 | The specified location.home_port.uuid is not valid. |
| 1967130 | The specified location.home_port.name is not valid. |
| 1967131 | The specified location.home_port.uuid and location.home_port.name are not valid. |
| 1967132 | The specified location.port.uuid is not valid. |
| 1967133 | The specified location.port.name is not valid. |
| 1967134 | The specified location.port.uuid and location.port.name are not valid. |
| 1967138 | Cannot patch port for a VIP interface. The specified parameter location.port.uuid is not valid. |
| 1967139 | Cannot patch port for a VIP interface. The specified parameter location.port.name is not valid. |
| 1967140 | Cannot patch port for a VIP interface. The specified parameters location.port.uuid and location.port.name are not valid. |
| 1967141 | Cannot patch home_port for a VIP interface. The specified parameter location.home_port.uuid is not valid. |
| 1967142 | Cannot patch home_port for a VIP interface. The specified parameter location.home_port.name is not valid. |
| 1967143 | Cannot patch home_port for a VIP interface. The specified parameters location.home_port.uuid and location.home_port.name  are not valid. |
| 1967145 | The specified location.failover is not valid. |
| 1967153 | No suitable port exists on location.home_node to host the interface. |
| 1967380 | Cannot patch home_port for a VIP interface. The specified parameter location.home_port.node.name is not valid. Consider using location.home_node.name instead. |
| 1967386 | Cannot patch port for a VIP interface. The specified parameter location.port.node.name is not valid. Consider using location.node.name instead. |
| 1967387 | The specified IP address is in use by a subnet in this IPspace. |
| 1967389 | Patching location.is_home to the value "false" is not supported. The value "true" would revert a network interface to its home port if the current value is "false". |
| 1967390 | Cannot patch a LIF revert as it requires an effective cluster version of 9.9.1 or later. |
| 1967391 | Patching the DNS zone requires an effective cluster version of 9.9.1 or later. |
| 1967392 | Patching the DDNS enable parameter requires an effective cluster version of 9.9.1 or later. |
| 53281065 | The service_policy does not exist in the SVM. |
| 53281086 | LIF would exceed the maximum number of supported intercluster LIFs in IPspace. |
| 53281089 | LIF on SVM cannot be updated to use service policy because that service policy includes SAN services and the target LIF is not home. |

*/
type NetworkIPInterfaceModifyDefault struct {
	_statusCode int
}

// Code gets the status code for the network ip interface modify default response
func (o *NetworkIPInterfaceModifyDefault) Code() int {
	return o._statusCode
}

func (o *NetworkIPInterfaceModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /network/ip/interfaces/{uuid}][%d] network_ip_interface_modify default ", o._statusCode)
}

func (o *NetworkIPInterfaceModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
