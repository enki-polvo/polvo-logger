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
// The 'timestamp' parameter is optional. If it's an empty string, the current time is used.
// Otherwise, it will be accepted if it matches one of the allowed layouts.
func BuildLog(source, eventName, eventLog, timestamp string, metadata map[string]interface{}) (*LogMessage, error) {
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

	// Define multiple acceptable layouts.
	layouts := []string{
		"2006-01-02T15:04:05Z07:00",           // RFC3339
		"2006-01-02T15:04:05.000000Z07:00",    // microsecond precision
		"2006-01-02T15:04:05.999999999Z07:00", // nanosecond precision
	}

	// If no timestamp provided, use current time.
	if timestamp == "" {
		timestamp = time.Now().Format(layouts[0])
	} else {
		valid := false
		for _, layout := range layouts {
			if _, err := time.Parse(layout, timestamp); err == nil {
				valid = true
				break
			}
		}
		if !valid {
			return nil, errors.New("invalid timestamp format")
		}
	}

	return &LogMessage{
		EventName: eventName,
		Source:    source,
		Timestamp: timestamp,
		Log:       eventLog,
		Metadata:  metadata,
	}, nil
}

// PrintLog prints the unified log message as a one-line JSON string.
func PrintLog(source, eventName, eventLog, timestamp string, metadata map[string]interface{}) {
	logMsg, err := BuildLog(source, eventName, eventLog, timestamp, metadata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b, err := json.Marshal(logMsg)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(b))
}

// PrintLogPretty prints the unified log message as a pretty-printed JSON.
func PrintLogPretty(source, eventName, eventLog, timestamp string, metadata map[string]interface{}) {
	logMsg, err := BuildLog(source, eventName, eventLog, timestamp, metadata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	b, err := json.MarshalIndent(logMsg, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(b))
}
