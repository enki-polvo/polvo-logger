package main

import (
	"fmt"
	"github.com/enki-polvo/polvo-logger/logger"
)

func main() {
	logMsg := polvo.Log("openat", 1234, 2345, "asdf...xyz")
	fmt.Println(logMsg)
}
