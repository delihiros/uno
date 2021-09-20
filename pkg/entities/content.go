package entities

import "time"

type Content struct {
	Characters []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Characters"`
	Maps []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Maps"`
	Chromas []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Chromas"`
	Skins []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Skins"`
	SkinLevels []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"SkinLevels"`
	Attachments []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Attachments"`
	Equips []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Equips"`
	Themes []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Themes"`
	GameModes []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"GameModes"`
	Sprays []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Sprays"`
	SprayLevels []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"SprayLevels"`
	Charms []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"Charms"`
	CharmLevels []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"CharmLevels"`
	PlayerCards []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"PlayerCards"`
	PlayerTitles []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"PlayerTitles"`
	StorefrontItems []struct {
		Name      string `json:"Name"`
		ID        string `json:"ID"`
		AssetName string `json:"AssetName"`
		IsEnabled bool   `json:"IsEnabled"`
	} `json:"StorefrontItems"`
	Seasons []struct {
		ID              string    `json:"ID"`
		Name            string    `json:"Name"`
		Type            string    `json:"Type"`
		StartTime       time.Time `json:"StartTime"`
		EndTime         time.Time `json:"EndTime"`
		IsEnabled       bool      `json:"IsEnabled"`
		IsActive        bool      `json:"IsActive"`
		DevelopmentOnly bool      `json:"DevelopmentOnly"`
	} `json:"Seasons"`
	CompetitiveSeasons []struct {
		ID              string    `json:"ID"`
		SeasonID        string    `json:"SeasonID"`
		StartTime       time.Time `json:"StartTime"`
		EndTime         time.Time `json:"EndTime"`
		DevelopmentOnly bool      `json:"DevelopmentOnly"`
	} `json:"CompetitiveSeasons"`
	Events []struct {
		ID              string    `json:"ID"`
		Name            string    `json:"Name"`
		StartTime       time.Time `json:"StartTime"`
		EndTime         time.Time `json:"EndTime"`
		IsEnabled       bool      `json:"IsEnabled"`
		IsActive        bool      `json:"IsActive"`
		DevelopmentOnly bool      `json:"DevelopmentOnly"`
	} `json:"Events"`
}
