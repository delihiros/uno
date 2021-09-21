package client

import (
	"reflect"
	"testing"
	"uno/pkg/entities"
)

func TestClient_GetMatchByID(t *testing.T) {
	type args struct {
		matchID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get match success",
			args: args{
				matchID: "2aa59334-e53a-415b-bb3d-4832305ee7db",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			_, err := c.GetMatchByID(tt.args.matchID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetAccountByNameTag(t *testing.T) {
	type args struct {
		name string
		tag  string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.Account
		wantErr bool
	}{
		{
			name: "get account success",
			args: args{
				name: "bobobobobobobo",
				tag:  "1212",
			},
			want: &entities.Account{
				Puuid: "1e98d5ed-2c63-5573-a564-110ddef7853f",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got, err := c.GetAccountByNameTag(tt.args.name, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByNameTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Puuid, tt.want.Puuid) {
				t.Errorf("GetAccountByNameTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetMMRDataByNameTag(t *testing.T) {
	type args struct {
		region string
		name   string
		tag    string
	}
	tests := []struct {
		name    string
		args    args
		want    *entities.MMRData
		wantErr bool
	}{
		{
			name: "get mmr data success",
			args: args{
				region: "ap",
				name:   "bobobobobobobo",
				tag:    "1212",
			},
			want: &entities.MMRData{
				Name: "bobobobobobobo",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			got, err := c.GetMMRDataByNameTag(tt.args.region, tt.args.name, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMMRDataByNameTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Name, tt.want.Name) {
				t.Errorf("GetMMRDataByNameTag() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetMMRHistory(t *testing.T) {
	type args struct {
		region string
		name   string
		tag    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get mmrhistory success",
			args: args{
				region: "ap",
				name:   "bobobobobobobo",
				tag:    "1212",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			_, err := c.GetMMRHistoryByNameTag(tt.args.region, tt.args.name, tt.args.tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMMRHistoryByNameTag() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetMatchHistory(t *testing.T) {
	type args struct {
		region string
		name   string
		tag    string
		filter string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get match history success",
			args: args{
				region: "ap",
				name:   "bobobobobobobo",
				tag:    "1212",
				filter: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			_, err := c.GetMatchHistory(tt.args.region, tt.args.name, tt.args.tag, tt.args.filter)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetContent(t *testing.T) {
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
			c := New()
			_, err := c.GetContent()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestClient_GetLeaderboard(t *testing.T) {
	type args struct {
		region string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				region: "ap",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			_, err := c.GetLeaderboard(tt.args.region)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLeaderboard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
