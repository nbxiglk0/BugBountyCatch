package httpx

import (
	"BugBountyCatch/src/Catchconfig"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"
	"log"
)

func Executable(urlFile string, outPutFile string) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)
	schPorts := []string{"http:80", "https:443"}
	httpxConfig := Catchconfig.InitConfig.HttpxConfig
	options := runner.Options{
		Methods:             httpxConfig.Methods,
		InputFile:           urlFile,
		StatusCode:          httpxConfig.StatusCode,
		Location:            httpxConfig.Location,
		ExtractTitle:        httpxConfig.ExtractTitle,
		TechDetect:          httpxConfig.TechDetect,
		OutputIP:            httpxConfig.OutputIP,
		OutputCName:         httpxConfig.OutputCName,
		Threads:             httpxConfig.Threads,
		FollowHostRedirects: httpxConfig.FollowHostRedirects,
		Silent:              httpxConfig.Silent,
		CustomPorts:         schPorts,
		Resolvers:           Catchconfig.GetResolversList(),
		Output:              outPutFile,
		//InputFile: "./targetDomains.txt", // path to file containing the target domains list
	}

	if err := options.ValidateOptions(); err != nil {
		log.Fatal(err)
	}

	httpxRunner, err := runner.New(&options)
	if err != nil {
		log.Fatal(err)
	}
	defer httpxRunner.Close()
	httpxRunner.RunEnumeration()
}
