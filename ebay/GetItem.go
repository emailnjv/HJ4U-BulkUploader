package ebay

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetItem returns the data about a specific item when passed the ite ID
// It returns an GetItemResponse object, or an error
func (e EbayClient) GetItem(itemID string) (GetItemResponse, error) {
	var result GetItemResponse
	client := &http.Client{}

	// Get Route
	route := e.tradingAPIRouteBuilder()

	// Create request body
	XMLBody := GetItemRequest{
		XMLName:       xml.Name{},
		Text:          "utf-8",
		Xmlns:         "urn:ebay:apis:eBLBaseComponents",
		ErrorLanguage: "en_US",
		WarningLevel:  "High",
		ItemID:        itemID,
	}

	// Marshall body into byte array
	out, _ := xml.Marshal(&XMLBody)

	// Create request
	req, err := http.NewRequest("POST", route, strings.NewReader(xml.Header + string(out)))
	if err != nil {
		return result, err
	}

	// Attach Request Headers
	req.Header.Add("X-EBAY-API-SITEID", "0")
	req.Header.Add("X-EBAY-API-COMPATIBILITY-LEVEL", "967")
	req.Header.Add("X-EBAY-API-CALL-NAME", "GetItem")
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("X-EBAY-API-IAF-TOKEN", e.oAuthKey)

	// Fire off request
	resp, err := client.Do(req)
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

	return result, err
}
