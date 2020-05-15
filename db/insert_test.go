package db

import (
	"fmt"
	"testing"
	"time"
)

var testProduct Product = Product{
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
		product *Product
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
