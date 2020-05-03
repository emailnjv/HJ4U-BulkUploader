package db

import (
	"testing"
)

func TestNewTargetDBClientAndClose(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"Should be able to connect, and then close it",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTargetDBClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			err = got.CloseConnection()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewDBClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
