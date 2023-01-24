package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadWebsite(url string) error {
	// Send GET request to website
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	// Create new file with website's domain name
	urlparts := strings.Split(url, ".")
	filename := urlparts[1] + ".html"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	defer file.Close()

	// Write response body to file
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Website downloaded successfully!")
	return nil
}

func main() {
	url := "https://www.youtube.com/"
	err := downloadWebsite(url)
	if err != nil {
		return
	}
}
