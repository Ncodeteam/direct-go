package main

import (
	"fmt"
	"os"
)

func main() {
	var url string
	fmt.Print("Direct Link Generator \n")
	fmt.Print("Enter URL: ")
	fmt.Scanln(&url)

	file_name, dlLink, err := direct_link(url)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println("File Name:", file_name)
	fmt.Println("Download Link:", dlLink)
}
