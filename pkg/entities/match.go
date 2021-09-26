package entities

import (
	"fmt"
	"math"
)

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

func (m *Match) FindPlayerByPUUID(puuid string) (*Player, error) {
	for _, player := range m.Players.AllPlayers {
		if player.Puuid == puuid {
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

func (m *Match) TierOf(team string) (float64, float64, float64) {
	var players []Player
	max, min, average := float64(math.MinInt), float64(math.MaxInt), 0.0
	if team == "Red" {
		players = m.Players.Red
	} else {
		players = m.Players.Blue
	}
	for _, player := range players {
		max = math.Max(max, float64(player.Currenttier))
		min = math.Min(min, float64(player.Currenttier))
		average += float64(player.Currenttier)
	}
	return average / float64(len(players)), max, min
}

func (m *Match) Tier() (float64, float64, float64) {
	rA, rMax, rMin := m.TierOf("Red")
	bA, bMax, bMin := m.TierOf("Blue")
	return (rA + bA) / 2.0, math.Max(rMax, bMax), math.Min(rMin, bMin)
}

func (m *Match) AttackerOf(round int) string {
	if round < 12 {
		return "Red"
	} else {
		return "Blue"
	}
}

func (m *Match) DefenderOf(round int) string {
	if round < 12 {
		return "Blue"
	} else {
		return "Red"
	}
}
