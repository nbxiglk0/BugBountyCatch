package naabu

import (
	"github.com/projectdiscovery/naabu/v2/pkg/result"
	"github.com/projectdiscovery/naabu/v2/pkg/runner"
	"log"
)

func ExecuteNaabu(hostFile string, outPutFile string) {
	options := runner.Options{
		ResumeCfg:  &runner.ResumeCfg{},
		Retries:    1,
		Silent:     true,
		Threads:    20,
		ExcludeCDN: true,
		TopPorts:   "1000",
		Output:     outPutFile,
		HostsFile:  hostFile,
		OnResult: func(hr *result.HostResult) {
			log.Println(hr.Host, hr.Ports)
		},
	}

	naabuRunner, err := runner.NewRunner(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer naabuRunner.Close()
	naabuRunner.RunEnumeration()
}
