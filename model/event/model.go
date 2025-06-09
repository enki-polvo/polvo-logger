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
	PID         int64  `json:"PID" mapstructure:"PID"`                 // example: 1234
	PPID        int64  `json:"PPID" mapstructure:"PPID"`               // example: 4
	UID         int64  `json:"UID" mapstructure:"UID"`                 // example: 1000
	Username    string `json:"Username" mapstructure:"Username"`       // example: "root"
	TGID        int64  `json:"TGID" mapstructure:"TGID"`               // example: 1234
	Commandline string `json:"Commandline" mapstructure:"Commandline"` // example: "bash rm -rf /tmp"
	ENV         string `json:"ENV" mapstructure:"ENV"`                 // example: "PATH=/usr/bin:/bin"
	Image       string `json:"Image" mapstructure:"Image"`             // example: "/usr/bin/bash"
}

// ProcessTerminateMetadata defines the Metadata structure for process termination events.
type ProcessTerminateMetadata struct {
	PID      int64  `json:"PID" mapstructure:"PID"`           // example: 1234
	Ret      int64  `json:"Ret" mapstructure:"Ret"`           // example: 0
	UID      int64  `json:"UID" mapstructure:"UID"`           // example: 1000
	Username string `json:"Username" mapstructure:"Username"` // example: "root"
}

// BashReadlineMetadata defines the Metadata structure for bash readline events.
type BashReadlineMetadata struct {
	PID         int64  `json:"PID" mapstructure:"PID"`                 // example: 1234
	Commandline string `json:"Commandline" mapstructure:"Commandline"` // example: "bash rm -rf /tmp"
	UID         int64  `json:"UID" mapstructure:"UID"`                 // example: 1000
	Username    string `json:"Username" mapstructure:"Username"`       // example: "root"
}

// ServiceMetadata defines the Metadata structure for service events.
type ServiceMetadata struct {
	PID         int64  `json:"PID" mapstructure:"PID"`                 // example: 1234
	UID         int64  `json:"UID" mapstructure:"UID"`                 // example: 1000
	TTY         string `json:"TTY" mapstructure:"TTY"`                 // example: "pts/0"
	Image       string `json:"Image" mapstructure:"Image"`             // example: "/usr/bin/bash"
	Commandline string `json:"Commandline" mapstructure:"Commandline"` // example: "bash rm -rf /tmp"
}

// TcpMetadata defines the Metadata structure for TCP events.
type TcpMetadata struct {
	PID      int64       `json:"PID" mapstructure:"PID"`           // example: 1234
	Daddr    string      `json:"Daddr" mapstructure:"Daddr"`       // example: "127.0.0.1"
	Dport    int64       `json:"Dport" mapstructure:"Dport"`       // example: 80
	Saddr    string      `json:"Saddr" mapstructure:"Saddr"`       // example: "127.0.0.1"
	Sport    int64       `json:"Sport" mapstructure:"Sport"`       // example: 80
	Protocol int64       `json:"Protocol" mapstructure:"Protocol"` // example: 4
	Op       state.TcpOp `json:"Op" mapstructure:"Op"`             // example: "CONNECT" "DISCONNECT" "ACCEPT" etc..
}

// FileMetadata defines the Metadata structure for file events.
type FileMetadata struct {
	PID            int64        `json:"PID" mapstructure:"PID"`                       // example: 1234
	UID            int64        `json:"UID" mapstructure:"UID"`                       // example: 1000
	TargetFilename string       `json:"TargetFilename" mapstructure:"TargetFilename"` // example: "/tmp/file.txt"
	Op             state.FileOp `json:"Op" mapstructure:"Op"`                         // example: "READ" "RENAME" "WRITE" etc..
	Mode           uint64       `json:"Mode" mapstructure:"Mode"`                     // example: 0
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
