// # Model
//
// Event model for eBPF events
// Copyright (c) 2025, ENKI, Inc Polvo
package model

import (
	commonModel "github.com/enki-polvo/polvo-logger/model"
)

// # EventCode
//
// EventCode defines the event code type.
type EventCode int

// # Event
//
// Event defines the interface for all event types.
type Event interface{}

const (
	// Event codes
	PROC_CREATE EventCode = iota
	PROC_TERMINATE
	PROC_BASH_READLINE
)

func (e EventCode) String() string {
	switch e {
	case PROC_CREATE:
		return "ProcessCreate"
	case PROC_TERMINATE:
		return "ProcessTerminate"
	case PROC_BASH_READLINE:
		return "BashReadline"
	default:
		return ""
	}
}

/*************************************************************************
* Metadata
*************************************************************************/

// # ProcessCreateMetadata
//
// ProcessCreateMetadata defines the metadata structure for process creation events.
type ProcessCreateMetadata struct {
	// process relation
	PID  int64 `json:"Pid"`  // example: 1234
	PPID int64 `json:"Ppid"` // example: 4
	// user info
	UID      int64  `json:"Uid"`      // example: 1000
	Username string `json:"Username"` // example: "root"
	TGID     int64  `json:"Tgid"`     // example: 1234
	// process info
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
	ENV         string `json:"Env"`         // example: "PATH=/usr/bin:/bin"
	Image       string `json:"Image"`       // example: "/usr/bin/bash"
}

// # ProcessTerminateMetadata
//
// ProcessTerminateMetadata defines the metadata structure for process termination events.
type ProcessTerminateMetadata struct {
	// process relation
	PID int64 `json:"Pid"` // example: 1234
	// process info
	Ret int `json:"Ret"` // example: 0
	// user info
	UID      int64  `json:"Uid"`      // example: 1000
	Username string `json:"Username"` // example: "root"
}

// # BashReadlineMetadata
//
// BashReadlineMetadata defines the metadata structure for bash readline events.
type BashReadlineMetadata struct {
	// process relation
	PID int64 `json:"Pid"` // example: 1234
	// info
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
	// user info
	UID      int64  `json:"Uid"`      // example: 1000
	Username string `json:"Username"` // example: "root"
}

// # ServiceMetadata
//
// ServiceMetadata defines the metadata structure for service events.
type ServiceMetadata struct {
	PID         int64  `json:"Pid"`         // example: 1234
	UID         int64  `json:"Uid"`         // example: 1000
	TTY         string `json:"Tty"`         // example: "pts/0"
	Image       string `json:"Image"`       // example: "/usr/bin/bash"
	Commandline string `json:"Commandline"` // example: "bash rm -rf /tmp"
}

// # TcpConnectMetadata
//
// TcpConnectMetadata defines the metadata structure for TCP connection events.
type TcpConnectMetadata struct {
	PID   int64  `json:"Pid"`   // example: 1234
	DIP   string `json:"Dip"`   // example: "127.0.0.1"
	Dport int64  `json:"Dport"` // example: 80
	SIP   string `json:"Sip"`   // example: "127.0.0.1"
	Sport int64  `json:"Sport"` // example: 80
	Proto int64  `json:"Proto"` // example: 4
}

// # TcpDisconnectMetadata
//
// TcpDisconnectMetadata defines the metadata structure for TCP disconnection events.
type TcpDisConnectMetadata struct {
	PID   int64  `json:"Pid"`   // example: 1234
	DIP   string `json:"Dip"`   // example: "127.0.0.1"
	Dport int64  `json:"Dport"` // example: 80
	SIP   string `json:"Sip"`   // example: "127.0.0.1"
	Sport int64  `json:"Sport"` // example: 80
	Proto int64  `json:"Proto"` // example: 4
}

/*************************************************************************
* Event
*************************************************************************/

// # ProcessCreateEvent
//
// ProcessCreateEvent defines the event structure for process creation events.
type ProcessCreateEvent struct {
	commonModel.CommonHeader
	Metadata ProcessCreateMetadata `json:"metadata"`
}

// # ProcessTerminateEvent
//
// ProcessTerminateEvent defines the event structure for process termination events.
type ProcessTerminateEvent struct {
	commonModel.CommonHeader
	Metadata ProcessTerminateMetadata `json:"metadata"`
}

// # BashReadlineEvent
//
// BashReadlineEvent defines the event structure for bash readline events.
type BashReadlineEvent struct {
	commonModel.CommonHeader
	Metadata BashReadlineMetadata `json:"metadata"`
}

// # ServiceEvent
//
// ServiceEvent defines the event structure for service events.
type ServiceEvent struct {
	commonModel.CommonHeader
	Metadata ServiceMetadata `json:"metadata"`
}

// # TcpConnectEvent
//
// TcpConnectEvent defines the event structure for TCP connection events.
type TcpConnectEvent struct {
	commonModel.CommonHeader
	Metadata TcpConnectMetadata `json:"metadata"`
}

// # TcpDisConnectEvent
//
// TcpDisConnectEvent defines the event structure for TCP disconnection events.
type TcpDisConnectEvent struct {
	commonModel.CommonHeader
	Metadata TcpDisConnectMetadata `json:"metadata"`
}
