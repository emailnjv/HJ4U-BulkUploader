package utils

import (
	"testing"
)

func TestNewFTPStruct(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Should create FTP connection", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFTPStruct()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFTPStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err = got.connection.Quit(); (err != nil) != tt.wantErr {
				t.Errorf("uploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}