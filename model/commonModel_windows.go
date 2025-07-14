//go:build windows
// +build windows

// model/commonModel.go
package commonModel

import (
	"time"
)

// EventCode defines the event code type.
type EventCode int

// Event codes
const (
	PROCESS_CREATION EventCode = iota
	NETWORK_CONNECTION
	DRIVER_LOAD
	IMAGE_LOAD
	PROCESS_ACCESS
	FILE_EVENT
	REGISTRY_ADD
	REGISTRY_DELETE
	REGISTRY_SET
	REGISTRY_EVENT
	CREATE_STREAM_HASH
	DNS_QUERY
	FILE_DELETE
)

// EventCodeToString converts an EventCode to its string representation.
func (e EventCode) String() string {
	switch e {
	case PROCESS_CREATION:
		return "process_creation"
	case NETWORK_CONNECTION:
		return "network_connection"
	case DRIVER_LOAD:
		return "driver_load"
	case IMAGE_LOAD:
		return "image_load"
	case PROCESS_ACCESS:
		return "process_access"
	case FILE_EVENT:
		return "file_event"
	case REGISTRY_ADD:
		return "registry_add"
	case REGISTRY_DELETE:
		return "registry_delete"
	case REGISTRY_SET:
		return "registry_set"
	case REGISTRY_EVENT:
		return "registry_event"
	case CREATE_STREAM_HASH:
		return "create_stream_hash"
	case DNS_QUERY:
		return "dns_query"
	case FILE_DELETE:
		return "file_delete"
	default:
		return "UnknownEvent"
	}
}

// CommonHeader defines the common header structure for all events.
type CommonHeader struct {
	EventCode EventCode `json:"EventCode"` // example: 1
	EventName string    `json:"EventName"` // example: "ProcessCreate"
	Source    string    `json:"Source"`    // example: "eBPF"
	Timestamp time.Time `json:"Timestamp"` // example: "2023-10-01T12:00:00Z"
}

// CommonModel defines the common structure for all events and entity.
type CommonModel struct {
	CommonHeader
	Metadata any `json:"Metadata"`
}

// CommonModelWrapper is a wrapper for CommonModel that includes a Metadata field as map.
// This is useful for decoding purposes, where the Metadata can be a map of any type.
type CommonModelWrapper struct {
	CommonHeader
	Metadata map[string]any `json:"Metadata"`
}
