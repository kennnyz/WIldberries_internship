package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func downloadWebsite(url string) {
	// Send GET request to website
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// Create new file with website's domain name
	urlparts := strings.Split(url, ".")
	filename := urlparts[1] + ".html"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	// Write response body to file
	_, err = file.Write(body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println("Website downloaded successfully!")
}

func main() {
	url := "https://www.youtube.com/"
	downloadWebsite(url)
}
