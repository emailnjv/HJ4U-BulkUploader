package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/emailnjv/HJ4U-BulkUploader/siteClient"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

func init() {
}

func main() {
	err := run("Art Posters", 11, 26)
	if err != nil {
		panic(err)
	}
	err = run("Art Prints", 11, 25)
	if err != nil {
		panic(err)
	}
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	}
	os.Exit(1)
}

func run(targetCategory string, catID int, subCatID int) error {
	var csvLineArr []utils.CSVLine

	lines, err := ReadCsv("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/listings.csv")
	if err != nil {
		panic(err)
	}

	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}
	defer sc.CloseSiteClient()

	// Loop through lines & turn into object
	for index, line := range lines {
		if index != 0 {
			data := utils.CSVLine{
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
			if data.CategoryLeafName == targetCategory {
				csvLineArr = append(csvLineArr, data)
			}

		}
	}

	for _, csvLine := range csvLineArr {
		printUsage(csvLine.ItemID)
		err = sc.InsertListing(catID, subCatID, csvLine)
		if err != nil {
			return fmt.Errorf("error inserting item %v; err = %v", csvLine.ItemID, err)
		}
	}

	return nil
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

func printUsage(itemID string) {
	fmt.Println("Inserting Product ID %v", itemID)
}
