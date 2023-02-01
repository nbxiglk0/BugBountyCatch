package main

import (
	"BugBountyCatch/src"
	"BugBountyCatch/src/config"
)

func main() {
	config.ParseConfig("watsons.com.ph")
	src.CatchRunning("watsons.com.ph")
}
