// # Model
//
// Entity model for Traces
// Copyright (c) 2025, ENKI, Inc Polvo
package model

import (
	entityState "github.com/enki-polvo/polvo-logger/model/state"
)

// # ProcessEntity
//
// ProcessEntity defines the process entity structure.
type ProcessEntity struct {
	PID   int64             `json:"Pid"`   // example: 1234
	PPID  int64             `json:"Ppid"`  // example: 4
	CMD   string            `json:"Cmd"`   // example: "bash"
	State entityState.State `json:"State"` // example: 0
}

// # ShellEntity
//
// ShellEntity defines the shell entity structure.
type ShellEntity struct {
	PID         int64             `json:"Pid"`         // example: 1234
	Commandline string            `json:"Commandline"` // example: "bash rm -rf /tmp"
	UID         int64             `json:"Uid"`         // example: 1000
	TTY         string            `json:"Tty"`         // example: "/dev/pts/0"
	Username    string            `json:"Username"`    // example: "root"
	State       entityState.State `json:"State"`       // example: 0
}
