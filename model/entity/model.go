// model/entity/model.go
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

// EntityTypeToString converts an EntityType to its string representation.
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

// CommonEntityModel defines the structure for all entity types.
type CommonEntityModel struct {
	EntityType     EntityType  `json:"EntityType"`               // example: 0
	State          state.State `json:"State"`                    // example: 0
	MatchedRuleIDs string      `json:"MatchedRuleIDs,omitempty"` // example: "rule1"
}

// ProcessEntityModel defines the structure for process entities.
type ProcessEntityModel struct {
	CommonEntityModel
}

// NetworkEntityModel defines the structure for network entities.
type NetworkEntityModel struct {
	CommonEntityModel
	// TODO:
	// NumRecvOps   int64 `json:"NumRecvOps"`   // example: 100 (Number of Receive operations)
	// NumSentOps   int64 `json:"NumSentOps"`   // example: 100 (Number of Send operations)
	// NumRecvBytes int64 `json:"NumRecvBytes"` // example: 100 (Number of bytes received)
	// NumSentBytes int64 `json:"NumSentBytes"` // example: 100 (Number of bytes sent)
}

// FileEntityModel defines the structure for file entities.
// TODO: make file entity(trace) in polvo-architecture
// type FileEntityModel struct {
// 	CommonEntityModel
// 	NumReadOps    int64 `json:"NumReadOps"`    // example: 100 (Number of Read operations)
// 	NumWriteOps   int64 `json:"NumWriteOps"`   // example: 100 (Number of Write operations)
// 	NumReadBytes  int64 `json:"NumReadBytes"`  // example: 100 (Number of bytes read)
// 	NumWriteBytes int64 `json:"NumWriteBytes"` // example: 100 (Number of bytes written)
// }
