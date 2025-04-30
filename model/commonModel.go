package commonModel

import "time"

// ## CommonHeader
//
// CommonHeader defines the common header structure for all events.
type CommonHeader struct {
	EventName string    `json:"Eventname"` // example: "ProcessCreate"
	Source    string    `json:"Source"`    // example: "eBPF"
	Timestamp time.Time `json:"Timestamp"` // example: "2023-10-01T12:00:00Z"
}

// # CommonModel
//
// CommonModel defines the common structure for all events and entity.
type CommonModel struct {
	CommonHeader
	Metadata interface{} `json:"metadata"`
}
