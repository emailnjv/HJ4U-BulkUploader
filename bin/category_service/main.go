package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/emailnjv/HJ4U-BulkUploader/siteClient"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

func init() {
}

func main() {
	categories := map[string]bool{
		"Beads":                          false,
		"Ceramic, Clay, Porcelain":       false,
		"Art Posters Done":               true,
		"Art Prints":                     true,
		"Collections, Lots":              true,
		"Connectors":                     true,
		"Bead Caps":                      true,
		"Chains":                         true,
		"Clasps & Hooks":                 true,
		"Chains, Necklaces & Pendants":   true,
		"Charms & Pendants":              true,
		"Bracelets":                      true,
		"Cabochons":                      true,
		"Carved Figures":                 true,
		"Denby/Langley/Lovatts":          true,
		"Earring Findings":               true,
		"Earrings":                       true,
		"Frames":                         true,
		"Franciscan":                     true,
		"Jewelry Clasps & Hooks":         true,
		"Jewelry Making Chains":          true,
		"Metals":                         true,
		"Eggs":                           true,
		"Other Craft Jewelry Findings":   true,
		"Other Fine Necklaces, Pendants": true,
		"Other Jewelry Design Findings":  true,
		"Other Loose Gemstones":          true,
		"Other Sapphires":                true,
		"Owls":                           true,
		"Rhinestones":                    true,
		"Stone":                          true,
		"Trinket Boxes":                  true,
	}

	for err := range runGroupedRespDownload(categories, "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses") {
		fmt.Println(err)
	}

	fmt.Println("Finished")
}

func runGroupedRespDownload(categories map[string]bool, downloadDirectory string) <-chan *error {
	var IDArr []string

	lines, err := ReadCsv("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/listings.csv")
	if err != nil {
		panic(err)
	}

	sc, err := siteClient.NewSiteClient()
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for index, line := range lines {
		if index != 0 {
			if categories[line[14]] {
				IDArr = append(IDArr, line[0])
			}
		}
	}

	return sc.EbayClient.DownloadAllResponses(IDArr, downloadDirectory)
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	}
	os.Exit(1)
}

func run(targetCategory string, catID int, subCatID int) error {
	var csvLineArr []utils.CSVLine

	lines, err := ReadCsv("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/listings_mod.csv")
	if err != nil {
		panic(err)
	}

	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}

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

	errChan := make(chan error)
	var csvWg sync.WaitGroup

	csvWg.Add(len(csvLineArr))
	for _, csvLine := range csvLineArr {
		time.Sleep(2 * time.Second)
		go func(catID int, subCatID int, csvLine utils.CSVLine) {

			err := sc.InsertListing(catID, subCatID, csvLine)
			// if err != nil {
			// 	return err
			// }
			defer csvWg.Done()

			errChan <- err
		}(catID, subCatID, csvLine)
	}
	// defer sc.CloseSiteClient()

	go func() {
		csvWg.Wait()
		close(errChan)
		sc.CloseSiteClient()
	}()

	for errResult := range errChan {
		if errResult != nil {
			return errResult
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
