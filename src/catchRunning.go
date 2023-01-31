package src

import (
	"BugBountyCatch/src/config"
	"BugBountyCatch/src/moudle/domain/assetfinder"
	"BugBountyCatch/src/moudle/domain/shuffledns"
	"BugBountyCatch/src/moudle/domain/subfinder"
	"BugBountyCatch/src/moudle/logger"
	"fmt"
	"os"
	"path/filepath"
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
	shufflednsResult := shuffledns.ExecuteShuffledns(domain)
	copy(config.Domains, shufflednsResult)
	logger.Logging("shuffledns 找到域名数" + string(rune(len(shufflednsResult))) + "总域名数: " + string(rune(len(config.Domains))))
	domainsFile := domain + "_domains.txt"
	path, _ := os.Getwd()
	domainsFile = filepath.Join(path, domainsFile)
	file, err := os.OpenFile(domainsFile, os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logger.Logging("保存域名结果失败 " + err.Error())
	}
	for _, subdomain := range config.Domains {
		_, err := file.WriteString("\r\n" + subdomain)
		if err != nil {
			logger.Logging("写入域名结果失败 " + err.Error())
		}
	}
	fmt.Println("共找到域名: " + string(rune(len(config.Domains))))
}
