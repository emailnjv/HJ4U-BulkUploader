package utils

import (
	"bufio"
	"os"
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
			if err = got.CloseClients(); (err != nil)  != tt.wantErr {
				t.Errorf("SCPClient().client.Close() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSCPClient_UploadFile(t *testing.T) {
	// Load the image to upload
	file, _ := os.Open("./testImage.jpg")
	defer file.Close()


	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(file)
	_, err := buffer.Read(bytes)
	if err != nil {
		t.Error(err)
	}

	type args struct {
		srcImage  []byte
		destImagePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{ "Upload image", args{
			srcImage: bytes,
			destImagePath: "/home/igivnqlrr5nm/public_html/assets/img/destImagePath.jpg",
		},
		false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			s, err := NewSCPClient()
			defer s.CloseClients()

			if err = s.UploadFile( tt.args.srcImage, tt.args.destImagePath); (err != nil) != tt.wantErr {
				t.Errorf("UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err = s.DeleteFile(tt.args.destImagePath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}