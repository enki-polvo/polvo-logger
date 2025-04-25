// # Model
//
// Event model for eBPF events
// Copyright (c) 2025, ENKI, Inc Polvo
package model

import "time"

const (
	PROC_CREATE        string = "ProcessCreate"
	PROC_TERMINATE     string = "ProcessTerminate"
	PROC_BASH_READLINE string = "BashReadline"
)

// ## CommonHeader
//
// CommonHeader defines the common header structure for all events.
type CommonHeader struct {
	EventName string    `json:"Eventname"` // example: "ProcessCreate"
	Source    string    `json:"Source"`    // example: "eBPF"
	Timestamp time.Time `json:"Timestamp"` // example: "2023-10-01T12:00:00Z"
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

/*************************************************************************
* Event
*************************************************************************/

// # ProcessCreateEvent
//
// ProcessCreateEvent defines the event structure for process creation events.
type ProcessCreateEvent struct {
	CommonHeader
	Metadata ProcessCreateMetadata `json:"metadata"`
}

// # ProcessTerminateEvent
//
// ProcessTerminateEvent defines the event structure for process termination events.
type ProcessTerminateEvent struct {
	CommonHeader
	Metadata ProcessTerminateMetadata `json:"metadata"`
}

// # BashReadlineEvent
//
// BashReadlineEvent defines the event structure for bash readline events.
type BashReadlineEvent struct {
	CommonHeader
	Metadata BashReadlineMetadata `json:"metadata"`
}
