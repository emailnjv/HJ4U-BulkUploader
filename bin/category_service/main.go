package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/siteClient"
)

var categories = map[string]bool{
	"Beads":                          false,
	"Ceramic, Clay, Porcelain":       false,
	"Art Posters Done":               false,
	"Art Prints":                     false,
	"Collections, Lots":              false,
	"Connectors":                     false,
	"Bead Caps":                      false,
	"Chains":                         false,
	"Clasps & Hooks":                 false,
	"Chains, Necklaces & Pendants":   false,
	"Charms & Pendants":              false,
	"Bracelets":                      false,
	"Cabochons":                      false,
	"Carved Figures":                 false,
	"Denby/Langley/Lovatts":          false,
	"Earring Findings":               false,
	"Earrings":                       false,
	"Frames":                         false,
	"Franciscan":                     false,
	"Jewelry Clasps & Hooks":         false,
	"Jewelry Making Chains":          false,
	"Metals":                         false,
	"Eggs":                           false,
	"Other Craft Jewelry Findings":   false,
	"Other Fine Necklaces, Pendants": false,
	"Other Jewelry Design Findings":  false,
	"Other Loose Gemstones":          false,
	"Other Sapphires":                false,
	"Owls":                           false,
	"Rhinestones":                    false,
	"Stone":                          false,
	"Trinket Boxes":                  false,
	"Single Flatware Pieces":         false,
	"Buttons":                        false,
	"Other China & Dinnerware":       false,
	"Other French Art Glass":         false,
	"Jewelry Boxes":                  false,
	"Movements":                      false,
	"Wristwatch Bands":               false,
	"Salt & Pepper Shakers":          false,
	"Jewelry Sets":                   false,
	"Brooches, Pins":                 false,
	"Pins, Brooches":                 false,
	"Pin Backs & Brooch Components":  false,
	"Spacer Beads & Stoppers":        false,
	"Limoges":                        true,
}

func init() {
}

func main() {
	// for err := range runGroupedRespDownload(categories, "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/LimogesResponses") {
	// 	fmt.Printf("Error returned from runGroupedRespDownload: %#v\n", err)
	// }
	// fmt.Println(uploadLocalListings())
	fmt.Println(uploadSpecificLocalListings("Fine Porcelain", "Limoges"))

	fmt.Println("Finished")
}

func uploadLocalListings() error {
	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}
	categoryMapping := sampleMapping()
	// NOTICE: Pulls in the category mapping here
	return sc.UploadLocalListings("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/limogesResponses", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/pictures", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/limogesVarianceResponses", categoryMapping)
}
func uploadSpecificLocalListings(targetCategory string, targetCSVCategory string) error {
	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}
	// categoryMapping := sampleMapping()
	// NOTICE: Pulls in the category mapping here
	return sc.UploadSpecificLocalListings("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/limogesResponses", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/pictures", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/limogesVarianceResponses", targetCategory, targetCSVCategory)
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
			if useBool, ok := categories[line[14]]; ok && useBool {
				IDArr = append(IDArr, line[0])
			}
		}
	}

	file, _ := json.MarshalIndent(IDArr, "", " ")

	_ = ioutil.WriteFile("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/ebay_listings/targetIDs.json", file, 0644)

	return sc.EbayClient.DownloadAllResponses(IDArr, downloadDirectory)
}

func exit(err error) {
	if err == nil {
		os.Exit(0)
	}
	os.Exit(1)
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

func sampleMapping() map[string]db.CategoryIDStruct {
	result := make(map[string]db.CategoryIDStruct)
	result["Beads"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  16,
	}
	result["Ceramic, Clay, Porcelain"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  16,
	}
	result["Art Posters Done"] = db.CategoryIDStruct{
		MainCategoryID: 11,
		SubCategoryID:  26,
	}
	result["Art Prints"] = db.CategoryIDStruct{
		MainCategoryID: 11,
		SubCategoryID:  26,
	}
	result["Collections, Lots"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Connectors"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Bead Caps"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Chains"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  13,
	}
	result["Clasps & Hooks"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Chains, Necklaces & Pendants"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  21,
	}
	result["Charms & Pendants"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  18,
	}
	result["Bracelets"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  12,
	}
	result["Cabochons"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  24,
	}
	result["Carved Figures"] = db.CategoryIDStruct{
		MainCategoryID: 11,
		SubCategoryID:  28,
	}
	result["Denby/Langley/Lovatts"] = db.CategoryIDStruct{ // dinnerware
		MainCategoryID: 12,
		SubCategoryID:  33,
	}
	result["Earring Findings"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Earrings"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  14,
	}
	result["Frames"] = db.CategoryIDStruct{
		MainCategoryID: 11,
		SubCategoryID:  27,
	}
	result["Franciscan"] = db.CategoryIDStruct{ // dinnerware
		MainCategoryID: 12,
		SubCategoryID:  12,
	}
	result["Jewelry Clasps & Hooks"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  21,
	}
	result["Jewelry Making Chains"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  21,
	}
	result["Metals"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  21,
	}
	result["Eggs"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  37,
	}
	result["Other Craft Jewelry Findings"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Other Fine Necklaces, Pendants"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  13,
	}
	result["Other Jewelry Design Findings"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  19,
	}
	result["Other Loose Gemstones"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  17,
	}
	result["Other Sapphires"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  17,
	}
	result["Owls"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  36,
	}
	result["Rhinestones"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  17,
	}
	result["Stone"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  17,
	}
	result["Trinket Boxes"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  37,
	}
	result["Single Flatware Pieces"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  33,
	}
	result["Buttons"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  39,
	}
	result["Other China & Dinnerware"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  31,
	}
	result["Other French Art Glass"] = db.CategoryIDStruct{
		MainCategoryID: 11,
		SubCategoryID:  28,
	}
	result["Jewelry Boxes"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  37,
	}
	result["Movements"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  40,
	}
	result["Wristwatch Bands"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  40,
	}
	result["Salt & Pepper Shakers"] = db.CategoryIDStruct{
		MainCategoryID: 12,
		SubCategoryID:  41,
	}
	result["Jewelry Sets"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  42,
	}
	result["Brooches, Pins"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  43,
	}
	result["Pins, Brooches"] = db.CategoryIDStruct{
		MainCategoryID: 13,
		SubCategoryID:  43,
	}
	result["Pin Backs & Brooch Components"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  44,
	}
	result["Spacer Beads & Stoppers"] = db.CategoryIDStruct{
		MainCategoryID: 14,
		SubCategoryID:  16,
	}

	return result
}
