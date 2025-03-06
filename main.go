package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/enki-polvo/polvo-logger/logger"
)

func main() {
	// Build metadata directly as a map.
	metadata := map[string]interface{}{
		"pid":          1234,
		"uid":          2345,
		"process_name": "myapp",
		"user_name":    "alice",
	}

	// Log an event from an eBPF source.
	logMsg, err := logger.Log("eBPF", "openat", "File descriptor opened successfully", metadata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Toggle pretty print. Set to false to print in a single line.
	pretty := true
	if pretty {
		var prettyJSON bytes.Buffer
		err := json.Indent(&prettyJSON, []byte(logMsg), "", "  ")
		if err != nil {
			fmt.Println("Error pretty printing JSON:", err)
			fmt.Println("Log message:", logMsg)
		} else {
			fmt.Println("Log message (pretty printed):")
			fmt.Println(prettyJSON.String())
		}
	} else {
		fmt.Println("Log message:", logMsg)
	}

	// Directly print a network event.
	err = logger.PrintLog("libpcap", "connect", "Established connection to 192.168.1.100:80", map[string]interface{}{
		"ip":   "192.168.1.100",
		"port": 80,
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
