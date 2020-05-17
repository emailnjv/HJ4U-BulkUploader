package utils

import (
	"testing"
)

func Test_parseHTML(t *testing.T) {
	type args struct {
		htmlString string
	}
	tests := []struct {
		name string
		args args
		want string
		wantErr bool
	}{
		{"fontDivUlLi", args{htmlString: `<font size="4" style="font-family:Arial" rwr="1">
<div><ul><li>Vasily Kandinsky "Little Painting With Yellow", 1914 Poster</li><li>Philadelphia Museum Of Art</li><li>Published &amp; Distributed By Graphique De France Printed In France</li></ul></div>




</font>`}, `<ul style="font-family:Arial;font-size:14px;"><li>Vasily Kandinsky "Little Painting With Yellow", 1914 Poster</li><li>Philadelphia Museum Of Art</li><li>Published & Distributed By Graphique De France Printed In France</li></ul>`, false},
		{"divDivb", args{htmlString: `<div>
<div><b>Henri Maltisse Lithograph</b></div><div><b>(1869-1954)</b></div><div><b>Christmas Eve</b></div><div><b>11" X 14"&nbsp;</b></div><div><b>Color Lithograph</b></div><div><b>42.5 X 65 cm</b></div><div><b>From Verve</b></div><div><b>Issue # 35/36</b></div><div><b>Published By Graphique Di France</b></div><div><b>Printed In France</b></div><div><b>1994 Succession H Maltisse / A.R.S., N.Y. LP 122</b></div>
</div>`}, `<ul style="font-family:Arial;font-size:14px;"><li>Henri Maltisse Lithograph</li><li>(1869-1954)</li><li>Christmas Eve</li><li>11" X 14" </li><li>Color Lithograph</li><li>42.5 X 65 cm</li><li>From Verve</li><li>Issue # 35/36</li><li>Published By Graphique Di France</li><li>Printed In France</li><li>1994 Succession H Maltisse / A.R.S., N.Y. LP 122</li></ul>`, false},
{"divTable", args{htmlString: `<font rwr='1' size='4' style='font-family:Arial'>
<div>
<table cellpadding="0" cellspacing="0"><tbody><tr>



</tr><tr><td>&nbsp;</td>
<td colspan="2">
<div><ul><li>NEW ELVIS PRESLEY POSTER BLACK &amp; WHITE NYC JULY,1956</li></ul></div>
<table><tbody><tr><td> 
<table><tbody><tr><td>
<table cellpadding="8"><tbody><tr><td>
<p><ul><li><b>SIZE 15.5" X 12"</b></li></ul></p><p><br></p><p><ul><li><b>ELVIS REHEARSES FOR THE STEVEN ALLEN TV SHOW. NYC JULY 1, 1956</b></li><li><b>"YOU AIN'T NOTHIN' BUT A HOUND DOG"</b></li><li><b>PHOTOGRAPH BY ALFRED WERTHEIMER, 1979</b></li><li><u><strong>WRITTEN&nbsp;BELOW ELVIS PICTURE:</strong></u></li><li><b>ALL RIGHTS RESERVED./DAUMAN PICTURES, NYC. PRODUCED BY CINERGY GRAPHICS, NYC 1996.</b></li><li><b>ELVIS AND ELVIS PRESLEY ARE REGISTERED TRADEMARKS OF ELVIS PRESLEY ENTERPRISES, INC.</b></li><li><b>LP420&nbsp;EXCLUSIVELY&nbsp;DISTRIBUTED BY GRAPHIQUE DE FRANCE</b></li></ul></p><p><ul><li><strong>PRINTED IN FRANCE</strong></li></ul></p></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table></td></tr></tbody></table></div>




</font>`}, `<ul style="font-family:Arial;font-size:14px;"><li>NEW ELVIS PRESLEY POSTER BLACK & WHITE NYC JULY,1956</li><li>SIZE 15.5" X 12"</li><li>ELVIS REHEARSES FOR THE STEVEN ALLEN TV SHOW. NYC JULY 1, 1956</li><li>"YOU AIN'T NOTHIN' BUT A HOUND DOG"</li><li>PHOTOGRAPH BY ALFRED WERTHEIMER, 1979</li><li>WRITTEN BELOW ELVIS PICTURE:</li><li>ALL RIGHTS RESERVED./DAUMAN PICTURES, NYC. PRODUCED BY CINERGY GRAPHICS, NYC 1996.</li><li>ELVIS AND ELVIS PRESLEY ARE REGISTERED TRADEMARKS OF ELVIS PRESLEY ENTERPRISES, INC.</li><li>LP420 EXCLUSIVELY DISTRIBUTED BY GRAPHIQUE DE FRANCE</li><li>PRINTED IN FRANCE</li></ul>`, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := HTMLParser{}
			result, err := parser.ParseHTML(tt.args.htmlString)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateThumbnailFromJPG() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if result != tt.want {
				t.Errorf("parseHTML() got = %v, want %v", result, tt.want)
			}		})
	}
}
