package src

import (
	"BugBountyCatch/src/config"
	"BugBountyCatch/src/moudle/domain/assetfinder"
	"BugBountyCatch/src/moudle/domain/shuffledns"
	"BugBountyCatch/src/moudle/domain/subfinder"
	"BugBountyCatch/src/moudle/logger"
	"BugBountyCatch/src/moudle/scan/httpx"
	"BugBountyCatch/src/moudle/scan/naabu"
	"fmt"
	mapSet "github.com/deckarep/golang-set"
	"github.com/fatih/color"
	"github.com/projectdiscovery/fileutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var mes string

func CatchRunning(domain string) {
	var nabbuOutPut = filepath.Join(config.Homedir, domain+"_NabbuScan.txt")
	var httpxOutPut = filepath.Join(config.Homedir, domain+"_HttpxScan.txt")
	var domainsFile = filepath.Join(config.Homedir, domain+"_domains.txt")
	mes = fmt.Sprintf("%s", time.Now().Format("2006-01-02 15:04:05")+": 开始收集子域名: "+domain)
	color.Green(mes)
	logger.Logging(mes)
	//subfinder
	//
	mes = time.Now().Format("2006-01-02 15:04:05") + ":	开始运行 subfinder"
	color.White(mes)
	subfinderResult := string(subfinder.ExecuteSubfinder(domain))
	domains := strings.Split(subfinderResult, "\n")
	for _, domain := range domains {
		config.Domains = append(config.Domains, domain)
	}
	mes = fmt.Sprintf("%s%d%s%d", "subfinder 找到域名数", len(domains), "    总域名数: ", len(config.Domains))
	color.Green(mes)
	logger.Logging(mes)
	//assetfinder
	//
	//
	mes = time.Now().Format("2006-01-02 15:04:05") + ":	开始运行 assetfinder"
	color.White(mes)
	assetfinderResult := assetfinder.ExecuteAssetfinder(domain)
	config.Domains = append(config.Domains, assetfinderResult...)
	mes = fmt.Sprintf("%s%d%s%d", "assetfinder 找到域名数", len(assetfinderResult), "    总域名数: ", len(config.Domains))
	color.Green(mes)
	logger.Logging(mes)
	// shuffledns
	// DNS子域名爆破
	//
	mes = time.Now().Format("2006-01-02 15:04:05") + ":	开始运行 shuffledns"
	color.White(mes)
	shufflednsResult := shuffledns.ExecuteShuffledns(domain, false, "")
	config.Domains = append(config.Domains, shufflednsResult...)
	mes = fmt.Sprintf("%s%d%s%d", "shuffledns 找到域名数", len(shufflednsResult), "    总域名数: ", len(config.Domains))
	color.GreenString(mes)
	logger.Logging(mes)
	//去重
	//
	//
	mes = time.Now().Format("2006-01-02 15:04:05") + ":	开始去重"
	color.White(mes)
	var a []interface{}
	for _, t := range config.Domains {
		a = append(a, t)
	}
	s := mapSet.NewSetFromSlice(a)
	domainsSlict := s.ToSlice()
	mes = fmt.Sprintf("%s%d", "去重后总域名数:  ", len(domainsSlict))
	color.Green(mes)
	logger.Logging(mes)
	//清洗无效域名
	//使用shuffledns验证域名
	//
	mes = time.Now().Format("2006-01-02 15:04:05") + ":	开始清除无效域名"
	color.White(mes)
	logger.Logging(mes)
	tmpDomainsFile := filepath.Join(config.Homedir, "tmpDomains.txt")
	file, err := os.OpenFile(tmpDomainsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	for index, subdomain := range domainsSlict {
		t := fmt.Sprintf("%s", subdomain)
		if t != "" {
			if index != len(domainsSlict)-1 {
				t = t + "\n"
			}
			_, err := file.WriteString(t)
			if err != nil {
				logger.Logging("清洗无效域名结果失败 " + err.Error())
			}
		}
	}
	shufflednsValidationResult := shuffledns.ExecuteShuffledns(domain, true, tmpDomainsFile)
	_ = os.Remove(tmpDomainsFile)
	//
	//
	//写入保存
	if fileutil.FileExists(domainsFile) {
		_ = os.Remove(domainsFile)
	}
	outputFile, err := os.OpenFile(domainsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		logger.Logging("保存域名结果失败 " + err.Error())
	}
	for _, subdomain := range shufflednsValidationResult {
		_, err = outputFile.WriteString(subdomain + "\n")
		if err != nil {
			logger.Logging("写入域名结果失败 " + err.Error())
		}
	}
	err = outputFile.Close()
	if err != nil {
		logger.Logging("写入域名结果失败 " + err.Error())
	}
	mes = fmt.Sprintf("%s%d", "共找到有效存活域名: ", len(shufflednsValidationResult))
	color.Green(mes)
	logger.Logging(mes)
	//扫描端口
	//
	//
	mes = time.Now().Format("2006-01-02 15:04:05") + ": 开始运行Naabu扫描端口,输出到 " + nabbuOutPut
	color.Green(mes)
	logger.Logging(mes)
	naabu.ExecuteNaabu(domainsFile, nabbuOutPut)
	//Httpx扫描
	mes = time.Now().Format("2006-01-02 15:04:05") + ": 开始运行httpx收集web信息,输出到 " + httpxOutPut
	color.Green(mes)
	logger.Logging(mes)
	httpx.Executehttpx(filepath.Join(config.Homedir, domain+"_NabbuScan.txt"), httpxOutPut)
}
