package ebay

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

// getItem returns the data about a specific item when passed the ite ID
// It returns an GetItemResponse object, or an error
func (ec *EbayClient) GetItem(itemID string) (GetItemResponse, error) {
	var result GetItemResponse

	// Get Route
	route := ec.tradingAPIRouteBuilder()

	// Create request body
	XMLBody := GetItemRequest{
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

	// Create request
	req, err := http.NewRequest("POST", route, strings.NewReader(xml.Header+string(out)))
	if err != nil {
		return result, err
	}

	// Attach Request Headers
	req.Header.Add("X-EBAY-API-SITEID", "0")
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", "967")
	req.Header.Add("X-EBAY-API-CALL-NAME", "GetItem")
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("X-EBAY-API-IAF-TOKEN", ec.oAuthKey)

	// Fire off request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	// Read the body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Unmarshall body into GetItemResponse
	if err := xml.Unmarshal(body, &result); err != nil {
		return result, err
	}

	if result.Ack == "Failure" {
		return result, fmt.Errorf("ebay connection failed, check OAuth Key")
	}

	return result, err
}

type httpRespObj struct {
	Response *http.Response
	Error    *error
}

// getItemRawResponse returns the data about a specific item when passed the ite ID
// It returns an GetItemResponse object, or an error
func (ec *EbayClient) getItemRawResponse(itemIDs ...string) <-chan *httpRespObj {
	out := make(chan *httpRespObj)
	var wg sync.WaitGroup

	wg.Add(len(itemIDs))

	// Get Route
	route := ec.tradingAPIRouteBuilder()

	for _, itemID := range itemIDs {
		go func(itemID string) {
			defer wg.Done()
			var result httpRespObj

			// Create request body
			XMLBody := GetItemRequest{
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

			// Create request
			req, err := http.NewRequest("POST", route, strings.NewReader(xml.Header+string(reqBody)))
			if err != nil {
				result.Error = &err
				out <- &result
			}

			// Attach Request Headers
			req.Header.Add("X-EBAY-API-SITEID", "0")
			req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", "967")
			req.Header.Add("X-EBAY-API-CALL-NAME", "GetItem")
			req.Header.Add("Content-Type", "text/xml")
			req.Header.Add("X-EBAY-API-IAF-TOKEN", ec.oAuthKey)

			// Fire off request
			result.Response, err = http.DefaultClient.Do(req)
			if err != nil {
				result.Error = &err
			}

			out <- &result
		}(itemID)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func (ec *EbayClient) downloadResp(in <-chan *httpRespObj, downloadPath string) <-chan *error {
	var wg sync.WaitGroup
	counter := 0
	out := make(chan *error)

	for resp := range in {
		wg.Add(1)

		go func(resp *httpRespObj, counter int) {
			defer wg.Done()

			if resp.Error != nil {
				out <- resp.Error
			}
			defer resp.Response.Body.Close()

			file, err := os.Create(fmt.Sprintf("%v/%v.xml", downloadPath, counter))
			if err != nil {
				out <- &err
			}
			defer file.Close()

			_, err = io.Copy(file, resp.Response.Body)
			if err != nil {
				out <- &err
			}

		}(resp, counter)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}