package ebay

import (
	"strconv"
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

	timeStamp := time.Now()

	p := db.Product{
		ID:             0,
		Name:           csvData.ItemTitle,
		Description:    apiCall.Item.Description,
		Price:          price,
		Featured:       0,
		Date:           &timeStamp,
		Main_cat:       catID,
		Sub_cat:        subCatID,
		Qty:            int(csvData.QuantityAvailable),
		Sku:            csvData.ItemID,
		Upc:            apiCall.Item.ProductListingDetails.UPC,
		Product_type:   "simple",
	}

	return p, apiCall.Item.PictureDetails.PictureURL, err
}


