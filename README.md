# go-ffmpeg

This is a small project showing how to transcode audio files using Go and FFMpeg.
It's not intended as production-ready code, rather code to show how something works and which forms the basis of an upcoming [Twilio](https://twilio.com/blog/author/msetter) tutorial.

## Application Overview

The application downloads an MP3 file from the Twilio library and then transcodes it to *OGG*, *WAV*, and *FLAC* formats with [FFMpeg](https://ffmpeg.org/).

[Check out this article on transcoding media files](https://opensource.com/article/17/6/ffmpeg-convert-media-file-formats), if you'd like to know more. 

## Prerequisites

- [FFMpeg](https://ffmpeg.org/)
- [Go](https://go.dev/)

## Usage

To use the project, after cloning the repository, run the following command.

```bash
go run main.go
```