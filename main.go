package main

import (
	"fmt"
	"github.com/enki-polvo/polvo-logger/logger"
)

func main() {
	pid := 123
	uid := 4567

	logMsg, err := polvo.Log("sensor-example", &pid, &uid, "This is a test log message.")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(logMsg)
}
