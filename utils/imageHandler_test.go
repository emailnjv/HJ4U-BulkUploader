package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestImageHandler_DownloadImage(t *testing.T) {
	type args struct {
		imagePath string
		imageURL  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Download Image",
			args{
				imagePath: "./testImage.jpg",
				imageURL:  "https://i.ebayimg.com/00/s/MTAyMlgxMDg3/z/loAAAOSwed9dlimX/$_12.JPG?set_id=880000500F",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := ImageHandler{}
			got, err := i.DownloadImage(tt.args.imagePath, tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("DownloadImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestNewImageHandler(t *testing.T) {
	tests := []struct {
		name string
		want ImageHandler
	}{
		{"NewImageHandler base test", ImageHandler{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewImageHandler(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewImageHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
