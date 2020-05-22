package ebay

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

type EbayClient struct {
	oAuthKey string
	Log      log.Logger
}

// NewEbayClient instantiates, and returns a new Ebay client
func NewEbayClient() (EbayClient, error) {
	err := godotenv.Load("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/.env")
	if err != nil {
		return EbayClient{}, err
	}
	result := EbayClient{
		oAuthKey: os.Getenv("EBAY_OAUTH_TOKEN"),
		Log:      log.Logger{},
	}
	if result.oAuthKey != "" {
		return result, err
	}

	return result, err
}

func (ec *EbayClient) GetProductInfo(catID int, subCatID int, csvData utils.CSVLine) (db.Product, []string, error) {
	apiResponse, err := ec.GetItem(csvData.ItemID)
	if err != nil {
		return db.Product{}, []string{}, err
	}
	return ec.parseItem(catID, subCatID, csvData, apiResponse)
}

func (ec *EbayClient) EbayClientLog(logString string) {
	ec.Log.Printf("Ebay Client Log: %s\n", logString)
}
