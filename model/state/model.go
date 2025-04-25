// # Model
//
// State model for entities and events.
// Copyright (c) 2025, ENKI, Inc Polvo
package model

// State defines the state of an entitys and events.
type State int

const (
	// CREATED indicates the entity was recently created
	CREATED State = iota
	// MODIFIED indicates that some attributes of the entity were modified.
	MODIFIED
	// REUP indicates that the entity already existed in system, but needs to be exported.
	REUP
)
