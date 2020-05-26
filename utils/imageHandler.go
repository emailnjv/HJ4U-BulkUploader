package utils

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"net"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/valyala/fasthttp"
)

type ImageHandler struct {
	Client *fasthttp.Client
}

type ImageDownloadInfo struct {
	Err          error
	PhotoName    string
	ImageURL     string
	ImageBytes   []byte
	DownloadPath string
}

type BytesError struct {
	Bytes []byte
	Err   error
}

// NewImageHandler returns an initialized image handler
func NewImageHandler() ImageHandler {
	var result ImageHandler
	result.Client = &fasthttp.Client{
		MaxConnWaitTimeout: time.Second * 600,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*600)
		},
		MaxConnsPerHost: 50000,
	}
	return result
}

// func (i *ImageHandler) CallDownloadImage(wg *sync.WaitGroup, out chan error, info *ImageDownloadInfo) {
// 	defer wg.Done()
// 	out <- i.downloadImage(*info)
// }

const GatewayTimeoutEx string = "<HTML><HEAD>\n<TITLE>Gateway Timeout - In read </TITLE>\n</HEAD><BODY>\n<H1>Gateway Timeout</H1>\nThe proxy server did not receive a timely response from the upstream server.<P>\nReference&#32;&#35;1&#46;97ba1002&#46;1577198088&#46;5759d5b\n</BODY></HTML>"

func (i *ImageHandler) DownloadImage(wg *sync.WaitGroup, out chan error, imageInfo ImageDownloadInfo) {
	statusCode, imageBytes, err := i.Client.Get(nil, imageInfo.ImageURL)
	if err != nil || statusCode != 200 {
		defer wg.Done()
		out <- fmt.Errorf("status code: %v\nbody: %v\nImageName: %v\nImageURL: %v\nErr: %v\n", statusCode, string(imageBytes), imageInfo.PhotoName, imageInfo.ImageURL, err)
		return
	}

	imageInfo.ImageBytes = imageBytes

	for err := range i.downloadImagesFromBytes(wg, imageInfo) {
		if err != nil {
			out <- err
			return
		}
	}
}

func (i *ImageHandler) getImage(imageURL string) ([]byte, error) {
	var result []byte

	// Fire off the get request of the image
	_, _, err := i.Client.Get(result, imageURL)
	if err != nil {
		return result, err
	}
	return result, err
}

func (i *ImageHandler) downloadImagesFromBytes(wg *sync.WaitGroup, info ImageDownloadInfo) <-chan error {
	var innerWg sync.WaitGroup
	out := make(chan error)

	innerWg.Add(2)

	go i.downloadImageFromBytes(&innerWg, info)

	nameSplit := strings.Split(info.PhotoName, ".")
	imgExt := nameSplit[len(nameSplit)-1]

	if imgExt == "jpg" {
		go i.downloadImageFromBytes(&innerWg, i.createThumbnailFromJPG(info, 200, 200))
	} else if imgExt == "png" {
		info := i.createThumbnailFromPNG(info, 200, 200)
		go i.downloadImageFromBytes(&innerWg, info)
	} else {
		fmt.Printf("image extension not found\n%#v\n%v\n", info, nameSplit[0])
		innerWg.Done()
	}

	go func() {
		innerWg.Wait()
		wg.Done()
		close(out)
	}()

	return out
}

func (i *ImageHandler) downloadImageFromBytes(wg *sync.WaitGroup, imageInfo ImageDownloadInfo) {
	defer wg.Done()
	if imageInfo.Err != nil {
		fmt.Println(imageInfo.Err)
		return
	}
	// Create blank file at path
	file, err := os.Create(imageInfo.DownloadPath + "/" + imageInfo.PhotoName)
	if err != nil {
		imageInfo.Err = err
		fmt.Println(imageInfo)
		return
	}

	reader := bytes.NewReader(imageInfo.ImageBytes)

	// Copy the image to the file
	_, err = io.Copy(file, reader)
	if err != nil {
		imageInfo.Err = err
		fmt.Println(imageInfo)
		return
	}
	defer file.Close()
}

// CreateThumbnailFromJPG creates a new thumbnail from a JPG byte array
// srcImage is the initial image in bytes
// width is the target width for the thumbnail
// height is the target height for the thumbnail
// It returns the byte array of the thumbnail, and an error
func (i *ImageHandler) createThumbnailFromJPG(info ImageDownloadInfo, width int, height int) ImageDownloadInfo {

	// Decode the source image
	imageStruct, _, err := image.Decode(bytes.NewReader(info.ImageBytes))
	if err != nil {
		info.Err = err
		return info
	}

	// Dimension of new thumbnail
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Thumbnail function of Graphics
	err = graphics.Thumbnail(dstImage, imageStruct)
	if err != nil {
		info.Err = err
		return info
	}
	byteBuffer := new(bytes.Buffer)

	// Encode the newly created thumbnail to new file
	err = jpeg.Encode(byteBuffer, dstImage, &jpeg.Options{Quality: 100})
	if err != nil {
		info.Err = err
		return info
	}

	// result.Bytes = byteBuffer.Bytes()
	// out <- &result
	info.ImageBytes = byteBuffer.Bytes()
	info.DownloadPath = info.DownloadPath + "/thumb"
	return info
}

// CreateThumbnailFromPNG creates a new thumbnail from a JPG byte array
// srcImage is the initial image in bytes
// width is the target width for the thumbnail
// height is the target height for the thumbnail
// It returns the byte array of the thumbnail, and an error
func (i *ImageHandler) createThumbnailFromPNG(info ImageDownloadInfo, width int, height int) ImageDownloadInfo {
	// Decode the source image
	imageStruct, _, err := image.Decode(bytes.NewReader(info.ImageBytes))
	if err != nil {
		info.Err = err
		return info
	}

	// Dimension of new thumbnail
	dstImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Thumbnail function of Graphics
	err = graphics.Thumbnail(dstImage, imageStruct)
	if err != nil {
		info.Err = err
		return info
	}

	byteBuffer := new(bytes.Buffer)

	// Encode the newly created thumbnail to new file
	err = png.Encode(byteBuffer, dstImage)
	if err != nil {
		info.Err = err
		return info
	}

	// result.Bytes = byteBuffer.Bytes()
	// out <- &result
	info.ImageBytes = byteBuffer.Bytes()
	info.DownloadPath = info.DownloadPath + "/thumb"
	return info
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
