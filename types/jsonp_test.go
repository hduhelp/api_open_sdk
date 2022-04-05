package types_test

import (
	"testing"

	"github.com/gin-gonic/gin/render"
	"github.com/hduhelp/api_open_sdk/types"
)

func TestJsonp_UnmarshalJSON(t *testing.T) {
	type fields struct {
		JsonpJSON render.JsonpJSON
	}
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_jsonp",
			fields: fields{
				JsonpJSON: render.JsonpJSON{
					Callback: "jQuery",
					Data: map[string]interface{}{
						"str": "8cec71817787de6fd05",
						"num": 1637719161,
						"array": []interface{}{
							"8cec71817787de6fd05",
							1637719161,
						},
					},
				},
			},
			args: args{
				bytes: []byte(`jQuery({
	"str": "8cec71817787de6fd05",
	"num":        1637719161,
	"array": [
        "8cec71817787de6fd05",
        1637719161
    ]
})`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &types.Jsonp{
				JsonpJSON: tt.fields.JsonpJSON,
			}
			if err := j.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
