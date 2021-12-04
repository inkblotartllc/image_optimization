package main

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/image/bmp"
)

var (
	// Image Path
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
	// All files in the image directory
	allFilesInDirectory []string
)

func main() {
	// Check if the image directory exists
	for _, directory := range imageDirectory {
		if directoryExists(directory) {
			allFilesInDirectory = walkAndAppendPath(directory)
		}
	}
	for _, approvedFileTypes := range imageExtensions {
		for _, file := range allFilesInDirectory {
			if filepath.Ext(file) == approvedFileTypes {
				optimizeImage(file)
			}
		}
	}
}

//
func optimizeImage(path string) {
	switch getFileExtension(path) {
	case ".bmp":
		if getContentType(path) == "image/bmp" {
			removeEXIFDataFromBMPImage(path)
		}
	case ".gif":
		if getContentType(path) == "image/gif" {
			removeEXIFDataFromGIFImage(path)
		}
	case ".jpg", ".jpeg":
		if getContentType(path) == "image/jpeg" {
			removeEXIFDataFromJPEGFile(path)
		}
	case ".png":
		if getContentType(path) == "image/png" {
			removeEXIFDataFromPNGImage(path)
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

// Remove all the exif data from a bmp file.
func removeEXIFDataFromBMPImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	outfile, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = bmp.Encode(outfile, img)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = outfile.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// Remove all the EXIF data from an GIF.
func removeEXIFDataFromGIFImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	outfile, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = gif.Encode(outfile, img, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = outfile.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// Remove EXIF Data from JPEG file.
func removeEXIFDataFromJPEGFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	outfile, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = jpeg.Encode(outfile, img, nil)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = outfile.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

// Remove all the EXIF data from an PNG file
func removeEXIFDataFromPNGImage(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}
	outfile, err := os.Create(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = png.Encode(outfile, img)
	if err != nil {
		log.Fatalln(err)
	}
	err = file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	err = outfile.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
