package db

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
)

var testProduct Products = Products{
	Name:         "Test Product",
	Description:  "Lorem ipsum is placeholder text commonly used in the graphic, print, and publishing industries for previewing layouts and visual mockups.",
	Price:        22.15,
	Featured:     0,
	Main_cat:     13,
	Sub_cat:      12,
	Qty:          10,
	Sku:          "260789710597",
	Upc:          "260789710597",
	Product_type: "simple",
}

func TestTargetDBClient_insertProduct(t *testing.T) {
	now := time.Now()
	testProduct.Date = &now
	type args struct {
		product *Products
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Insert Product",
			args:    args{product: &testProduct},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := NewTargetDBClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTargetDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := d.InsertProduct(tt.args.product)
			if (err != nil) != tt.wantErr {
				t.Errorf("insertProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if recordNew := d.db.NewRecord(*tt.args.product); recordNew != false {
				t.Errorf("NewRecord() = %v, want %v", got, false)
			}

			fmt.Println(got)

			deleteRef := d.db.Delete(*tt.args.product)
			if (deleteRef.Error != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			err = d.CloseConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("CloseConnection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestTargetDBClient_LevenshteinRatio(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		targetWord1 string
		targetWord2 string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "LevenshteinRatio Word Passes",
			args: args{
				targetWord1: "Cabochon",
				targetWord2: "Cabachon",
			},
			want:    88,
			wantErr: false,
		},
		{
			name: "LevenshteinRatio Word Fails",
			args: args{
				targetWord1: "Cabochon",
				targetWord2: "Applesauce",
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d, err := NewTargetDBClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTargetDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got, err := d.LevenshteinRatio(tt.args.targetWord1, tt.args.targetWord2)
			if (err != nil) != tt.wantErr {
				t.Errorf("LevenshteinRatio() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("LevenshteinRatio() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBClient_MainCatsSubCats(t *testing.T) {
	type args struct {
		mainCategoriesName string
	}
	tests := []struct {
		name string
		args args
		want    CategoryIDMapping
		wantErr bool
	}{
		{
			name: "MainCatsSubCats Retrieval",
			args: args{
				mainCategoriesName: "Fine Porcelain",
			},
			want: CategoryIDMapping{
				"Baskets":{
					12,
					29,
				},
				"Cache Pots":{
					12,
					30,
				},
				"Candle Holders":{
					12,
					31,
				},
				"Candy Dishes / Trays":{
					12,
					32,
				},
				"Dinnerware":{
					12,
					33,
				},
				"Figurines":{
					12,
					36,
				},
				"Jars":{
					12,
					34,
				},
				"Pin Cushions":{
					12,
					35,
				},
				"Salt & Pepper Shakers":{
					12,
					41,
				},
				"Trinket Boxes":{
					12,
					37,
				},
				"Vases":{
					12,
					38,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbc, err := NewTargetDBClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTargetDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := dbc.MainCatsSubCats(tt.args.mainCategoriesName)
			if (err != nil) != tt.wantErr {
				t.Errorf("MainCatsSubCats() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MainCatsSubCats() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBClient_LevenshteinMatch(t *testing.T) {

	type args struct {
		input        string
		categoryName string
		targetMatch  int
	}
	tests := []struct {
		name    string
		args    args
		want    CategoryIDStruct
		wantErr bool
	}{
		{
			name: "LevenshteinMatch",
			args: args{
				input:        "Cabachon",
				categoryName: "Jewelry Supplies",
				targetMatch:  80,
			},
			want: CategoryIDStruct{
				MainCategoryID: 14,
				SubCategoryID:  24,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbc, err := NewTargetDBClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTargetDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := dbc.LevenshteinMatch(tt.args.input, tt.args.categoryName, tt.args.targetMatch)
			if (err != nil) != tt.wantErr {
				t.Errorf("LevenshteinMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevenshteinMatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDBClient_SQLMatchMatch(t *testing.T) {

	type args struct {
		input          string
		targetCategory string
	}
	tests := []struct {
		name    string
		args    args
		want    []CategoryIDStruct
		wantErr bool
	}{
		{
			name: "SQLMatchMatch",
			args: args{
				input:          "pots",
				targetCategory: "Fine Porcelain",
			},
			want:    []CategoryIDStruct{
				{
					MainCategoryID: 12,
					SubCategoryID:  30,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbc, err := NewTargetDBClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTargetDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := dbc.SQLMatchMatch(tt.args.input, tt.args.targetCategory)
			if (err != nil) != tt.wantErr {
				t.Errorf("SQLMatchMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SQLMatchMatch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
