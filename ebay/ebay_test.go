package ebay

import (
	"fmt"
	"testing"
)

func TestNewEbayClient(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Instantiate A Ebay Client"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ec, err := NewEbayClient()
			if err != nil {
				t.Errorf("Error creating new Ebay client\n%#v", err)
			}
			fmt.Printf("%#v\n", ec)
		})
	}
}
