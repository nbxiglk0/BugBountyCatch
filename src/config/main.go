package config

import (
	"BugBountyCatch/src/moudle/logger"
	"bufio"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var Domains []string
var config = new(Config)
var defaultResolvers []string

type Config struct {
	SubfinderConfigFile string `yaml:"subfinderConfigFile"`
	Thread              int    `yaml:"Thread"`
	ResolversList       string `yaml:"resolversList"`
	SubdomainWordList   string `yaml:"subdomainWordList"`
	MassdnsPath         string `yaml:"massdnsPath"`
}

const (
	configFileName = "catchConfig.yaml"
)

func ParseConfig() {
	path, _ := os.Getwd()
	defaultConfigFile := filepath.Join(path, configFileName)
	_, err := os.Stat(defaultConfigFile)
	if err != nil {
		logger.Logging("找不到配置文件")
		os.Exit(1)
	}
	data, err := os.ReadFile(defaultConfigFile)
	if err != nil {
		logger.Logging("读取配置文件失败")
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		logger.Logging("解析配置文件失败" + err.Error())
	}
}

func GetSubfinder() string {
	return config.SubfinderConfigFile
}
func GetThread() int {
	return config.Thread
}
func GetSubdomainList() string {
	return config.SubdomainWordList
}
func GetResolversFilePath() string {
	return config.ResolversList
}
func GetResolversList() []string {
	f, err := os.Open(config.ResolversList)
	if err != nil {
		logger.Logging("读取ResolversList失败,使用默认Resolvers")
		defaultResolvers = []string{
			"1.1.1.1:53",        // Cloudflare primary
			"1.0.0.1:53",        // Cloudflare secondary
			"8.8.8.8:53",        // Google primary
			"8.8.4.4:53",        // Google secondary
			"9.9.9.9:53",        // Quad9 Primary
			"9.9.9.10:53",       // Quad9 Secondary
			"77.88.8.8:53",      // Yandex Primary
			"77.88.8.1:53",      // Yandex Secondary
			"208.67.222.222:53", // OpenDNS Primary
			"208.67.220.220:53", // OpenDNS Secondary
		}
	} else {
		r := bufio.NewScanner(f)
		r.Split(bufio.ScanLines)
		for r.Scan() {
			resolver := r.Text()
			if resolver == "" {
				continue
			}
			resolver = resolver + ":53"
			defaultResolvers = append(defaultResolvers, resolver)
		}
	}
	return defaultResolvers
}
