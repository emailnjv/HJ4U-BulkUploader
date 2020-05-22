package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/BurntSushi/graphics-go/graphics"
)

type ImageHandler struct{}

// NewImageHandler returns an initialized image handler
func NewImageHandler() ImageHandler {
	var result ImageHandler
	return result
}

func (i *ImageHandler) GetImage(imageURL string) ([]byte, error) {
	var result []byte

	// Fire off the get request of the image
	resp, err := http.Get(imageURL)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	result, err = ioutil.ReadAll(resp.Body)

	return result, err
}

func (i *ImageHandler) GetImageExtension(imgURL string) string {
	urlStruct, err := url.Parse(imgURL)
	if err != nil {
		log.Fatal(err)
	}

	pathSplit := strings.Split(urlStruct.Path, "/")
	imgName := pathSplit[len(pathSplit)-1]
	nameSplit := strings.Split(imgName, ".")
	imgExt := nameSplit[len(nameSplit)-1]

	return imgExt
}

// DownloadImage downloads an image, and loads it into a file
// imagePath is the path including the name desired for the image
// imageURL is the url to send the request for the image to
// It returns the path, and an error
func (i *ImageHandler) DownloadImage(imagePath string, imageURL string) (string, error) {
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

func (i *ImageHandler) DownloadImageFromBytes(imageBytes []byte, imagePath string) error {
	// Create blank file at path
	file, err := os.Create(imagePath)
	if err != nil {
		return err
	}

	reader := bytes.NewReader(imageBytes)

	// Copy the image to the file
	_, err = io.Copy(file, reader)
	if err != nil {
		return err
	}
	defer file.Close()

	return err
}

// CreateThumbnailFromJPG creates a new thumbnail from a JPG byte array
// srcImage is the initial image in bytes
// width is the target width for the thumbnail
// height is the target height for the thumbnail
// It returns the byte array of the thumbnail, and an error
func (i *ImageHandler) CreateThumbnailFromJPG(srcImage []byte, width int, height int) ([]byte, error) {
	// Decode the source image
	imageStruct, _, err := image.Decode(bytes.NewReader(srcImage))
	if err != nil {
		return nil, err
	}

	// Dimension of new thumbnail
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Thumbnail function of Graphics
	graphics.Thumbnail(dstImage, imageStruct)

	byteBuffer := new(bytes.Buffer)

	// Encode the newly created thumbnail to new file
	err = jpeg.Encode(byteBuffer, dstImage, &jpeg.Options{Quality: 100})
	if err != nil {
		return nil, err
	}

	return byteBuffer.Bytes(), err
}

// CreateThumbnailFromPNG creates a new thumbnail from a JPG byte array
// srcImage is the initial image in bytes
// width is the target width for the thumbnail
// height is the target height for the thumbnail
// It returns the byte array of the thumbnail, and an error
func (i *ImageHandler) CreateThumbnailFromPNG(srcImage []byte, width int, height int) ([]byte, error) {
	// Decode the source image
	imageStruct, _, err := image.Decode(bytes.NewReader(srcImage))
	if err != nil {
		return nil, err
	}

	// Dimension of new thumbnail
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Thumbnail function of Graphics
	graphics.Thumbnail(dstImage, imageStruct)

	byteBuffer := new(bytes.Buffer)

	// Encode the newly created thumbnail to new file
	err = png.Encode(byteBuffer, dstImage)
	if err != nil {
		return nil, err
	}

	return byteBuffer.Bytes(), err
}
