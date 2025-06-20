// model/commonModel.go
package commonModel

import (
	"time"
)

// EventCode defines the event code type.
type EventCode int

// Event codes
const (
	PROC_CREATE EventCode = iota
	PROC_TERMINATE
	PROC_BASH_READLINE
	PROC_SERVICE
	TCP_EVENT
	FILE_OPEN_EVENT
	FILE_RENAME_EVENT
)

// EventCodeToString converts an EventCode to its string representation.
func (e EventCode) String() string {
	switch e {
	case PROC_CREATE:
		return "ProcessCreate"
	case PROC_TERMINATE:
		return "ProcessTerminate"
	case PROC_BASH_READLINE:
		return "BashReadline"
	case PROC_SERVICE:
		return "Service"
	case TCP_EVENT:
		return "TcpEvent"
	case FILE_OPEN_EVENT:
		return "FileOpenEvent"
	case FILE_RENAME_EVENT:
		return "FileRenameEvent"
	default:
		return ""
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
