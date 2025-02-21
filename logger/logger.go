package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// ExecutionContext holds common process-related details.
type ExecutionContext struct {
	PID         int    `json:"pid,omitempty"`
	UID         int    `json:"uid,omitempty"`
	ProcessName string `json:"process_name,omitempty"`
	UserName    string `json:"user_name,omitempty"`
}

// ToMap converts ExecutionContext into a metadata map.
// Fields that have a zero value are omitted.
func (ec ExecutionContext) ToMap() map[string]interface{} {
	metadata := make(map[string]interface{})
	if ec.PID > 0 {
		metadata["pid"] = ec.PID
	}
	if ec.UID > 0 {
		metadata["uid"] = ec.UID
	}
	if ec.ProcessName != "" {
		metadata["process_name"] = ec.ProcessName
	}
	if ec.UserName != "" {
		metadata["user_name"] = ec.UserName
	}
	return metadata
}

// BuildLog constructs the log message in the unified format.
// The resulting format is:
// [timestamp] [eventName@source] [eventLog] [metadata]
func BuildLog(source, eventName, eventLog string, metadata map[string]interface{}) (string, error) {
	// Validate required fields.
	if source == "" {
		return "", errors.New("source cannot be empty")
	}
	if eventName == "" {
		return "", errors.New("eventName cannot be empty")
	}
	if eventLog == "" {
		return "", errors.New("eventLog cannot be empty")
	}

	// Generate timestamp.
	timestamp := time.Now().Format(time.RFC3339)

	// Convert metadata to JSON string.
	metadataStr := "{}"
	if metadata != nil {
		bytes, err := json.Marshal(metadata)
		if err != nil {
			return "", fmt.Errorf("failed to marshal metadata: %v", err)
		}
		metadataStr = string(bytes)
	}

	// Build the unified log message.
	logMsg := fmt.Sprintf("[%s@%s] [%s] [%s] [%s]", eventName, source, timestamp, eventLog, metadataStr)
	return logMsg, nil
}

// Log returns the unified log message string based on the provided parameters.
func Log(source, eventName, eventLog string, metadata map[string]interface{}) (string, error) {
	return BuildLog(source, eventName, eventLog, metadata)
}

// PrintLog prints the unified log message to the console.
func PrintLog(source, eventName, eventLog string, metadata map[string]interface{}) error {
	logMsg, err := BuildLog(source, eventName, eventLog, metadata)
	if err != nil {
		return err
	}
	fmt.Println(logMsg)
	return nil
}
