package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"time"
)

func main() {
	url := *flag.String("url", "", "URL to check")
	filePath := *flag.String("file", "", "File path to check")
	timeout := *flag.Duration("timeout", time.Second*5, "Timeout for the request")

	flag.Parse()

	if url != "" {
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			os.Exit(1)
		}

		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			os.Exit(1)
		}

		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			os.Exit(1)
		}
	}

	if filePath != "" {
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			os.Exit(1)
		}
	}
}
