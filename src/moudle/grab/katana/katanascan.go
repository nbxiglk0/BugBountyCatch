package katana

import (
	"BugBountyCatch/src/Catchconfig"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/katana/pkg/engine/standard"
	"github.com/projectdiscovery/katana/pkg/output"
	"github.com/projectdiscovery/katana/pkg/types"
)

func Executable(urls []string, outputFile string) {
	extensionFilter := Catchconfig.InitConfig.KatanaConfig.ExtensionFilter
	options := &types.Options{
		OutputFile:        outputFile,
		Headless:          true,
		AutomaticFormFill: true,
		HeadlessNoSandbox: true,
		ExtensionFilter:   extensionFilter,
		Silent:            true,
		MaxDepth:          3,               // Maximum depth to crawl
		FieldScope:        "rdn",           // Crawling Scope Field
		BodyReadSize:      2 * 1024 * 1024, // Maximum response size to read
		RateLimit:         150,             // Maximum requests to send per second
		OnResult: func(result output.Result) { // Callback function to execute for result
			gologger.Info().Msg(result.URL)
		},
	}
	crawlerOptions, err := types.NewCrawlerOptions(options)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	defer crawlerOptions.Close()
	crawler, err := standard.New(crawlerOptions)
	if err != nil {
		gologger.Fatal().Msg(err.Error())
	}
	defer crawler.Close()
	for _, url := range urls {
		err = crawler.Crawl(url)
		if err != nil {
			gologger.Warning().Msgf("Could not crawl %s: %s", url, err.Error())
		}
	}
}
