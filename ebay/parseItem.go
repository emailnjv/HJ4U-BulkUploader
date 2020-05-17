package ebay

import (
	"html"
	"strconv"
	"strings"
	"time"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

// parseItem parse the raw product data into ready to use structs to insert
// catID is the main category ID
// subCatID is the sub-category ID
// csvData is the product data row returned from the base CSV
// apiCall is the result from calling the GetItem api call from Ebay
// it returns the filled out product struct, an array of image URLs, and an error
func (e EbayClient) parseItem(catID int, subCatID int, csvData utils.CSVLine, apiCall GetItemResponse) (db.Product, []string, error){
	var product db.Product
	var photoArray []string

	rawPrice := csvData.Price[1:]
	price, err := strconv.ParseFloat(rawPrice, 64)
	if err != nil {
		return product, photoArray, err
	}

	quantity, err := strconv.Atoi(csvData.QuantityAvailable)
	if err != nil {
		return product, photoArray, err
	}

	timeStamp := time.Now()

	p := db.Product{
		Name:           strings.TrimSpace(html.UnescapeString(apiCall.Item.Title)),
		Description:    apiCall.Item.Description,
		Price:          price,
		Featured:       0,
		Date:           &timeStamp,
		Main_cat:       catID,
		Sub_cat:        subCatID,
		Qty:            quantity,
		Sku:            csvData.ItemID,
		Product_type:   "simple",
	}
	if apiCall.Item.ProductListingDetails.UPC != "" {
		p.Upc = apiCall.Item.ProductListingDetails.UPC
	} else {
		p.Upc = csvData.ItemID
	}

	return p, apiCall.Item.PictureDetails.PictureURL, err
}


