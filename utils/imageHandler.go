package utils

import (
	"image"
	"image/jpeg"
	"io"
	"net/http"
	"os"

	"github.com/BurntSushi/graphics-go/graphics"
)

type ImageHandler struct{}

// NewImageHandler returns an initialized image handler
func NewImageHandler() ImageHandler {
	var result ImageHandler
	return result
}

// DownloadImage downloads an image, and loads it into a file
// imagePath is the path including the name desired for the image
// imageURL is the url to send the request for the image to
// It returns the path, and an error
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
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return result, err
	}
	defer file.Close()

	return imagePath, err
}

// CreateThumbnailFromJPG creates a new thumbnail from  a JPG
// srcImagePath is the path including the name desired for the source image
// destImagePath is the path including the name desired for the image thumbnail
// width is the target width for the thumbnail
// height is the target height for the thumbnail
// It returns the path, and an error
func (i ImageHandler) CreateThumbnailFromJPG(srcImagePath string, destImagePath string, width int, height int) (string, error) {
	imagePath, _ := os.Open(srcImagePath)
	defer imagePath.Close()

	// Decode the source image
	srcImage, _, err := image.Decode(imagePath)
	if err != nil {
		return "", err
	}

	// Dimension of new thumbnail
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Thumbnail function of Graphics
	graphics.Thumbnail(dstImage, srcImage)

	// Create new file for thumbnail
	newImage, _ := os.Create(destImagePath)
	defer newImage.Close()

	// Encode the newly created thumbnail to new file
	err = jpeg.Encode(newImage, dstImage, &jpeg.Options{Quality: 100})
	if err != nil {
		return "", err
	}

	return destImagePath, err
}