package main

import (
	"BugBountyCatch/src"
	"BugBountyCatch/src/config"
)

func main() {
	config.ParseConfig()
	src.CatchRunning("test.com")
}
