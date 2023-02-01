package httpx

import (
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/httpx/runner"
	"log"
)

func Executehttpx(urlFile string, outPutFile string) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelVerbose) // increase the verbosity (optional)
	options := runner.Options{
		Methods:             "GET",
		InputFile:           urlFile,
		StatusCode:          true,
		Location:            true,
		ExtractTitle:        true,
		TechDetect:          true,
		OutputIP:            true,
		OutputCName:         true,
		Threads:             100,
		FollowHostRedirects: true,
		Silent:              true,
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
