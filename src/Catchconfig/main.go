package Catchconfig

import (
	"BugBountyCatch/src/moudle/logger"
	"bufio"
	"github.com/fatih/color"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strings"
)

var Domains []string
var InitConfig = new(Config)
var defaultResolvers []string
var Homedir string
var path, _ = os.Getwd()
var TargetDomain string

type subfinderConfig struct {
	SubfinderConfigFile string `yaml:"subfinderConfigFile"`
	Threads             int    `yaml:"Threads"`
	Silent              bool   `yaml:"Silent"`
	All                 bool   `yaml:"All"`
	RemoveWildcard      bool   `yaml:"RemoveWildcard"`
}
type shufflednsConfig struct {
	SubdomainWordList string `yaml:"subdomainWordList"`
	Threads           int64  `yaml:"Threads"`
	WildcardThreads   int    `yaml:"WildcardThreads"`
	StrictWildcard    bool   `yaml:"StrictWildcard"`
	Silent            bool   `yaml:"Silent"`
}

type naabuConfig struct {
	Threads   int    `yaml:"Threads"`
	TopPorts  string `yaml:"TopPorts"`
	Silent    bool   `yaml:"Silent"`
	OutputCDN bool   `yaml:"OutputCDN"`
}
type httpxConfig struct {
	Threads             int    `yaml:"Threads"`
	Methods             string `yaml:"Methods"`
	Silent              bool   `yaml:"Silent"`
	StatusCode          bool   `yaml:"StatusCode"`
	Location            bool   `yaml:"Location"`
	ExtractTitle        bool   `yaml:"ExtractTitle"`
	TechDetect          bool   `yaml:"TechDetect"`
	OutputIP            bool   `yaml:"OutputIP"`
	OutputCName         bool   `yaml:"OutputCName"`
	FollowHostRedirects bool   `yaml:"FollowHostRedirects"`
}
type nucleiConfig struct {
	Eid       string `yaml:"Eid"`
	ParsedEid []string
	Threads   int  `yaml:"Threads"`
	Debug     bool `yaml:"Debug"`
	Silent    bool `yaml:"Silent"`
}
type katanaConfig struct {
	FilterExt       string `yaml:"FilterExt"`
	ExtensionFilter []string
}
type Config struct {
	SubfinderConfig  subfinderConfig  `yaml:"subfinder"`
	ShufflednsConfig shufflednsConfig `yaml:"shuffledns"`
	NaabuConfig      naabuConfig      `yaml:"naabu"`
	HttpxConfig      httpxConfig      `yaml:"httpx"`
	NucleiConfig     nucleiConfig     `yaml:"nuclei"`
	KatanaConfig     katanaConfig     `yaml:"katana"`
	ResolversList    string           `yaml:"resolversList"`
	Proxy            string           `yaml:"proxy"`
}

const (
	configFileName = "catchConfig.yaml"
)

func CreateHome() {
	err := os.Mkdir(filepath.Join(path, TargetDomain), 0777)
	if err != nil {
		color.Red("创建域名文件夹失败" + err.Error())
		os.Exit(-1)
	}
	Homedir = filepath.Join(path, TargetDomain)
}
func ParseConfig(domain string) {
	TargetDomain = domain
	CreateHome()
	logger.InitLogFile(Homedir)
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
	err = yaml.Unmarshal(data, &InitConfig)
	InitConfig.NucleiConfig.ParsedEid = strings.Split(InitConfig.NucleiConfig.Eid, ",")
	InitConfig.KatanaConfig.ExtensionFilter = strings.Split(InitConfig.KatanaConfig.FilterExt, ",")
	if err != nil {
		logger.Logging("解析配置文件失败" + err.Error())
	}
}
func GetResolversList() []string {
	f, err := os.Open(InitConfig.ResolversList)
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
