package naabu

import (
	"BugBountyCatch/src/config"
	"github.com/projectdiscovery/naabu/v2/pkg/result"
	"github.com/projectdiscovery/naabu/v2/pkg/runner"
	"log"
)

func Executable(hostFile string, outPutFile string) {
	naabuConfig := config.InitConfig.NaabuConfig
	options := runner.Options{
		ResumeCfg: &runner.ResumeCfg{},
		Retries:   1,
		Silent:    naabuConfig.Silent,
		Threads:   naabuConfig.Threads,
		TopPorts:  naabuConfig.TopPorts,
		Output:    outPutFile,
		OutputCDN: naabuConfig.OutputCDN,
		HostsFile: hostFile,
		OnResult: func(hr *result.HostResult) {
		},
	}

	naabuRunner, err := runner.NewRunner(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer naabuRunner.Close()
	naabuRunner.RunEnumeration()
}
