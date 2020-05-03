package ebay

import (
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
			_, err := NewEbayClient()
			if err != nil {
				t.Errorf("Error creating new Ebay client")
			}
		})
	}
}