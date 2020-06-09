package siteClient

import (
	"testing"
)

func TestSiteClient_DownloadAllImages(t *testing.T) {
	type args struct {
		responseDirectory string
		targetDirectory   string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Test", args: args{
			responseDirectory: "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/resp",
			targetDirectory:   "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/prod",
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, err := NewSiteClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSiteClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			sc.DownloadAllImages(tt.args.responseDirectory, tt.args.targetDirectory)

			// if err := sc.DownloadAllImages(tt.args.responseDirectory, tt.args.targetDirectory); (err != nil) != tt.wantErr {
			// 	t.Errorf("DownloadAllImages() error = %v, wantErr %v", err, tt.wantErr)
			// }
		})
	}
}

func TestSiteClient_DownloadAllIDs(t *testing.T) {
	type args struct {
		responseDirectory string
		targetDirectory   string
	}
	tests := []struct {
		name   string
		args   args
		wantErr bool
	}{
		{"ID download test", args{
			responseDirectory: "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/BulkUploader/resources/resp",
			targetDirectory:   "/home/nick/Documents/Projects/Work/Dad/HotJewelry4U/ebay_listings",
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, err := NewSiteClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSiteClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			sc.DownloadAllIDs(tt.args.responseDirectory, tt.args.targetDirectory)
		})
	}
}
