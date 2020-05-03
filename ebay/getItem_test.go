package ebay

import (
	"testing"
)

func TestEbayClient_GetItem(t *testing.T) {

	type args struct {
		itemID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Successfully Queries Item",
			args{
				"260789710597",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ebayClient, err := NewEbayClient()
			if err != nil {
				t.Errorf("error creating ebay client")
				return
			}

			got, err := ebayClient.getItem(tt.args.itemID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Ack == "Failure" {
				t.Error("Ebay connection failed, check OAuth Key")
			}
		})
	}
}
