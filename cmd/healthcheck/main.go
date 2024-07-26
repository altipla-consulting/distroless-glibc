package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"
)

func main() {
	var flagURL, flagFile string
	var flagTimeout time.Duration
	flag.StringVar(&flagURL, "url", "", "URL to check")
	flag.StringVar(&flagFile, "file", "", "File path to check")
	flag.DurationVar(&flagTimeout, "timeout", time.Second*5, "Timeout for the request")
	flag.Parse()

	if flagURL != "" {
		ctx, cancel := context.WithTimeout(context.Background(), flagTimeout)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, flagURL, nil)
		if err != nil {
			os.Exit(1)
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			os.Exit(1)
		}
	}

	if flagFile != "" {
		if _, err := os.Stat(flagFile); err != nil {
			os.Exit(1)
		}
	}
}
