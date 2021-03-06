package main

import (
	"image"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

var (
	// Image
	imageDirectory = []string{
		"/",
	}
	// Image Extensions
	imageExtensions = []string{
		".jpg",
		".jpeg",
		".png",
		".gif",
		".bmp",
	}
)

func init() {
	// Application check
	commandExists("cwebp")
}

func main() {
	// All files in the image directory
	var allFilesInDirectory []string
	// Check if the image directory exists
	for _, directory := range imageDirectory {
		if directoryExists(directory) {
			allFilesInDirectory = walkAndAppendPath(directory)
		}
	}
	for _, file := range allFilesInDirectory {
		for _, approvedFileTypes := range imageExtensions {
			if filepath.Ext(file) == approvedFileTypes {
				if strings.Contains(getContentType(file), "image") {
					getAllDataFromImage(file)
				}
			}
		}
	}
}

// Find all the files in a given directory and append that to a path.
func walkAndAppendPath(walkPath string) []string {
	var filePath []string
	err := filepath.Walk(walkPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fileExists(path) {
			filePath = append(filePath, path)
		}
		return err
	})
	if err != nil {
		log.Fatalln(err)
	}
	return filePath
}

// Check if the file exists and return a bool.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

// Check if a directory exists
func directoryExists(path string) bool {
	directory, err := os.Stat(path)
	if err != nil {
		return false
	}
	return directory.IsDir()
}

// Get the file extension of a file
func getFileExtension(path string) string {
	return filepath.Ext(path)
}

// Get the content type of a file and return it as a string
func getContentType(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	return http.DetectContentType(buffer)
}

// Get all the data from an given image.
func getAllDataFromImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	imageData, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	// Encode the image to webp
	encodeImageToWEBP(path, imageData, 100)
}

// Encode data to webp format.
func encodeImageToWEBP(path string, content image.Image, quality float32) {
	file, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, quality)
	if err != nil {
		log.Fatalln(err)
	}
	err = webp.Encode(file, content, options)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// Check if the application is installed and in path
func commandExists(application string) {
	_, err := exec.LookPath(application)
	if err != nil {
		log.Fatalln(err)
	}
}

/*
Make sure webp is installed in the system.
https://developers.google.com/speed/webp
*/
