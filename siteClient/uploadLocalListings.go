package siteClient

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

type listing struct {
	XML    utils.GetItemResponse
	VarXML utils.GetItemGroupResponse
	Images []string
}

type CategoryStruct struct {
	MainCategory int
	SubCategory  int
}

func (sc *SiteClient) UploadLocalListings(listingDirectory string, imageDirectory string, varianceDirectory string, categoryMapping map[string]CategoryStruct) error {
	listingMap := make(map[string]*listing)
	var missingXMLIDArr []string
	var missingPhotoIDArr []string
	var extraPhotoIDArr []string
	var emptyProductAttributeArr []string

	// map xml responses to images
	for xmlErrStruct := range sc.readDir(listingDirectory) {
		if xmlErrStruct.Error != nil {
			return xmlErrStruct.Error
		} else {
			// Create entries in listing mapping
			listingMap[xmlErrStruct.Response.Item.ItemID] = &listing{
				XML:    xmlErrStruct.Response,
				VarXML: utils.GetItemGroupResponse{},
				Images: []string{},
			}
		}
	}

	// map xml variance responses to listings
	for jsonErrStruct := range sc.ReadVarDir(varianceDirectory) {
		if jsonErrStruct.Error != nil {
			continue
		} else {
			if len(jsonErrStruct.Response.Items) > 0 {
				// Create entries in listing mapping
				listingMap[jsonErrStruct.Response.Items[0].LegacyItemID].VarXML =  jsonErrStruct.Response
			}
		}
	}

	// Read image directory
	imageDirectoryContents, err := ioutil.ReadDir(imageDirectory)
	if err != nil {
		return err
	}

	// Iterate over images and insert name into string array
	for _, item := range imageDirectoryContents {

		// Get Name
		photoName := item.Name()

		// Check for regular file
		if item.Mode().IsRegular() && photoName[0] != '.' {

			// Split ID from name
			itemSplitOne := strings.Split(photoName, "_")
			itemID := itemSplitOne[0]

			// Check to see if image name exists
			listing, ok := listingMap[itemID]
			if ok {
				// Insert Image name
				listing.Images = append(listing.Images, photoName)
			}
		}
	}

	lines, err := ReadCsv("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/listings.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & parse listings into DB
	for index, line := range lines {
		// If first line of CSV or listingMap[line[0]] != nil
		if index != 0 || listingMap[line[0]] != nil {
			// Check to see if in category mapping
			categoryStruct, ok := categoryMapping[line[14]]
			// if found
			if ok {
				csvLine := utils.CSVLine{
					ItemID:             line[0],
					CustomLabel:        line[1],
					ProductIDType:      line[2],
					ProductIDValue:     line[3],
					ProductIDValue2:    line[4],
					QuantityAvailable:  line[5],
					Purchases:          line[6],
					Bids:               line[7],
					Price:              line[8],
					StartDate:          line[9],
					EndDate:            line[10],
					Condition:          line[11],
					Type:               line[12],
					ItemTitle:          line[13],
					CategoryLeafName:   line[14],
					CategoryNumber:     line[15],
					PrivateNotes:       line[16],
					SiteListed:         line[17],
					DownloadDate:       line[18],
					VariationDetails:   line[19],
					ProductReferenceID: line[20],
					ConditionID:        line[21],
					OutOfStockControl:  line[22],
				}

				// check if item id for image is in listing map
				listing, ok := listingMap[csvLine.ItemID]
				if ok {

					// Get the ready to go product struct
					product, err := sc.getProductObjects(categoryStruct, listing.XML, csvLine)
					if err != nil {
						return err
					}

					// Insert product
					id, err := sc.DBClient.InsertProduct(&product)
					if err != nil {
						return err
					}

					if len(listingMap[csvLine.ItemID].VarXML.Items) > 0 {
						productAttributes, err := sc.VarianceParser.HandleVariances(id, listingMap[csvLine.ItemID].VarXML.Items)
						if err != nil {
							panic(fmt.Errorf("Error parsing variances\nItem ID: %v\nErr: %#v", csvLine.ItemID, err))
						}
						if len(productAttributes) > 0 {
							err = sc.DBClient.GroupInsertProductAtt(productAttributes)
						} else {
							emptyProductAttributeArr = append(emptyProductAttributeArr, csvLine.ItemID)
						}
					}

					// If images in array
					if len(listing.Images) == 0 {
						missingPhotoIDArr = append(missingPhotoIDArr, csvLine.ItemID)
					} else {
						err = sc.handleImageURLs(id, listing.Images)
						if err != nil {
							return err
						}
					}
				} else {
					extraPhotoIDArr = append(extraPhotoIDArr, line[0])
				}
			} else {
				missingXMLIDArr = append(missingXMLIDArr, line[0])
			}
		}
	}

	if len(missingPhotoIDArr) > 0 {
		file, _ := json.MarshalIndent(missingPhotoIDArr, "", " ")
		_ = ioutil.WriteFile("missingPhotoIDs.json", file, 0644)
	}

	if len(extraPhotoIDArr) > 0 {
		file, _ := json.MarshalIndent(extraPhotoIDArr, "", " ")
		_ = ioutil.WriteFile("extraPhotoIDs.json", file, 0644)
	}

	if len(emptyProductAttributeArr) > 0 {
		file, _ := json.MarshalIndent(emptyProductAttributeArr, "", " ")
		_ = ioutil.WriteFile("emptyProductAttributeIDs.json", file, 0644)
	}

	if len(missingXMLIDArr) > 0 {
		file2, _ := json.MarshalIndent(missingXMLIDArr, "", " ")
		_ = ioutil.WriteFile("missingXMLIDs.json", file2, 0644)
	}

	return nil
}

func (sc *SiteClient) getProductObjects(categoryStruct CategoryStruct, xmlStruct utils.GetItemResponse, line utils.CSVLine) (db.Products, error) {
	var result db.Products

	// Get the product info
	product, _, err := sc.EbayClient.ParseItem(categoryStruct.MainCategory, categoryStruct.SubCategory, line, xmlStruct)
	if err != nil {
		return result, err
	}

	// Standardize the item description
	formattedDescription, err := sc.HTMLParser.ParseHTML(product.Description)
	if err != nil {
		return result, err
	}

	product.Description = formattedDescription

	return product, err
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
