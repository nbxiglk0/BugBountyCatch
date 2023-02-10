package main

import (
	"BugBountyCatch/src"
	"BugBountyCatch/src/Catchconfig"
)

func main() {
	Catchconfig.ParseConfig("watsons.com.ph")
	src.CatchRunning()
}
