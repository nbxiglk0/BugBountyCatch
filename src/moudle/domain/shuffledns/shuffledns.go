package shuffledns

import (
	"BugBountyCatch/src/config"
	"BugBountyCatch/src/moudle/logger"
	"bufio"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/shuffledns/pkg/runner"
	"os"
	"path/filepath"
)

var domains []string

func ExecuteShuffledns(domain string) []string {
	// Parse the command line flags and read config files
	options := &runner.Options{}
	options.Domain = domain
	outfile := "shufflednsdomain.txt"
	path, _ := os.Getwd()
	tmp := filepath.Join(path, outfile)
	_, err := os.OpenFile(tmp, os.O_APPEND|os.O_CREATE, 0644)
	options.Output = tmp
	options.ResolversFile = config.GetResolversFilePath()
	options.Wordlist = config.GetSubdomainList()
	massdnsRunner, err := runner.New(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}
	massdnsRunner.RunEnumeration()
	massdnsRunner.Close()
	file, err := os.OpenFile(tmp, os.O_RDONLY, 0644)
	if err != nil {
		logger.Logging("shuffledns 获取子域名失败 " + err.Error())
		return domains
	}
	r := bufio.NewScanner(file)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		domains = append(domains, r.Text())
	}
	return domains

}
