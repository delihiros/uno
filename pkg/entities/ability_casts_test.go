package entities

import (
	"reflect"
	"testing"
)

func TestAbilityCasts_UnmarshalJSON(t *testing.T) {
	type fields struct {
		CCast int
		QCast int
		ECast int
		XCast int
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *AbilityCasts
		wantErr bool
	}{
		{
			name: "parse normal",
			args: args{
				data: []byte(`{"c_cast":1,"q_cast":2,"e_cast":3,"x_cast":4}`),
			},
			want: &AbilityCasts{
				CCast: 1,
				QCast: 2,
				ECast: 3,
				XCast: 4,
			},
			wantErr: false,
		},
		{
			name: "parse null",
			args: args{
				data: []byte(`{"c_cast":null,"q_cast":null,"e_cast":null,"x_cast":null}`),
			},
			want: &AbilityCasts{
				CCast: 0,
				QCast: 0,
				ECast: 0,
				XCast: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ac := &AbilityCasts{}
			if err := ac.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(ac, tt.want) {
				t.Errorf("UnmarshalJSON() want = %v, got %v", tt.want, ac)
			}
		})
	}
}
