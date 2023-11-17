package netbox

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_handleCustomFieldUpdate(t *testing.T) {
	type args struct {
		old interface{}
		new interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "create new custom field",
			args: args{
				old: "",
				new: "{\"a\": \"b\", \"c\": true, \"d\": {\"e\": \"f\"}}",
			},
			want: map[string]interface{}{
				"a": "b",
				"c": true,
				"d": map[string]interface{}{
					"e": "f",
				},
			},
			wantErr: false,
		},
		{
			name: "update custom fields",
			args: args{
				old: "{\"a\": \"b\", \"c\": true, \"d\": {\"e\": \"f\"}}",
				new: "{\"a\": \"q\", \"c\": false}",
			},
			want: map[string]interface{}{
				"a": "q",
				"c": false,
				"d": nil,
			},
			wantErr: false,
		},
		{
			name: "remove custom fields",
			args: args{
				old: "{\"a\": \"b\", \"c\": true, \"d\": {\"e\": \"f\"}}",
				new: "",
			},
			want: map[string]interface{}{
				"a": nil,
				"c": nil,
				"d": nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handleCustomFieldUpdate(tt.args.old, tt.args.new)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleCustomFieldUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_handleCustomFieldRead(t *testing.T) {
	tests := []struct {
		name    string
		cf      interface{}
		want    string
		wantErr bool
	}{
		{
			name: "all fields are nil",
			cf: map[string]interface{}{
				"a": nil,
				"c": nil,
				"d": nil,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "one field is valid",
			cf: map[string]interface{}{
				"a": nil,
				"c": true,
				"d": nil,
			},
			want:    "{\"a\":null,\"c\":true,\"d\":null}",
			wantErr: false,
		},
		{
			name:    "cannot marshal nil",
			cf:      nil,
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handleCustomFieldRead(tt.cf)
			if (err != nil) != tt.wantErr {
				t.Errorf("handleCustomFieldRead() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
