package gau

import (
	"crypto/tls"
	"github.com/lc/gau/v2/pkg/output"
	"github.com/lc/gau/v2/pkg/providers"
	"github.com/lc/gau/v2/runner"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"io"
	"os"
	"sync"
)

type URLScanConfig struct {
	Host   string `mapstructure:"host"`
	APIKey string `mapstructure:"apikey"`
}

type Config struct {
	Filters           providers.Filters `mapstructure:"filters"`
	Proxy             string            `mapstructure:"proxy"`
	Threads           uint              `mapstructure:"threads"`
	Timeout           uint              `mapstructure:"timeout"`
	Verbose           bool              `mapstructure:"verbose"`
	MaxRetries        uint              `mapstructure:"retries"`
	IncludeSubdomains bool              `mapstructure:"subdomains"`
	RemoveParameters  bool              `mapstructure:"parameters"`
	Providers         []string          `mapstructure:"providers"`
	Blacklist         []string          `mapstructure:"blacklist"`
	JSON              bool              `mapstructure:"json"`
	URLScan           URLScanConfig     `mapstructure:"urlscan"`
	OTX               string            `mapstructure:"otx"`
	Outfile           string            // output file to write to
}

func Executable(scanDomains []string, outputFile string) {
	c := &Config{
		Filters:           providers.Filters{},
		Proxy:             "",
		Timeout:           45,
		Threads:           10,
		Verbose:           false,
		MaxRetries:        5,
		IncludeSubdomains: true,
		RemoveParameters:  false,
		Providers:         []string{"wayback", "commoncrawl", "otx"},
		Blacklist:         []string{"png", "jpg", "gif", "jpeg", "svg", "ico", "css", "vue", "webp"},
		JSON:              false,
		Outfile:           outputFile,
	}
	pMap := make(runner.ProvidersMap)
	for _, provider := range c.Providers {
		pMap[provider] = c.Filters
	}
	var dialer fasthttp.DialFunc
	pc := &providers.Config{
		Threads:           c.Threads,
		Timeout:           c.Timeout,
		Verbose:           c.Verbose,
		MaxRetries:        c.MaxRetries,
		IncludeSubdomains: c.IncludeSubdomains,
		RemoveParameters:  c.RemoveParameters,
		Client: &fasthttp.Client{
			TLSConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			Dial: dialer,
		},
		Providers: c.Providers,
		Output:    c.Outfile,
		JSON:      c.JSON,
		URLScan: providers.URLScan{
			Host:   c.URLScan.Host,
			APIKey: c.URLScan.APIKey,
		},
		OTX: c.OTX,
	}

	gau := &runner.Runner{}
	if err := gau.Init(pc, pMap); err != nil {
		log.Warn(err)
	}
	results := make(chan string)
	var out io.Writer
	// Handle results in background
	if pc.Output == "" {
		out = os.Stdout
	} else {
		ofp, err := os.OpenFile(pc.Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("Could not open output file: %v\n", err)
		}
		defer ofp.Close()
		out = ofp
	}
	writeWg := &sync.WaitGroup{}
	writeWg.Add(1)
	if pc.JSON {
		go func() {
			defer writeWg.Done()
			output.WriteURLsJSON(out, results, pc.Blacklist, pc.RemoveParameters)
		}()
	} else {
		go func() {
			defer writeWg.Done()
			if err := output.WriteURLs(out, results, pc.Blacklist, pc.RemoveParameters); err != nil {
				log.Fatalf("error writing results: %v\n", err)
			}
		}()
	}
	domains := make(chan string)
	gau.Start(domains, results)
	for _, domain := range scanDomains {
		domains <- domain
	}
	close(domains)
	// wait for providers to fetch URLS
	gau.Wait()
	// close results channel
	close(results)
	// wait for writer to finish output
	writeWg.Wait()
}
