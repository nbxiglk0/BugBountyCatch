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

var path, _ = os.Getwd()

func Executable(domain string, validation bool, subdomains string) []string {
	// Parse the command line flags and read config files
	shufflednsConfig := config.InitConfig.ShufflednsConfig
	var domains []string
	options := &runner.Options{}
	options.ResolversFile = config.InitConfig.ResolversList
	options.Threads = 10000
	options.Retries = 5
	options.StrictWildcard = shufflednsConfig.StrictWildcard
	options.WildcardThreads = shufflednsConfig.WildcardThreads
	options.WildcardOutputFile = ""
	options.Silent = shufflednsConfig.Silent
	options.Json = false
	options.Domain = domain
	options.SubdomainsList = ""
	options.Wordlist = ""
	options.MassdnsRaw = ""
	options.Output = ""
	if validation == false {
		tmp := filepath.Join(path, "shufflednsDomain.txt")
		_, err := os.OpenFile(tmp, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			logger.Logging("创建shuffledns输出文件失败")
		}
		options.Output = tmp
		options.Wordlist = shufflednsConfig.SubdomainWordList
	} else {
		options.SubdomainsList = subdomains
		tmp := filepath.Join(path, "shufflednsValidationDomain.txt")
		_, err := os.OpenFile(tmp, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			logger.Logging("创建shuffledns输出文件失败")
		}
		options.Output = tmp
	}
	massdnsRunner, err := runner.New(options)
	if err != nil {
		gologger.Fatal().Msgf("Could not create runner: %s\n", err)
	}
	massdnsRunner.RunEnumeration()
	massdnsRunner.Close()
	file, err := os.OpenFile(options.Output, os.O_RDONLY, 0644)
	if err != nil {
		logger.Logging("shuffledns 获取子域名失败 " + err.Error())
		return domains
	}
	r := bufio.NewScanner(file)
	r.Split(bufio.ScanLines)
	for r.Scan() {
		t := r.Text()
		domains = append(domains, t)
	}
	_ = os.Remove(options.Output)
	return domains

}
