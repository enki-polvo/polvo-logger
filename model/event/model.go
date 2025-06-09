// event/model.go
package eventModel

import (
	commonModel "github.com/enki-polvo/polvo-logger/model"
	state "github.com/enki-polvo/polvo-logger/model/state"
)

// Event defines the interface for all event types.
type Event any

// --------------------------------------------------
// Event Metadata
// --------------------------------------------------

// ProcessCreateMetadata defines the Metadata structure for process creation events.
type ProcessCreateMetadata struct {
	PID         int64  `json:"PID"`         // example: 1234
	PPID        int64  `json:"PPID"`        // example: 4
	UID         int64  `json:"UID"`         // example: 1000
	Username    string `json:"Username"`    // example: "root"
	TGID        int64  `json:"TGID"`        // example: 1234
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
	ENV         string `json:"ENV"`         // example: "PATH=/usr/bin:/bin"
	Image       string `json:"Image"`       // example: "/usr/bin/bash"
}

// ProcessTerminateMetadata defines the Metadata structure for process termination events.
type ProcessTerminateMetadata struct {
	PID      int64  `json:"PID"`      // example: 1234
	Ret      int64  `json:"Ret"`      // example: 0
	UID      int64  `json:"UID"`      // example: 1000
	Username string `json:"Username"` // example: "root"
}

// BashReadlineMetadata defines the Metadata structure for bash readline events.
type BashReadlineMetadata struct {
	PID         int64  `json:"PID"`         // example: 1234
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
	UID         int64  `json:"UID"`         // example: 1000
	Username    string `json:"Username"`    // example: "root"
}

// ServiceMetadata defines the Metadata structure for service events.
type ServiceMetadata struct {
	PID         int64  `json:"PID"`         // example: 1234
	UID         int64  `json:"UID"`         // example: 1000
	TTY         string `json:"TTY"`         // example: "pts/0"
	Image       string `json:"Image"`       // example: "/usr/bin/bash"
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
}

// TcpMetadata defines the Metadata structure for TCP events.
type TcpMetadata struct {
	PID      int64       `json:"PID"`      // example: 1234
	Daddr    string      `json:"Daddr"`    // example: "127.0.0.1"
	Dport    int64       `json:"Dport"`    // example: 80
	Saddr    string      `json:"Saddr"`    // example: "127.0.0.1"
	Sport    int64       `json:"Sport"`    // example: 80
	Protocol int64       `json:"Protocol"` // example: 4
	Op       state.TcpOp `json:"Op"`       // example: "CONNECT" "DISCONNECT" "ACCEPT" etc..
}

// FileOpenMetadata defines the Metadata structure for file open events
// for specific purposes (e.g., file opened to write data to it).
type FileOpenMetadata struct {
	PID               int64                   `json:"PID"`               // example: 8080
	FileOpenerUID     int64                   `json:"FileOpenerUID"`     // example: 1200
	FileOpenerGID     int64                   `json:"FileOpenerGID"`     // example: 1000
	FileOwnerUID      int64                   `json:"FileOwnerUID"`      // example: 1200
	FileOwnerGID      int64                   `json:"FileOwnerGID"`      // example: 1000
	Mode              int64                   `json:"Mode"`              // example: 0444
	Fmode             int64                   `json:"Fmode"`             // example: 0100644
	FileOpenPurposeOp state.FileOpenPurposeOp `json:"FileOperationType"` // example: "FILE_OPEN_TO_WRITE"
	Indoe             int64                   `json:"Indoe"`             // example: 17986650
	Size              int64                   `json:"Size"`              // example: 1048576
	ProcessName       string                  `json:"ProcessName"`       // example: "bash"
	Path              string                  `json:"Path"`              // example: "/var/log/syslog"
}

// FileMetadata defines the Metadata structure for file events.
// It includes file open, close, and rename events.
type FileRenameMetadata struct {
	PID     int64  `json:"PID"`     // example: 8080
	UID     int64  `json:"UID"`     // example: 1200
	GID     int64  `json:"GID"`     // example: 1000
	Command string `json:"Command"` // example: "mv"
	OldPath string `json:"OldPath"` // example: "/var/log/syslog"
	NewPath string `json:"NewPath"` // example: "/var/log/syslog.backup"
}

// --------------------------------------------------
// System events Metadata
//
// They define the event structures for each type of event.
// --------------------------------------------------

type ProcessCreateEvent struct {
	commonModel.CommonHeader
	Metadata ProcessCreateMetadata `json:"Metadata"`
}

type ProcessTerminateEvent struct {
	commonModel.CommonHeader
	Metadata ProcessTerminateMetadata `json:"Metadata"`
}

type BashReadlineEvent struct {
	commonModel.CommonHeader
	Metadata BashReadlineMetadata `json:"Metadata"`
}

type ServiceEvent struct {
	commonModel.CommonHeader
	Metadata ServiceMetadata `json:"Metadata"`
}

type TcpEvent struct {
	commonModel.CommonHeader
	Metadata TcpMetadata `json:"Metadata"`
}

type FileOpenEvent struct {
	commonModel.CommonHeader
	Metadata FileOpenMetadata `json:"Metadata"`
}

type FileRenameEvent struct {
	commonModel.CommonHeader
	Metadata FileRenameMetadata `json:"Metadata"`
}
