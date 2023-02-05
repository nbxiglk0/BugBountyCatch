package main

import (
	//"bytes"
	//"fmt"
	//"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	//"io"
	//"log"
	"bytes"
	"fmt"
	"io"
	"log"

	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

func main() {
	//config.ParseConfig("watsons.com.ph")
	//src.CatchRunning()
	////var crawlUrls []string
	////f, _ := os.Open("/mnt/e/Github/BugBountyCatch/watsons.com.ph/watsons.com.ph_HttpxScan.txt")
	////r := bufio.NewScanner(f)
	////r.Split(bufio.ScanLines)
	////for r.Scan() {
	////	url := strings.Split(r.Text(), "")[0]
	////	crawlUrls = append(crawlUrls, url)
	////}
	////katana.Executable(crawlUrls, "/mnt/e/Github/BugBountyCatch/watsons.com.ph/res.txt")
	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            10,                       // Thread controls the number of threads to use for active enumerations
		Timeout:            30,                       // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10,                       // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		Resolvers:          resolve.DefaultResolvers, // Use the default list of resolvers by marshaling it to the config
		ResultCallback: func(s *resolve.HostEntry) { // Callback function to execute for available host
			log.Println(s.Host, s.Source)
		},
	})

	buf := bytes.Buffer{}
	err = runnerInstance.EnumerateSingleDomain("projectdiscovery.io", []io.Writer{&buf})
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", data)
}
