package subfinder

import (
	"BugBountyCatch/src/Catchconfig"
	"BugBountyCatch/src/moudle/logger"
	"bytes"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
	"io"
	"log"
)

func Executable(domain string) []byte {
	subfinderConfig := Catchconfig.InitConfig.SubfinderConfig
	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            subfinderConfig.Threads,        // Thread controls the number of threads to use for active enumerations
		Timeout:            30,                             // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10,                             // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		Resolvers:          Catchconfig.GetResolversList(), // Use the default list of resolvers by marshaling it to the Catchconfig
		ResultCallback: func(s *resolve.HostEntry) { // Callback function to execute for available host
			log.Println(s.Host, s.Source)
		},
	})
	buf := bytes.Buffer{}
	err = runnerInstance.EnumerateSingleDomain(domain, []io.Writer{&buf})
	if err != nil {
		logger.Logging(err.Error())
		log.Fatal(err)

	}
	data, err := io.ReadAll(&buf)
	if err != nil {
		logger.Logging(err.Error())
		log.Fatal(err)
	}
	return data
}
