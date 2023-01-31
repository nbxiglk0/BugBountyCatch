package shuffledns

import (
	"BugBountyCatch/src/config"
	"BugBountyCatch/src/moudle/logger"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/shuffledns/pkg/runner"
)

func ExecuteShuffledns(domain string) {
	// Parse the command line flags and read config files
	options := &runner.Options{}
	options.Domain = domain
	options.ResolversFile = config.GetResolversFilePath()
	options.Wordlist = config.GetSubdomainList()
	massdnsRunner, err := runner.New(options)
	if err != nil {
		logger.Logging(err.Error())
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}
	massdnsRunner.RunEnumeration()
	massdnsRunner.Close()
}
