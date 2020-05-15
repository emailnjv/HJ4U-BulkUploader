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
				productID: 26078971,
				imageURLs: []string{
					"https://i.ebayimg.com/00/s/MTAyMlgxMDg3/z/loAAAOSwed9dlimX/$_12.JPG?set_id=880000500F",
					"https://i.ebayimg.com/00/s/MTA4MlgxMTgw/z/i9sAAOSw1EZdlim4/$_12.JPG?set_id=880000500F",
					"https://i.ebayimg.com/00/s/MTM2OVgxNTE0/z/EiEAAOSwQvhdlinX/$_12.JPG?set_id=880000500F",
					"https://i.ebayimg.com/00/s/MTYwMFgxMjAw/z/mewAAOSwH7ddlily/$_12.JPG?set_id=880000500F",
					"https://i.ebayimg.com/00/s/MTYwMFgxNjAw/z/pDgAAOSwUMpdlils/$_12.JPG?set_id=880000500F",
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