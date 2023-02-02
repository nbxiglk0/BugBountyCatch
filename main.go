package main

import (
	"BugBountyCatch/src/moudle/scan/katana"
	"bufio"
	"os"
	"strings"
)

func main() {
	//config.ParseConfig("watsons.com.ph")
	//src.CatchRunning()
	var crawlUrls []string
	f, _ := os.Open("/mnt/d/Github/BugBountyCatch/watsons.com.ph/watsons.com.ph_HttpxScan.txt")
	r := bufio.NewScanner(f)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		url := strings.Split(r.Text(), "")[0]
		crawlUrls = append(crawlUrls, url)
	}
	katana.Executable(crawlUrls, "/mnt/d/Github/BugBountyCatch/watsons.com.ph/res.txt")
}
