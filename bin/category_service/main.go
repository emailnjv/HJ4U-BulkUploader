package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/emailnjv/HJ4U-BulkUploader/siteClient"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

func init() {
}

func main() {
	// err := runGroupedRespDownload("Art Posters", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Art Posters Done")
	// }
	// err = runGroupedRespDownload("Art Prints", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println("Art Prints")
	// }
	categories := map[string]bool{
		"Beads": true,
		"Ceramic, Clay, Porcelain": true,
	}
	// var wg sync.WaitGroup
	// out := make(chan *error)

	// wg.Add(len(categories))

	// for _, category := range categories {
		// go func(category string, wg *sync.WaitGroup) {
			for err := range runGroupedRespDownload(categories, "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses") {
				fmt.Println(err)
			}
				// out <- err
	// 		}
	// 	}(category, &wg)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(out)
	// }()

	// for err := range out {
	// 	fmt.Println(err)
	// }
	fmt.Println("Finished")


	// err := runGroupedRespDownload("Beads", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("Beads Done")
// }
// err = runResponseDownload("Ceramic, Clay, Porcelain", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// } else {
// 	fmt.Println("Ceramic, Clay, Porcelain Done")
// }
// err := runResponseDownload("Collections, Lots", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Bead Caps", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Connectors", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Chains", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Clasps & Hooks", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Chains, Necklaces & Pendants", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Charms & Pendants", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Bracelets", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Cabochons", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Carved Figures", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Denby/Langley/Lovatts", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Earring Findings", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Earrings", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Eggs", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Frames", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Franciscan", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Jewelry Clasps & Hooks", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Jewelry Clasps & Hooks", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Jewelry Making Chains", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Metals", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Other Craft Jewelry Findings", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Other Fine Necklaces, Pendants", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Other Jewelry Design Findings", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Other Loose Gemstones", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Other Sapphires", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Owls", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Rhinestones", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Stone", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
// err = runResponseDownload("Trinket Boxes", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/responses")
// if err != nil {
// 	panic(err)
// }
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

func runResponseDownload(targetCategory string, downloadDirectory string) error {

	var APICallDelayString string
	if APICallDelayString = os.Getenv("API_CALL_DELAY"); APICallDelayString == "" {
		APICallDelayString = "0"
	}
	APICallDelay, err := strconv.Atoi(APICallDelayString)
	if err != nil {
		return err
	}

	lines, err := ReadCsv("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/listings.csv")
	if err != nil {
		panic(err)
	}

	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}

	var csvWG sync.WaitGroup
	itemIDChan := make(chan string)

	foundLineCounter := 0
	// Loop through lines & turn into object
	for index, line := range lines {
		if index != 0 {
			if line[14] == targetCategory {
				fmt.Printf("Found Item ID: %v; Index: %v\n", line[0], foundLineCounter)
				foundLineCounter++
				csvWG.Add(1)
				go func(csvWG *sync.WaitGroup) {
					defer csvWG.Done()
					itemIDChan <- line[0]
					return
				}(&csvWG)
			}
		}
	}

	go func() {
		defer close(itemIDChan)
		csvWG.Wait()
		return
	}()

	var responseWG sync.WaitGroup
	errChan := make(chan error)
	itemIDCounter := 0
	for itemID := range itemIDChan {
		responseWG.Add(1)
		time.Sleep(time.Duration(APICallDelay) * time.Millisecond)
		itemID := itemID
		go func(responseWG *sync.WaitGroup, itemIDCounter int) {
			fmt.Printf("calling item ID: %v; Index: %v\n", itemID, itemIDCounter)
			apiResponse, err := sc.EbayClient.GetItem(itemID)
			if err != nil {
				errChan <- err
				responseWG.Done()
				return
			}
			if apiResponse.Ack == "Failure" {
				errChan <- fmt.Errorf("response not found")
				responseWG.Done()
				return
			}
			fmt.Printf("downloaded file ID: %v; Index: %v\n", itemID, itemIDCounter)
			err = apiResponse.ToFile(responseWG, downloadDirectory)
			errChan <- err
			return
		}(&responseWG, itemIDCounter)
		itemIDCounter++
	}

	go func() {
		defer close(errChan)
		responseWG.Wait()
		return
	}()

	errChanCounter := 0
	for errResult := range errChan {
		fmt.Printf("result returned: %v; counter: %v\n", errResult, errChanCounter)
		errChanCounter++
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
