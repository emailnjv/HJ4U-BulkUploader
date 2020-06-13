package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/emailnjv/HJ4U-BulkUploader/siteClient"
)

var categories = map[string]bool{
	"Beads":                          true,
	"Ceramic, Clay, Porcelain":       true,
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
	"Single Flatware Pieces":         true,
	"Buttons":                        true,
	"Other China & Dinnerware":       true,
	"Other French Art Glass":         true,
	"Jewelry Boxes":                  true,
	"Movements":                      true,
	"Wristwatch Bands":               true,
	"Salt & Pepper Shakers":          true,
	"Jewelry Sets":                   true,
	"Brooches, Pins":                 true,
	"Pins, Brooches":                 true,
	"Pin Backs & Brooch Components":  true,
	"Spacer Beads & Stoppers":        true,
}

func init() {
}

func main() {

	// for err := range runGroupedRespDownload(categories, "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/resp") {
	// 	fmt.Printf("Error returned from runGroupedRespDownload: %#v\n", err)
	// }
		fmt.Println(uploadLocalListings())

	fmt.Println("Finished")
}

func uploadLocalListings() error {
	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}
	categoryMapping := sampleMapping()

	return sc.UploadLocalListings("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/responses2","/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/pictures", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/varianceResponses2", categoryMapping)
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

	// jsonFile, err := os.Open("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/siteClient/missingIDs.json")
	// if we os.Open returns an error then handle it
	// if err != nil {
	// 	fmt.Printf("Error opening json file: %#v\n", err)
	// }
	// defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	// byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	// var IDs []string

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	// err = json.Unmarshal(byteValue, &IDs)
	// if err != nil {
	// 	panic(err)
	// }

	// Loop through lines & turn into object
	for index, line := range lines {
		if index != 0 {
			if _, ok := categories[line[14]]; ok {
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

// func run(targetCategory string, catID int, subCatID int) error {
// 	var csvLineArr []utils.CSVLine
//
// 	lines, err := ReadCsv("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/listings_mod.csv")
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	sc, err := siteClient.NewSiteClient()
// 	if err != nil {
// 		return err
// 	}
//
// 	// Loop through lines & turn into object
// 	for index, line := range lines {
// 		if index != 0 {
// 			data := utils.CSVLine{
// 				ItemID:             line[0],
// 				CustomLabel:        line[1],
// 				ProductIDType:      line[2],
// 				ProductIDValue:     line[3],
// 				ProductIDValue2:    line[4],
// 				QuantityAvailable:  line[5],
// 				Purchases:          line[6],
// 				Bids:               line[7],
// 				Price:              line[8],
// 				StartDate:          line[9],
// 				EndDate:            line[10],
// 				Condition:          line[11],
// 				Type:               line[12],
// 				ItemTitle:          line[13],
// 				CategoryLeafName:   line[14],
// 				CategoryNumber:     line[15],
// 				PrivateNotes:       line[16],
// 				SiteListed:         line[17],
// 				DownloadDate:       line[18],
// 				VariationDetails:   line[19],
// 				ProductReferenceID: line[20],
// 				ConditionID:        line[21],
// 				OutOfStockControl:  line[22],
// 			}
// 			if data.CategoryLeafName == targetCategory {
// 				csvLineArr = append(csvLineArr, data)
// 			}
//
// 		}
// 	}
//
// 	errChan := make(chan error)
// 	var csvWg sync.WaitGroup
//
// 	csvWg.Add(len(csvLineArr))
// 	for _, csvLine := range csvLineArr {
// 		time.Sleep(2 * time.Second)
// 		go func(catID int, subCatID int, csvLine utils.CSVLine) {
//
// 			err := sc.InsertListing(catID, subCatID, csvLine)
// 			// if err != nil {
// 			// 	return err
// 			// }
// 			defer csvWg.Done()
//
// 			errChan <- err
// 		}(catID, subCatID, csvLine)
// 	}
// 	// defer sc.CloseSiteClient()
//
// 	go func() {
// 		csvWg.Wait()
// 		close(errChan)
// 		sc.CloseSiteClient()
// 	}()
//
// 	for errResult := range errChan {
// 		if errResult != nil {
// 			return errResult
// 		}
// 	}
//
// 	return nil
// }

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

func sampleMapping() map[string]siteClient.CategoryStruct {
	result := make(map[string]siteClient.CategoryStruct)
	result["Beads"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  16,
	}
	result["Ceramic, Clay, Porcelain"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  16,
	}
	result["Art Posters Done"] = siteClient.CategoryStruct{
		MainCategory: 11,
		SubCategory:  26,
	}
	result["Art Prints"] = siteClient.CategoryStruct{
		MainCategory: 11,
		SubCategory:  26,
	}
	result["Collections, Lots"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Connectors"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Bead Caps"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Chains"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  13,
	}
	result["Clasps & Hooks"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Chains, Necklaces & Pendants"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Charms & Pendants"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  18,
	}
	result["Bracelets"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  12,
	}
	result["Cabochons"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  24,
	}
	result["Carved Figures"] = siteClient.CategoryStruct{
		MainCategory: 11,
		SubCategory:  28,
	}
	result["Denby/Langley/Lovatts"] = siteClient.CategoryStruct{ // dinnerware
		MainCategory: 12,
		SubCategory:  33,
	}
	result["Earring Findings"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Earrings"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  14,
	}
	result["Frames"] = siteClient.CategoryStruct{
		MainCategory: 11,
		SubCategory:  27,
	}
	result["Franciscan"] = siteClient.CategoryStruct{ // dinnerware
		MainCategory: 12,
		SubCategory:  12,
	}
	result["Jewelry Clasps & Hooks"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Jewelry Making Chains"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Metals"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Eggs"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  37,
	}
	result["Other Craft Jewelry Findings"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Other Fine Necklaces, Pendants"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  13,
	}
	result["Other Jewelry Design Findings"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Other Loose Gemstones"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Other Sapphires"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Owls"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  36,
	}
	result["Rhinestones"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Stone"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Trinket Boxes"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  37,
	}
	result["Single Flatware Pieces"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  33,
	}
	result["Buttons"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  39,
	}
	result["Other China & Dinnerware"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  31,
	}
	result["Other French Art Glass"] = siteClient.CategoryStruct{
		MainCategory: 11,
		SubCategory:  28,
	}
	result["Jewelry Boxes"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  37,
	}
	result["Movements"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  40,
	}
	result["Wristwatch Bands"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  40,
	}
	result["Salt & Pepper Shakers"] = siteClient.CategoryStruct{
		MainCategory: 12,
		SubCategory:  41,
	}
	result["Jewelry Sets"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  42,
	}
	result["Brooches, Pins"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  43,
	}
	result["Pins, Brooches"] = siteClient.CategoryStruct{
		MainCategory: 13,
		SubCategory:  43,
	}
	result["Pin Backs & Brooch Components"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  44,
	}
	result["Spacer Beads & Stoppers"] = siteClient.CategoryStruct{
		MainCategory: 14,
		SubCategory:  16,
	}

	return result
}
