// # Model
//
// Entity model for Traces
// Copyright (c) 2025, ENKI, Inc Polvo
package model

import (
	state "github.com/enki-polvo/polvo-logger/model/state"
)

type EntityType int

const (
	PROCESS_ENTITY EntityType = iota
	NETWORK_ENTITY
	FILE_ENTITY
)

func (e EntityType) String() string {
	switch e {
	case PROCESS_ENTITY:
		return "PROCESS"
	case NETWORK_ENTITY:
		return "NETWORK"
	case FILE_ENTITY:
		return "FILE"
	default:
		return ""
	}
}

// # CommonEntityModel
//
// CommonEntityModel defines the structure for all entity types.
type CommonEntityModel struct {
	EntityType     EntityType  `json:"EntityType"`               // example: 0
	State          state.State `json:"State"`                    // example: 0
	MatchedRuleIDs string      `json:"MatchedRuleIDs,omitempty"` // example: "rule1"
}

// # ProcessEntityModel
//
// ProcessEntityModel defines the structure for process entities.
type ProcessEntityModel struct {
	CommonEntityModel
}

// # NetworkEntityModel
//
// NetworkEntityModel defines the structure for network entities.
type NetworkEntityModel struct {
	CommonEntityModel
	NumRRecvOps   int64 `json:"NumRRecvOps"`   // example: 100
	NumWSentOps   int64 `json:"NumWSentOps"`   // example: 100
	NumRRecvBytes int64 `json:"NumRRecvBytes"` // example: 100
	NumWSentBytes int64 `json:"NumWSentBytes"` // example: 100
}
