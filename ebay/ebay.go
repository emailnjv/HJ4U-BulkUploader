package ebay

import (
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

type EbayClient struct {
	Client   *fasthttp.Client
	oAuthKey string
}

// NewEbayClient instantiates, and returns a new Ebay client
func NewEbayClient() (EbayClient, error) {
	var result EbayClient
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			return EbayClient{}, err
		}
	}

	result.oAuthKey = os.Getenv("EBAY_OAUTH_TOKEN")
	if result.oAuthKey != "" {
		return result, err
	}

	client := fasthttp.Client{
		MaxConnWaitTimeout: time.Second * 600,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*600)
		},
		MaxConnsPerHost: 50000,
	}

	result.Client = &client

	return result, err
}

func (ec *EbayClient) GetProductInfo(catID int, subCatID int, csvData utils.CSVLine) (db.Products, []string, error) {
	apiResponse, err := ec.GetItem(csvData.ItemID)
	if err != nil {
		return db.Products{}, []string{}, err
	}
	return ec.ParseItem(catID, subCatID, csvData, apiResponse)
}
