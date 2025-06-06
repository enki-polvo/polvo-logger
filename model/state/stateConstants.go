// model/state/stateConstants.go
package stateConstants

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

// FileOp defines the file operation types.
type FileOp int

const (
	// default value for FileOp
	FILE_OP_UNSET FileOp = iota
	// File creation
	FILE_CREATE
	// File deletion
	FILE_DELETE
	// File read
	FILE_READ
	// File write
	FILE_WRITE
	// File rename
	FILE_RENAME
)

// TcpOp defines the TCP operation types.
type TcpOp int

const (
	// default value for TcpOp
	TCP_OP_UNSET TcpOp = iota
	// TCP connection establishment
	TCP_CONNECT
	// TCP connection termination
	TCP_DISCONNECT
	// TCP connection acceptance
	TCP_ACCEPT
)
