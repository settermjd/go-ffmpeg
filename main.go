package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func downloadFile(fileURL string) {
	req, _ := http.NewRequest("GET", fileURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		log.Fatalf("Error while downloading: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	fileName := path.Base(req.URL.Path)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
}

func main() {
	var fullURLFile string = "https://demo.twilio.com/docs/classic.mp3"
	downloadFile(fullURLFile)
}
