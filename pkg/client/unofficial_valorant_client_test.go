package client

import (
	"testing"
)

func TestUnofficialValorantAPIClient_GetWeapons(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "success",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewUnofficialValorantAPIClient()
			_, err := c.GetWeapons()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWeapons() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
