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
	FILE_EVENT
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
	case FILE_EVENT:
		return "FileEvent"
	default:
		return ""
	}
}

// CommonHeader defines the common header structure for all events.
type CommonHeader struct {
	EventCode EventCode `json:"-"`         // example: 1
	EventName string    `json:"EventName"` // example: "ProcessCreate"
	Source    string    `json:"Source"`    // example: "eBPF"
	Timestamp time.Time `json:"Timestamp"` // example: "2023-10-01T12:00:00Z"
}

// CommonModel defines the common structure for all events and entity.
type CommonModel struct {
	CommonHeader
	Metadata any `json:"Metadata"`
}
