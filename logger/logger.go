package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// BuildLog constructs the log message in the unified JSON format.
// The resulting format is:
//
//	{
//	  "eventname": "<eventName>",
//	  "source": "<source>",
//	  "timestamp": "<timestamp>",
//	  "log": "<eventLog>",
//	  "metadata": {<metadata JSON>}
//	}
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

	// Build the unified log message in JSON format.
	logMsg := fmt.Sprintf("{\"eventname\": \"%s\",\n \"source\": \"%s\",\n \"timestamp\": \"%s\",\n \"log\": \"%s\",\n \"metadata\": %s\n}",
		eventName, source, timestamp, eventLog, metadataStr)
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
