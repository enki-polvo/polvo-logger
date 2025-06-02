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
	// process relation
	PID  int64 `json:"PID"`  // example: 1234
	PPID int64 `json:"PPID"` // example: 4

	// user info
	UID      int64  `json:"UID"`      // example: 1000
	Username string `json:"Username"` // example: "root"
	TGID     int64  `json:"TGID"`     // example: 1234

	// process info
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
	ENV         string `json:"ENV"`         // example: "PATH=/usr/bin:/bin"
	Image       string `json:"Image"`       // example: "/usr/bin/bash"
}

// ProcessTerminateMetadata defines the Metadata structure for process termination events.
type ProcessTerminateMetadata struct {
	// process relation
	PID int64 `json:"PID"` // example: 1234

	// process info
	Ret int64 `json:"Ret"` // example: 0

	// user info
	UID      int64  `json:"UID"`      // example: 1000
	Username string `json:"Username"` // example: "root"
}

// BashReadlineMetadata defines the Metadata structure for bash readline events.
type BashReadlineMetadata struct {
	// process relation
	PID int64 `json:"PID"` // example: 1234

	// info
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"

	// user info
	UID      int64  `json:"UID"`      // example: 1000
	Username string `json:"Username"` // example: "root"
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
	// process  relation
	PID int64 `json:"PID"` // example: 1234
	// tcp info
	Daddr    string `json:"Daddr"`    // example: "127.0.0.1"
	Dport    int64  `json:"Dport"`    // example: 80
	Saddr    string `json:"Saddr"`    // example: "127.0.0.1"
	Sport    int64  `json:"Sport"`    // example: 80
	Protocol int64  `json:"Protocol"` // example: 4
	// tcp operation
	Op state.TcpOp `json:"Op"` // example: "CONNECT" "DISCONNECT" "ACCEPT" etc..
}

// FileMetadata defines the Metadata structure for file events.
type FileMetadata struct {
	// process relation
	PID int64 `json:"PID"` // example: 1234
	UID int64 `json:"UID"` // example: 1000
	// file info
	TargetFilename string `json:"TargetFilename"` // example: "/tmp/file.txt"
	// file operation
	Op   state.FileOp `json:"Op"`   // example: "READ" "RENAME" "WRITE" etc..
	Mode uint64       `json:"Mode"` // example: 0
}

// --------------------------------------------------
// System events Metadata
// --------------------------------------------------

// ProcessCreateEvent defines the event structure for process creation events.
type ProcessCreateEvent struct {
	commonModel.CommonHeader
	Metadata ProcessCreateMetadata `json:"Metadata"`
}

// ProcessTerminateEvent defines the event structure for process termination events.
type ProcessTerminateEvent struct {
	commonModel.CommonHeader
	Metadata ProcessTerminateMetadata `json:"Metadata"`
}

// BashReadlineEvent defines the event structure for bash readline events.
type BashReadlineEvent struct {
	commonModel.CommonHeader
	Metadata BashReadlineMetadata `json:"Metadata"`
}

// ServiceEvent defines the event structure for service events.
type ServiceEvent struct {
	commonModel.CommonHeader
	Metadata ServiceMetadata `json:"Metadata"`
}

// TcpEvent defines the event structure for TCP events.
type TcpEvent struct {
	commonModel.CommonHeader
	Metadata TcpMetadata `json:"Metadata"`
}

// FileEvent defines the event structure for file events.
type FileEvent struct {
	commonModel.CommonHeader
	Metadata FileMetadata `json:"Metadata"`
}
