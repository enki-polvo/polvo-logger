package main

import (
	"fmt"

	"github.com/enki-polvo/polvo-logger/logger"
)

func main() {
	pid := 1234
	uid := 2345

	// Get the formatted log string
	logMsg, err := polvo.BuildLog("openat", &pid, &uid, "asdf...xyz")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("BuildLog output:", logMsg)

	// Directly print the log message to the console
	err = polvo.PrintLog("openat", &pid, &uid, "asdf...xyz")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
