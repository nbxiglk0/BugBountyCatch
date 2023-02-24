package main

import (
	"BugBountyCatch/src"
	"BugBountyCatch/src/Catchconfig"
	"flag"
	"github.com/fatih/color"
	"os"
)

func main() {
	target := flag.String("domain", "", "Input your target domain ./BugBountyCatch -domain reacted.com")
	flag.Parse()
	if *target == "" {
		color.Red("Input your target domain ./BugBountyCatch -domain reacted.com")
		os.Exit(0)
	}
	Catchconfig.ParseConfig(*target)
	src.CatchRunning()
	//domain := []string{"myaccount.findmespot.com"}
	//path := "./gau.txt"
	//gau.Executable(domain, path)
}
