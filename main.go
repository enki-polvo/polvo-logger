package main

import (
	"fmt"

	"github.com/enki-polvo/polvo-logger/logger"
)

func main() {
	// Example using ProcessInfo helper.
	executionContent := logger.ExecutionContext{
		PID:         1234,
		UID:         2345,
		ProcessName: "myapp",
		UserName:    "alice",
	}
	// Convert ProcessInfo to metadata map.
	metadata := executionContent.ToMap()

	// Log an event from an eBPF source.
	logMsg, err := logger.Log("eBPF", "openat", "File descriptor opened successfully", metadata)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Log message:", logMsg)

	// Directly print a network event.
	err = logger.PrintLog("libpcap", "connect", "Established connection to 192.168.1.100:80", map[string]interface{}{
		"ip":   "192.168.1.100",
		"port": 80,
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
