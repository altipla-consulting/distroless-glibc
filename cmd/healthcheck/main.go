package main

import (
	"flag"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "", "URL to check")
	filePath := flag.String("file", "", "File path to check")

	flag.Parse()

	if *url != "" {
		resp, err := http.Get(*url)
		if err != nil {
			os.Exit(1)
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			os.Exit(1)
		}
	}

	if *filePath != "" {
		if _, err := os.Stat(*filePath); os.IsNotExist(err) {
			os.Exit(1)
		}
	}
}
