package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "http://localhost:8080", "URL to check")
	filePath := flag.String("file", "/file/path", "File path to check")

	flag.Parse()

	if *url == "" && *filePath == "" {
		fmt.Println("Need URL or file path to check.")
		os.Exit(1)
	}

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
