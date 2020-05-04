package utils

import (
	"fmt"
	"testing"
)

func TestNewSCPClient(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"Should create a new SCPClient", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSCPClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSCPClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err = got.client.Close(); (err != nil)  != tt.wantErr {
				t.Errorf("SCPClient().client.Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSCPClient_uploadFile(t *testing.T) {
	type args struct {
		srcImagePath  string
		destImagePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// { "Upload image", args{
		// 	srcImagePath:  "./testImage.jpg",
		// 	destImagePath: "/home/igivnqlrr5nm/public_html/assets/img/destImagePath.jpg",
		// },
		// false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := NewSCPClient()
			defer s.client.Close()

			fmt.Println("b4 sesh")
			session, err := s.client.NewSession()
			if (err != nil) != tt.wantErr {
				t.Errorf("uploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("after session")

			if err = s.uploadFile(session, tt.args.srcImagePath, tt.args.destImagePath); (err != nil) != tt.wantErr {
				t.Errorf("uploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("after upload")

			defer session.Close()
		})
	}
}