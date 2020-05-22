package siteClient

import (
	"fmt"
	"strings"
	"time"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/ebay"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

const (
	ImageDirectory          = "/home/igivnqlrr5nm/public_html/admin/uploads/products/"
	ThumbnailImageDirectory = "/home/igivnqlrr5nm/public_html/admin/uploads/products/thumb/"
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

	scpClient, err := utils.NewSCPClient()
	if err != nil {
		return result, err
	}

	result.DBClient = dbClient
	result.EbayClient = ebayClient
	result.ImageHandler = imageHandler
	result.SCPClient = scpClient
	result.HTMLParser = utils.HTMLParser{}

	return result, err
}

func (sc *SiteClient) InsertListing(catID int, subCatID int, csvData utils.CSVLine) error {
	product, urlArr, err := sc.EbayClient.GetProductInfo(catID, subCatID, csvData)

	formattedDescription, err := sc.HTMLParser.ParseHTML(product.Description)
	if err != nil {
		return err
	}

	product.Description = formattedDescription

	productID, err := sc.DBClient.InsertProduct(&product)
	if err != nil {
		return fmt.Errorf("error inserting product %v; err = %v", csvData.ItemID, err)
	}

	err = sc.handleImageURLs(productID, urlArr)
	if err != nil {
		return fmt.Errorf("error inserting product images for product %v; imageURLArr = %v; err = %v", csvData.ItemID, urlArr, err)
	}

	return err
}
func (sc SiteClient) test() (string, error) {
	var result string
	var err error

	return result, err
}

func (sc *SiteClient) CallInsertListing(chann chan error, catID int, subCatID int, csvData utils.CSVLine) {
	err := sc.InsertListing(catID, subCatID, csvData)
	chann <- err
}

func (sc *SiteClient) handleImageURLs(productID int, imageURLs []string) error {

	for index, imageURL := range imageURLs {
		var thumbnailBytes []byte

		imageBytes, err := sc.ImageHandler.GetImage(imageURL)
		if err != nil {
			return err
		}

		extension := strings.ToLower(sc.ImageHandler.GetImageExtension(imageURL))

		if extension == "jpg" {
			thumbnailBytes, err = sc.ImageHandler.CreateThumbnailFromJPG(imageBytes, 200, 200)
			if err != nil {
				return err
			}
		} else if extension == "png" {
			thumbnailBytes, err = sc.ImageHandler.CreateThumbnailFromPNG(imageBytes, 200, 200)
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("missing extension; %v", extension)
		}

		fullSizeImageName := fmt.Sprintf("%v_%v.%v", productID, index, extension)
		fullSizeImageErr := sc.ImageHandler.DownloadImageFromBytes(imageBytes, ImageDirectory+fullSizeImageName)
		if fullSizeImageErr != nil {
			return fullSizeImageErr
		}

		thumbNailName := fmt.Sprintf("%v_%v_thumb.%v", productID, index, extension)
		thumbnailImageErr := sc.ImageHandler.DownloadImageFromBytes(thumbnailBytes, ThumbnailImageDirectory+thumbNailName)
		if thumbnailImageErr != nil {
			return thumbnailImageErr
		}

		mediaEntry := db.Media{
			PageID:        productID,
			PageType:      "products",
			MediaFileName: fullSizeImageName,
			MediaThumb:    "thumb/" + thumbNailName,
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

		_, err = sc.DBClient.InsertMedia(&mediaEntry)
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
