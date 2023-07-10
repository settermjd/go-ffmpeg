package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
	"time"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func downloadFile(fileURL string) string {
	req, _ := http.NewRequest("GET", fileURL, nil)
	resp, _ := http.DefaultClient.Do(req)
	if resp.StatusCode != 200 {
		log.Fatalf("Error while downloading: %v", resp.StatusCode)
	}
	defer resp.Body.Close()

	fileName := path.Base(fileURL)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d\n", fileName, size)

	return fileName
}

func transcodeAudioFile(inputFile string, outputFile string, wg *sync.WaitGroup) {
	defer wg.Done()

	err := ffmpeg.
		Input(inputFile).
		Output(outputFile).
		OverWriteOutput().
		Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transcoded %s to %s.\n", inputFile, outputFile)
}

func main() {
	start := time.Now()

	var fullURLFile string = "https://demo.twilio.com/docs/classic.mp3"
	fileName := downloadFile(fullURLFile)

	var wg sync.WaitGroup
	outputFilenames := [3]string{"classic.ogg", "classic.wav", "classic.flac"}
	for _, outputFile := range outputFilenames {
		wg.Add(1)
		go transcodeAudioFile(fileName, outputFile, &wg)
	}

	wg.Wait()

	elapsed := time.Since(start)
	log.Printf("Execution took %s", elapsed)
}
