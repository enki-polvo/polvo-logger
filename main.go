package main

import (
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

	// Test 1: Using an empty timestamp string (auto-fills current time)
	emptyTimestamp := ""
	fmt.Println("Test 1: Logger with empty timestamp (auto current time):")
	logger.PrintLogPretty("eBPF", "openat", "File descriptor opened successfully", emptyTimestamp, metadata)

	// Test 2: Using a valid timestamp string
	// Note: This valid timestamp must match the layout "2006-01-02T15:04:05.000000Z07:00".
	customTimestamp := "2004-09-26T12:34:56.123456+00:00"
	fmt.Println("\nTest 2: Logger with valid timestamp:")
	logger.PrintLog("libpcap", "connect", "Established connection to 192.168.1.100:80", customTimestamp, map[string]interface{}{
		"ip":   "192.168.1.100",
		"port": 80,
	})
}
