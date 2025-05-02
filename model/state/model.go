// model/state/model.go
package model

// State defines the state of an entities and events.
type State int

// TODO: Renaming state constants to deliver its intentions better later
const (
	// CREATED indicates the entity was recently created
	CREATED State = iota

	// MODIFIED indicates that some attributes of the entity were modified.
	MODIFIED

	// REUP indicates that the entity already existed in system, but needs to be exported.
	REUP
)
