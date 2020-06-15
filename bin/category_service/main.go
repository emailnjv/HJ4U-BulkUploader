package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

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
	fmt.Println(uploadLocalListings())

	fmt.Println("Finished")
}

func uploadLocalListings() error {
	sc, err := siteClient.NewSiteClient()
	if err != nil {
		return err
	}
	// categoryMapping := sampleMapping()
	// NOTICE: Pulls in the category mapping here
	return sc.UploadSpecificLocalListings("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/limogesResponses", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/pictures", "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/data/limogesVarianceResponses", "Limoges")
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
