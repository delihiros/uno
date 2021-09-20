package entities

import "fmt"

type Match struct {
	Metadata struct {
		Map              string `json:"map"`
		GameVersion      string `json:"game_version"`
		GameLength       int    `json:"game_length"`
		GameStart        int64  `json:"game_start"`
		GameStartPatched string `json:"game_start_patched"`
		RoundsPlayed     int    `json:"rounds_played"`
		Mode             string `json:"mode"`
		SeasonID         string `json:"season_id"`
		Platform         string `json:"platform"`
		Matchid          string `json:"matchid"`
	} `json:"metadata"`
	Players struct {
		AllPlayers []Player `json:"all_players"`
		Red        []Player `json:"red"`
		Blue       []Player `json:"blue"`
	} `json:"players"`
	Teams struct {
		Red  Team `json:"red"`
		Blue Team `json:"blue"`
	} `json:"teams"`
	Rounds []Round `json:"rounds"`
}

func (m *Match) FindPlayer(name, tag string) (*Player, error) {
	for _, player := range m.Players.AllPlayers {
		if player.Name == name && player.Tag == tag {
			return &player, nil
		}
	}
	return nil, fmt.Errorf("could not find player")
}

func (m *Match) NameTag(puuid string) (string, string, error) {
	for _, player := range m.Players.AllPlayers {
		if player.Puuid == puuid {
			return player.Name, player.Tag, nil
		}
	}
	return "", "", fmt.Errorf("could not find player")
}
