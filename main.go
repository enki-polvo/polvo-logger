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

	// Toggle pretty print. Set to false to print in a single line.
	pretty := true

	if pretty {
		fmt.Println("Log message (pretty printed):")
		logger.PrintLogPretty("eBPF", "openat", "File descriptor opened successfully", metadata)
	} else {
		fmt.Println("Log message (one-liner):")
		logger.PrintLog("eBPF", "openat", "File descriptor opened successfully", metadata)
	}

	// Directly print a network event in one-line format.
	fmt.Println("Network event log (one-liner):")
	logger.PrintLog("libpcap", "connect", "Established connection to 192.168.1.100:80", map[string]interface{}{
		"ip":   "192.168.1.100",
		"port": 80,
	})
}
