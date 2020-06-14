package ebay

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

type EbayClient struct {
	Client           *fasthttp.Client
	oAuthKey         string
	StoreCategoryMap map[int]string
}

// NewEbayClient instantiates, and returns a new Ebay client
func NewEbayClient() (EbayClient, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		err = godotenv.Load(".env")
		if err != nil {
			return EbayClient{}, err
		}
	}

	oAuthKey := os.Getenv("EBAY_OAUTH_TOKEN")
	if oAuthKey == "" {
		return EbayClient{}, fmt.Errorf("no OAuth key found")
	}

	client := fasthttp.Client{
		MaxConnWaitTimeout: time.Second * 600,
		Dial: func(addr string) (net.Conn, error) {
			return fasthttp.DialTimeout(addr, time.Second*600)
		},
		MaxConnsPerHost: 50000,
	}

	result := EbayClient{
		Client:           &client,
		oAuthKey:         oAuthKey,
		StoreCategoryMap: make(map[int]string),
	}

	err = result.loadStoreCategories()
	if err != nil {
		return EbayClient{}, err
	}

	return result, err
}

func (ec *EbayClient) GetProductInfo(catID int, subCatID int, csvData utils.CSVLine) (db.Products, []string, error) {
	apiResponse, err := ec.GetItem(csvData.ItemID)
	if err != nil {
		return db.Products{}, []string{}, err
	}
	return ec.ParseItem(catID, subCatID, csvData, apiResponse)
}
