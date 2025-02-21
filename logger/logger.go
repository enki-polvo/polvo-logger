package polvo

import (
	"errors"
	"fmt"
	"time"
)

// Creates a formatted log message.
// PID and UID are optional depending on the POLVO sensor's situation, so they can be nil.
func Log(eventName string, pid *int, uid *int, eventLog string) (string, error) {
	// Validate the values
	if eventName == "" {
		return "", errors.New("eventName cannot be empty")
	}

	if pid != nil && *pid <= 0 {
		return "", errors.New("pid must be greater than 0")
	}

	if uid != nil && *uid <= 0 {
		return "", errors.New("uid must be greater than 0")
	}

	// Use placeholders if PID or UID are nil
	var pidStr, uidStr string
	if pid != nil {
		pidStr = fmt.Sprintf("%d", *pid)
	} else {
		pidStr = "-"
	}

	if uid != nil {
		uidStr = fmt.Sprintf("%d", *uid)
	} else {
		uidStr = "-"
	}

	// Format the current time
	currentTime := time.Now().Format(time.RFC3339)

	// Create the formatted log message
	logMsg := fmt.Sprintf("[%s] [%s] (pid=%s, uid=%s) => %s", currentTime, eventName, pidStr, uidStr, eventLog)
	return logMsg, nil
}
