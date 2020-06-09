package siteClient

import (
	"fmt"
	"testing"

	"github.com/emailnjv/HJ4U-BulkUploader/db"
	"github.com/emailnjv/HJ4U-BulkUploader/utils"
)

// func TestSiteClient_UploadLocalListings(t *testing.T) {
// 	type args struct {
// 		listingDirectory string
// 		imageDirectory   string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{"Test", args{
// 			listingDirectory: "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/resp",
// 			imageDirectory:   "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/mergedProducts",
// 		}, false},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			sc, err := NewSiteClient()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("NewSiteClient() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
//
// 			sample := sampleMapping()
//
// 			if err := sc.UploadLocalListings(tt.args.listingDirectory, tt.args.imageDirectory, sample); (err != nil) != tt.wantErr {
// 				t.Errorf("UploadLocalListings() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestVarianceParser_HandleVariances(t *testing.T) {
	type args struct {
		productID        int
		varianceResponse []utils.GIGRItems
	}
	tests := []struct {
		name    string
		args    args
		want    []*db.ProductAtt
		wantErr bool
	}{
		{
			name: "Variance Parser Test",
			args: args{
				productID: 1124,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, err := NewSiteClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleVariances() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// map xml variance responses to listings
			for jsonErrStruct := range sc.ReadVarDir("/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/debugData/varianceResponses") {
				if jsonErrStruct.Error != nil {
					t.Errorf("ReadVarDir() error = %v, wantErr %v", err, tt.wantErr)
					return
				} else {
					got, err := sc.VarianceParser.HandleVariances(tt.args.productID, jsonErrStruct.Response.Items)
					if (err != nil) != tt.wantErr {
						t.Errorf("HandleVariances() error = %v, wantErr %v", err, tt.wantErr)
						return
					}
					fmt.Printf("%#v", got)
				}
			}
		})
	}
}

func sampleMapping() map[string]CategoryStruct {
	result := make(map[string]CategoryStruct)
	result["Beads"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  16,
	}
	result["Ceramic, Clay, Porcelain"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  16,
	}
	result["Art Posters Done"] = CategoryStruct{
		MainCategory: 11,
		SubCategory:  26,
	}
	result["Art Prints"] = CategoryStruct{
		MainCategory: 11,
		SubCategory:  26,
	}
	result["Collections, Lots"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Connectors"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Bead Caps"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Chains"] = CategoryStruct{
		MainCategory: 13,
		SubCategory:  13,
	}
	result["Clasps & Hooks"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Chains, Necklaces & Pendants"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Charms & Pendants"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  18,
	}
	result["Bracelets"] = CategoryStruct{
		MainCategory: 13,
		SubCategory:  12,
	}
	result["Cabochons"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  24,
	}
	result["Carved Figures"] = CategoryStruct{
		MainCategory: 11,
		SubCategory:  28,
	}
	result["Denby/Langley/Lovatts"] = CategoryStruct{ // dinnerware
		MainCategory: 12,
		SubCategory:  33,
	}
	result["Earring Findings"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Earrings"] = CategoryStruct{
		MainCategory: 13,
		SubCategory:  14,
	}
	result["Frames"] = CategoryStruct{
		MainCategory: 11,
		SubCategory:  27,
	}
	result["Franciscan"] = CategoryStruct{ // dinnerware
		MainCategory: 12,
		SubCategory:  12,
	}
	result["Jewelry Clasps & Hooks"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Jewelry Making Chains"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Metals"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  21,
	}
	result["Eggs"] = CategoryStruct{
		MainCategory: 12,
		SubCategory:  37,
	}
	result["Other Craft Jewelry Findings"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Other Fine Necklaces, Pendants"] = CategoryStruct{
		MainCategory: 13,
		SubCategory:  13,
	}
	result["Other Jewelry Design Findings"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  19,
	}
	result["Other Loose Gemstones"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Other Sapphires"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Owls"] = CategoryStruct{
		MainCategory: 12,
		SubCategory:  36,
	}
	result["Rhinestones"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Stone"] = CategoryStruct{
		MainCategory: 14,
		SubCategory:  17,
	}
	result["Trinket Boxes"] = CategoryStruct{
		MainCategory: 12,
		SubCategory:  37,
	}

	return result
}
