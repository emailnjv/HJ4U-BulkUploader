package ebay

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/valyala/fasthttp"

	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

// getItem returns the data about a specific item when passed the ite ID
// It returns an GetItemResponse object, or an error
func (ec *EbayClient) GetItem(itemID string) (utils.GetItemResponse, error) {
	var result utils.GetItemResponse
	var resp fasthttp.Response

	// Get Route
	route := ec.tradingAPIRouteBuilder()

	// Create request body
	XMLBody := utils.GetItemRequest{
		XMLName:       xml.Name{},
		Text:          "utf-8",
		Xmlns:         "urn:ebay:apis:eBLBaseComponents",
		ErrorLanguage: "en_US",
		WarningLevel:  "High",
		DetailLevel:   "ReturnAll",
		ItemID:        itemID,
	}

	// Marshall body into byte array
	out, _ := xml.Marshal(&XMLBody)

	req := fasthttp.AcquireRequest()

	// Create request
	req.SetBodyString(xml.Header + string(out))
	req.Header.SetRequestURI(route)
	req.Header.SetMethod("POST")

	// Attach Request Headers
	req.Header.Add("X-EBAY-API-SITEID", "0")
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", "967")
	req.Header.Add("X-EBAY-API-CALL-NAME", "GetItem")
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("X-EBAY-API-IAF-TOKEN", ec.oAuthKey)

	// Fire off request
	err := ec.Client.Do(req, &resp)
	if err != nil {
		return result, err
	}

	// Unmarshall body into GetItemResponse
	if err := xml.Unmarshal(resp.Body(), &result); err != nil {
		return result, err
	}

	if result.Ack == "Failure" {
		return result, fmt.Errorf("ebay connection failed, check OAuth Key")
	}

	return result, err
}

type byteErrObj struct {
	Body  []byte
	Error *error
}

// getItemRawResponse returns the data about a specific item when passed the ite ID
// It returns an GetItemResponse object, or an error
func (ec *EbayClient) getItemRawResponse(itemIDs ...string) <-chan byteErrObj {
	out := make(chan byteErrObj)
	var wg sync.WaitGroup

	client := fasthttp.Client{
		MaxConnWaitTimeout: time.Second * 600,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*600)
		},
		MaxConnsPerHost: 50000,
	}

	wg.Add(len(itemIDs))

	// Get Route
	route := ec.tradingAPIRouteBuilder()

	var waitTime string
	if waitTime = os.Getenv("API_CALL_DELAY"); waitTime == "" {
		waitTime = "0"
	}

	waitTimeInt, err := strconv.Atoi(waitTime)
	if err != nil {
		panic(err)
	}

	for _, itemID := range itemIDs {
		time.Sleep(time.Duration(waitTimeInt) * time.Millisecond)
		go func(client *fasthttp.Client, itemID string) {
			defer wg.Done()
			var response fasthttp.Response
			var result byteErrObj

			// Create request body
			XMLBody := utils.GetItemRequest{
				XMLName:       xml.Name{},
				Text:          "utf-8",
				Xmlns:         "urn:ebay:apis:eBLBaseComponents",
				ErrorLanguage: "en_US",
				WarningLevel:  "High",
				DetailLevel:   "ReturnAll",
				ItemID:        itemID,
			}

			// Marshall body into byte array
			reqBody, _ := xml.Marshal(&XMLBody)

			req := fasthttp.AcquireRequest()

			// Create request
			req.SetBodyString(xml.Header + string(reqBody))
			req.Header.SetRequestURI(route)
			req.Header.SetMethod("POST")

			// Attach Request Headers
			req.Header.Add("X-EBAY-API-SITEID", "0")
			req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", "967")
			req.Header.Add("X-EBAY-API-CALL-NAME", "GetItem")
			req.Header.Add("Content-Type", "text/xml")
			req.Header.Add("X-EBAY-API-IAF-TOKEN", ec.oAuthKey)

			fmt.Printf("Sending Request for: %v\n", itemID)

			// Fire off request
			err = client.Do(req, &response)
			if err != nil {
				result.Error = &err
			}
			result.Body = response.Body()

			out <- result
		}(&client, itemID)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (ec *EbayClient) downloadResp(in <-chan byteErrObj, downloadPath string) <-chan *error {
	var wg sync.WaitGroup
	counter := 0
	out := make(chan *error)

	for resp := range in {
		counter++
		wg.Add(1)

		go func(resp byteErrObj, counter int) {
			var result utils.GetItemResponse
			defer wg.Done()

			if resp.Error != nil {
				out <- resp.Error
				return
			}

			err := xml.Unmarshal(resp.Body, &result)
			if err != nil {
				out <- &err
			}

			file, _ := xml.MarshalIndent(result, "", " ")
			if err != nil {
				out <- &err
			}

			err = ioutil.WriteFile(fmt.Sprintf("%v/%v.xml", downloadPath, result.Item.ItemID), file, 0644)
			if err != nil {
				out <- &err
			}
			fmt.Printf("Downloaded XML Response: %v\n", counter)
		}(resp, counter)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
