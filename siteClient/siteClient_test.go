package siteClient

import (
	"reflect"
	"testing"
)

func TestNewSiteClient(t *testing.T) {
	tests := []struct {
		name    string
		want    SiteClient
		wantErr bool
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSiteClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSiteClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSiteClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSiteClient_handleImageURLs(t *testing.T) {

	type args struct {
		productID int
		imageURLs []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ImageArr Upload Test",
			args: args{
				productID: 99999999,
				imageURLs: []string{
					"https://i.ebayimg.com/00/s/MTA1M1gxMjQ2/z/hWEAAOSw9Ztcftdp/$_57.JPG?set_id=8800005007", "https://i.ebayimg.com/00/s/MTA2NlgxNjAw/z/wFwAAOSwo~dcftaF/$_57.JPG?set_id=880000500F", "https://i.ebayimg.com/00/s/MTA2NlgxNjAw/z/iYkAAOSwJoxcftZ~/$_57.JPG?set_id=880000500F", "https://i.ebayimg.com/00/s/MTA2NlgxNjAw/z/waQAAOSwqrVcftaE/$_57.JPG?set_id=880000500F", "https://i.ebayimg.com/00/s/MTA2NlgxNjAw/z/SpQAAOSwI7JcftaA/$_57.JPG?set_id=880000500F", "https://i.ebayimg.com/00/s/OTUyWDEyNzg=/z/lu8AAOSwUMhcftbL/$_57.JPG?set_id=8800005007", "https://i.ebayimg.com/00/s/MTA2NlgxNjAw/z/3J8AAOSwVbhcftaC/$_57.JPG?set_id=880000500F", "https://i.ebayimg.com/00/s/MTA2NlgxNjAw/z/4IYAAOSwzOxcftaf/$_57.JPG?set_id=880000500F",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sc, err := NewSiteClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSiteClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := sc.handleImageURLs(tt.args.productID, tt.args.imageURLs); (err != nil) != tt.wantErr {
				t.Errorf("handleImageURLs() error = %v, wantErr %v", err, tt.wantErr)
			}

			err = sc.CloseSiteClient()
			if (err != nil) != tt.wantErr {
				t.Errorf("CloseSiteClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
