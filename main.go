package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Ncodeteam/direct-go/services/autoload"
)

func main() {
	var url string
	fmt.Print("Direct Link Generator \n")
	fmt.Print("Enter URL: ")
	fmt.Scanln(&url)

	if url == "" {
		fmt.Println("Error: URL is empty")
		os.Exit(1)
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		fmt.Println("Error: Invalid URL format")
		os.Exit(1)
	}

	file_name, dlLink, err := autoload.DirectLink(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("File Name:", file_name)
	fmt.Println("Download Link:", dlLink)
}
