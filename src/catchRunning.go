package src

import (
	"BugBountyCatch/src/config"
	"BugBountyCatch/src/moudle/domain/assetfinder"
	"BugBountyCatch/src/moudle/domain/shuffledns"
	"BugBountyCatch/src/moudle/domain/subfinder"
	"BugBountyCatch/src/moudle/logger"
	"strings"
)

func CatchRunning(domain string) {
	subfinderResult := string(subfinder.ExecuteSubfinder(domain))
	domains := strings.Split(subfinderResult, "\r\n")
	logger.Logging("subfinder 找到域名数" + string(rune(len(domains))) + "总域名数: " + string(rune(len(config.Domains))))
	for _, domain := range domains {
		config.Domains = append(config.Domains, domain)
	}
	assetfinderResult := assetfinder.ExecuteAssetfinder(domain)
	copy(config.Domains, assetfinderResult)
	logger.Logging("assetfinder 找到域名数" + string(rune(len(assetfinderResult))) + "总域名数: " + string(rune(len(config.Domains))))
	shuffledns.ExecuteShuffledns(domain)
}
