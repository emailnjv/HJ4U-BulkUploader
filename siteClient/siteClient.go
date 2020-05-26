package siteClient

import (
	"sort"
	"time"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/ebay"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

const (
	ImageDirectory          = "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/products/"
	ThumbnailImageDirectory = "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/products/thumb/"
)

type SiteClient struct {
	DBClient     db.TargetDBClient
	EbayClient   ebay.EbayClient
	ImageHandler utils.ImageHandler
	SCPClient    utils.SCPClient
	HTMLParser   utils.HTMLParser
}

func NewSiteClient() (SiteClient, error) {
	var result SiteClient

	dbClient, err := db.NewTargetDBClient()
	if err != nil {
		return result, err
	}

	ebayClient, err := ebay.NewEbayClient()
	if err != nil {
		return result, err
	}

	imageHandler := utils.NewImageHandler()
	//
	// scpClient, err := utils.NewSCPClient()
	// if err != nil {
	// 	return result, err
	// }

	result.DBClient = dbClient
	result.EbayClient = ebayClient
	result.ImageHandler = imageHandler
	// result.SCPClient = scpClient
	result.HTMLParser = utils.HTMLParser{}

	return result, err
}

// func (sc *SiteClient) InsertListing(catID int, subCatID int, csvData utils.CSVLine) error {
// 	product, urlArr, err := sc.EbayClient.GetProductInfo(catID, subCatID, csvData)
//
// 	formattedDescription, err := sc.HTMLParser.ParseHTML(product.Description)
// 	if err != nil {
// 		return err
// 	}
//
// 	product.Description = formattedDescription
//
// 	productID, err := sc.DBClient.InsertProduct(&product)
// 	if err != nil {
// 		return fmt.Errorf("error inserting product %v; err = %v", csvData.ItemID, err)
// 	}
//
// 	err = sc.handleImageURLs(productID, urlArr)
// 	if err != nil {
// 		return fmt.Errorf("error inserting product images for product %v; imageURLArr = %v; err = %v", csvData.ItemID, urlArr, err)
// 	}
//
// 	return err
// }

// func (sc *SiteClient) CallInsertListing(chann chan error, catID int, subCatID int, csvData utils.CSVLine) {
// 	err := sc.InsertListing(catID, subCatID, csvData)
// 	chann <- err
// }

func (sc *SiteClient) handleImageURLs(productID int, imageURLs []string) error {

	sort.Strings(imageURLs)

	for index, imageName := range imageURLs {

		mediaEntry := db.Media{
			PageID:        productID,
			PageType:      "products",
			MediaFileName: imageName,
			MediaThumb:    "thumb/" + imageName,
			Author:        1,
		}

		now := time.Now()
		mediaEntry.Date = &now
		if index == 0 {
			mediaEntry.MediaType = "f_img"
			mediaEntry.OrderID = index + 1
		} else {
			mediaEntry.MediaType = "img"
			mediaEntry.OrderID = index - 1
		}

		_, err := sc.DBClient.InsertMedia(&mediaEntry)
		if err != nil {
			return err
		}
	}
	return nil
}

func (sc *SiteClient) CloseSiteClient() error {
	err := sc.SCPClient.CloseClients()
	if err != nil {
		return err
	}
	err = sc.DBClient.CloseConnection()
	return err
}
