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

// FileOpenPurposeOp defines the purposes of a specific file open operation
type FileOpenPurposeOp int

const (
	// default value for FileOp
	FILE_OPEN_TO_UNSET FileOpenPurposeOp = iota
	// File opened to read data from it
	FILE_OPEN_TO_READ
	// File opened to write data to it
	FILE_OPEN_TO_WRITE
	// File opened to do something else that is not belonged to read or write
	FILE_OPEN_TO_OTHER
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
