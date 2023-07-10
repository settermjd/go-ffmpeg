package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(fileURL string, fileName string) {
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(fileURL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

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
	var fullURLFile, fileName string = "https://demo.twilio.com/docs/classic.mp3", "classic.mp3"
	downloadFile(fullURLFile, fileName)
}
