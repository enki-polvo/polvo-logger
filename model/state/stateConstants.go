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

func (s State) String() string {
	switch s {
	case CREATED:
		return "CREATED"
	case MODIFIED:
		return "MODIFIED"
	case REUP:
		return "REUP"
	default:
		return ""
	}
}

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

func (f FileOpenPurposeOp) String() string {
	switch f {
	case FILE_OPEN_TO_UNSET:
		return "FILE_OPEN_TO_UNSET"
	case FILE_OPEN_TO_READ:
		return "FILE_OPEN_TO_READ"
	case FILE_OPEN_TO_WRITE:
		return "FILE_OPEN_TO_WRITE"
	case FILE_OPEN_TO_OTHER:
		return "FILE_OPEN_TO_OTHER"
	default:
		return ""
	}
}

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

func (t TcpOp) String() string {
	switch t {
	case TCP_OP_UNSET:
		return "TCP_OP_UNSET"
	case TCP_CONNECT:
		return "TCP_CONNECT"
	case TCP_DISCONNECT:
		return "TCP_DISCONNECT"
	case TCP_ACCEPT:
		return "TCP_ACCEPT"
	default:
		return ""
	}
}
