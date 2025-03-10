package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// LogMessage defines the unified log message structure.
type LogMessage struct {
	EventName string                 `json:"eventname"`
	Source    string                 `json:"source"`
	Timestamp string                 `json:"timestamp"`
	Log       string                 `json:"log"`
	Metadata  map[string]interface{} `json:"metadata"`
}

// BuildLog constructs the log message using a structured type.
func BuildLog(source, eventName, eventLog string, metadata map[string]interface{}) (*LogMessage, error) {
	// Validate required fields.
	if source == "" {
		return nil, errors.New("source cannot be empty")
	}
	if eventName == "" {
		return nil, errors.New("eventName cannot be empty")
	}
	if eventLog == "" {
		return nil, errors.New("eventLog cannot be empty")
	}

	return &LogMessage{
		EventName: eventName,
		Source:    source,
		Timestamp: time.Now().Format(time.RFC3339),
		Log:       eventLog,
		Metadata:  metadata,
	}, nil
}

// PrintLog prints the unified log message as a one-line JSON string.
func PrintLog(source, eventName, eventLog string, metadata map[string]interface{}) {
	logMsg, err := BuildLog(source, eventName, eventLog, metadata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Marshal without indent (one-line)
	b, err := json.Marshal(logMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(b))
}

// PrintLogPretty prints the unified log message as a pretty-printed JSON.
func PrintLogPretty(source, eventName, eventLog string, metadata map[string]interface{}) {
	logMsg, err := BuildLog(source, eventName, eventLog, metadata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Marshal with indentation for pretty printing
	b, err := json.MarshalIndent(logMsg, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(b))
}
