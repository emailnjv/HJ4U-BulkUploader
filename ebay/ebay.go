package ebay

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

type EbayClient struct {
	oAuthKey string
}

// NewEbayClient instantiates, and returns a new Ebay client
func NewEbayClient() (EbayClient, error) {
	var result EbayClient
	err := godotenv.Load("../.env")
	if err != nil {
		return result, err
	}
	result.oAuthKey = os.Getenv("EBAY_OAUTH_TOKEN")
	if result.oAuthKey != "" {
		return result, err
	}

	return result, err
}

func (ec *EbayClient) GetProductInfo(catID int, subCatID int, csvData utils.CSVLine) (db.Product, []string, error) {
	apiResponse, err := ec.getItem(csvData.ItemID)
	if err != nil {
		return db.Product{}, []string{}, err
	}
	return ec.parseItem(catID, subCatID, csvData, apiResponse)
}

