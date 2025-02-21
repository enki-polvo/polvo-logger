package polvo

import (
	"fmt"
	"time"
)

// Creates a formatted log message.
func Log(eventName string, pid int, uid int, eventLog string) string {
	currentTime := time.Now().Format(time.RFC3339)
	return fmt.Sprintf("[%s] [%s] (pid=%d, uid=%d) => %s", currentTime, eventName, pid, uid, eventLog)
}
