package siteClient

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/valyala/fasthttp"

	"github.com/emailnjv/HJ4U-BulkUploader/ebay"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

func (sc *SiteClient) DownloadAllImages(responseDirectory string, targetDirectory string) {
	var wg sync.WaitGroup
	out := make(chan error)

	var waitTime string
	if waitTime = os.Getenv("API_CALL_DELAY"); waitTime == "" {
		waitTime = "0"
	}

	waitTimeInt, err := strconv.Atoi(waitTime)
	if err != nil {
		panic(err)
	}


	for DownloadInfo := range sc.fetchImageURLS(sc.readDir(responseDirectory)) {
		if DownloadInfo.Err != nil {
			fmt.Println(DownloadInfo.Err)
		}
		DownloadInfo.DownloadPath = targetDirectory
		wg.Add(1)
		time.Sleep(time.Duration(waitTimeInt) * time.Millisecond)
		go sc.ImageHandler.DownloadImage(&wg, out, *DownloadInfo)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for err := range out {
		fmt.Println(err)
	}

	fmt.Println("Finished")
}

type xmlRespObj struct {
	Response ebay.GetItemResponse
	Error    error
}

func (sc *SiteClient) readDir(targetDir string) <-chan *xmlRespObj {
	var wg sync.WaitGroup
	out := make(chan *xmlRespObj)

	xmlFiles, err := ioutil.ReadDir(targetDir)
	if err != nil {
		out <- &xmlRespObj{
			Response: ebay.GetItemResponse{},
			Error:    err,
		}
		return out
	}

	for _, xmlFile := range xmlFiles {
		wg.Add(1)
		if xmlFile.Mode().IsRegular() {
			go func(targetDir string, xmlFile string) {
				var result ebay.GetItemResponse
				defer wg.Done()

				xmlFileByteReader, err := os.Open(fmt.Sprintf("%v/%v", targetDir, xmlFile))
				if err != nil {
					out <- &xmlRespObj{
						Response: result,
						Error:    err,
					}
					return
				}
				defer xmlFileByteReader.Close()

				byteValue, err := ioutil.ReadAll(xmlFileByteReader)
				if err != nil {
					out <- &xmlRespObj{
						Response: result,
						Error:    err,
					}
					return
				}

				err = xml.Unmarshal(byteValue, &result)
				out <- &xmlRespObj{
					Response: result,
					Error:    err,
				}
			}(targetDir, xmlFile.Name())
		}
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (sc *SiteClient) fetchImageURLS(in <-chan *xmlRespObj) <-chan *utils.ImageDownloadInfo {
	var wg sync.WaitGroup
	out := make(chan *utils.ImageDownloadInfo)

	for xmlResp := range in {
		wg.Add(1)
		go func(xmlResp *xmlRespObj) {
			defer wg.Done()
			if xmlResp.Error != nil || xmlResp.Response.Ack == "Failure" {
				out <- &utils.ImageDownloadInfo{
					Err: xmlResp.Error,
				}
				return
			}
			for index, photoURL := range xmlResp.Response.Item.PictureDetails.PictureURL {
				out <- &utils.ImageDownloadInfo{
					Err:       nil,
					PhotoName: fmt.Sprintf("%v_%v.%v", xmlResp.Response.Item.ItemID, index, strings.ToLower(sc.ImageHandler.GetImageExtension(photoURL))),
					ImageURL:  photoURL,
				}
			}
		}(xmlResp)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

type nameNbytes struct {
	PhotoName string
	Bytes     []byte
	Err       error
}

func (sc *SiteClient) fetchImgResponses(in <-chan *utils.ImageDownloadInfo) <-chan *nameNbytes {
	var wg sync.WaitGroup
	out := make(chan *nameNbytes)

	var waitTime string
	if waitTime = os.Getenv("API_CALL_DELAY"); waitTime == "" {
		waitTime = "0"
	}

	waitTimeInt, err := strconv.Atoi(waitTime)
	if err != nil {
		out <- &nameNbytes{
			Err: err,
		}
		return out
	}

	var netClient = &fasthttp.Client{
		MaxConnDuration:               time.Second * 30,
		MaxConnWaitTimeout:            time.Second * 60,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second)
		},
	}

	for urlObj := range in {
		wg.Add(1)
		time.Sleep(time.Duration(waitTimeInt) * time.Millisecond)

		go func(urlObj *utils.ImageDownloadInfo, client *fasthttp.Client) {
			defer wg.Done()
			if urlObj.Err != nil {
				out <- &nameNbytes{
					Err: urlObj.Err,
				}
				return
			}
			fmt.Printf("Requesting URL: %v\n", urlObj.ImageURL)

			var result []byte
			// Fire off the get request of the image
			_, _, err := client.Get(result, urlObj.ImageURL)
			if err != nil {
				out <- &nameNbytes{
					Err: err,
				}
				return
			}

			out <- &nameNbytes{
				PhotoName: strings.ToLower(urlObj.PhotoName),
				Bytes:     result,
				Err:       err,
			}
		}(urlObj, netClient)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (sc *SiteClient) addThumbnail(in <-chan *nameNbytes, downloadDirectory string) <-chan *nameNbytes {
	var wg sync.WaitGroup
	out := make(chan *nameNbytes)

	for byteObj := range in {
		wg.Add(2)

		go func(downloadDirectory string, byteObj *nameNbytes) {
			defer wg.Done()
			var result nameNbytes

			result.PhotoName = fmt.Sprintf("%v/%v", downloadDirectory, byteObj.PhotoName)
			result.Bytes = byteObj.Bytes
			out <- &result
		}(downloadDirectory, byteObj)

		go func(downloadDirectory string, byteObj *nameNbytes) {
			defer wg.Done()

			var result nameNbytes

			if byteObj.Err != nil {
				result.Err = byteObj.Err
				out <- &result
				return
			}

			nameSplit := strings.Split(byteObj.PhotoName, ".")
			imgExt := nameSplit[len(nameSplit)-1]

			if imgExt == "jpg" {

				// Decode the source image
				imageStruct, _, err := image.Decode(bytes.NewReader(byteObj.Bytes))
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				// Dimension of new thumbnail
				dstImage := image.NewRGBA(image.Rect(0, 0, 200, 200))

				// Thumbnail function of Graphics
				graphics.Thumbnail(dstImage, imageStruct)

				byteBuffer := new(bytes.Buffer)

				// Encode the newly created thumbnail to new file
				err = jpeg.Encode(byteBuffer, dstImage, &jpeg.Options{Quality: 100})
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				result.PhotoName = fmt.Sprintf("%v/thumb/%v", downloadDirectory, byteObj.PhotoName)
				result.Bytes = byteBuffer.Bytes()

				out <- &result
			} else if imgExt == "png" {

				// Decode the source image
				imageStruct, _, err := image.Decode(bytes.NewReader(byteObj.Bytes))
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				// Dimension of new thumbnail
				dstImage := image.NewRGBA(image.Rect(0, 0, 200, 200))

				// Thumbnail function of Graphics
				graphics.Thumbnail(dstImage, imageStruct)

				byteBuffer := new(bytes.Buffer)

				// Encode the newly created thumbnail to new file
				err = png.Encode(byteBuffer, dstImage)
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				result.PhotoName = fmt.Sprintf("%v/thumb/%v", downloadDirectory, byteObj.PhotoName)
				result.Bytes = byteBuffer.Bytes()

				out <- &result
			}

		}(downloadDirectory, byteObj)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (sc *SiteClient) addThumbnails(in <-chan *nameNbytes, downloadDirectory string) <-chan *nameNbytes {
	var wg sync.WaitGroup
	out := make(chan *nameNbytes)

	for byteObj := range in {
		wg.Add(2)

		go func(downloadDirectory string, byteObj *nameNbytes) {
			defer wg.Done()
			var result nameNbytes

			result.PhotoName = fmt.Sprintf("%v/%v", downloadDirectory, byteObj.PhotoName)
			result.Bytes = byteObj.Bytes
			out <- &result
		}(downloadDirectory, byteObj)

		go func(downloadDirectory string, byteObj *nameNbytes) {
			defer wg.Done()

			var result nameNbytes

			if byteObj.Err != nil {
				result.Err = byteObj.Err
				out <- &result
				return
			}

			nameSplit := strings.Split(byteObj.PhotoName, ".")
			imgExt := nameSplit[len(nameSplit)-1]

			if imgExt == "jpg" {

				// Decode the source image
				imageStruct, _, err := image.Decode(bytes.NewReader(byteObj.Bytes))
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				// Dimension of new thumbnail
				dstImage := image.NewRGBA(image.Rect(0, 0, 200, 200))

				// Thumbnail function of Graphics
				graphics.Thumbnail(dstImage, imageStruct)

				byteBuffer := new(bytes.Buffer)

				// Encode the newly created thumbnail to new file
				err = jpeg.Encode(byteBuffer, dstImage, &jpeg.Options{Quality: 100})
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				result.PhotoName = fmt.Sprintf("%v/thumb/%v", downloadDirectory, byteObj.PhotoName)
				result.Bytes = byteBuffer.Bytes()

				out <- &result
			} else if imgExt == "png" {

				// Decode the source image
				imageStruct, _, err := image.Decode(bytes.NewReader(byteObj.Bytes))
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				// Dimension of new thumbnail
				dstImage := image.NewRGBA(image.Rect(0, 0, 200, 200))

				// Thumbnail function of Graphics
				graphics.Thumbnail(dstImage, imageStruct)

				byteBuffer := new(bytes.Buffer)

				// Encode the newly created thumbnail to new file
				err = png.Encode(byteBuffer, dstImage)
				if err != nil {
					result.Err = err
					out <- &result
					return
				}

				result.PhotoName = fmt.Sprintf("%v/thumb/%v", downloadDirectory, byteObj.PhotoName)
				result.Bytes = byteBuffer.Bytes()

				out <- &result
			}

		}(downloadDirectory, byteObj)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (sc *SiteClient) createImage(in <-chan *nameNbytes) <-chan error {
	var wg sync.WaitGroup
	out := make(chan error)

	for byteObj := range in {
		wg.Add(1)
		go func(byteObj *nameNbytes) {
			if byteObj.Err != nil {
				out <- byteObj.Err
				return
			}

			defer wg.Done()

			file, err := os.Create(byteObj.PhotoName)
			if err != nil {
				out <- err
				return
			}

			reader := bytes.NewReader(byteObj.Bytes)

			// Copy the image to the file
			byteSize, err := io.Copy(file, reader)
			if err != nil {
				out <- err
				return
			}

			fmt.Printf("Downloaded Image: %v; Size: %v\n", byteObj.PhotoName, byteSize)

			defer file.Close()

			out <- err
		}(byteObj)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
