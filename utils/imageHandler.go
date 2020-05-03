package utils

import (
	"io"
	"net/http"
	"os"
)

type ImageHandler struct{}

// NewImageHandler returns an initialized image handler
func NewImageHandler() ImageHandler {
	var result ImageHandler
	return result
}

// func (i ImageHandler) DownloadImage(path string, imageName string, imageURL string) (string, error) {
func (i ImageHandler) DownloadImage(imagePath string, imageURL string) (string, error) {
	var result string

	// Fire off the get request of the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	// Create blank file at path
	file, err := os.Create(imagePath)
	if err != nil {
		return result, err
	}

	// Copy the image to the file
	size, err := io.Copy(file, resp.Body)
	if err != nil {
		return result, err
	}

	defer file.Close()

	return "Just Downloaded " + imagePath + " with size " + string(size), err
}