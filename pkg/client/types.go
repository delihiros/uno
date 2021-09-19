package client

import (
	"encoding/json"
	"time"
)

type Response struct {
	Status string          `json:"status"`
	Data   json.RawMessage `json:"data"`
}

type Account struct {
	Puuid        string `json:"puuid"`
	Region       string `json:"region"`
	AccountLevel int    `json:"account_level"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
}

type MMRData struct {
	Name        string `json:"name"`
	Tag         string `json:"tag"`
	CurrentData struct {
		Currenttier          int    `json:"currenttier"`
		Currenttierpatched   string `json:"currenttierpatched"`
		RankingInTier        int    `json:"ranking_in_tier"`
		MmrChangeToLastGame  int    `json:"mmr_change_to_last_game"`
		Elo                  int    `json:"elo"`
		GamesNeededForRating int    `json:"games_needed_for_rating"`
	} `json:"current_data"`
	BySeason struct {
		E3A2 SeasonMMR `json:"e3a2"`
		E3A1 SeasonMMR `json:"e3a1"`
		E2A3 SeasonMMR `json:"e2a3"`
		E2A2 SeasonMMR `json:"e2a2"`
		E2A1 SeasonMMR `json:"e2a1"`
		E1A3 SeasonMMR `json:"e1a3"`
		E1A2 SeasonMMR `json:"e1a2"`
		E1A1 SeasonMMR `json:"e1a1"`
	} `json:"by_season"`
}

type MMRHistory struct {
	Currenttier         int    `json:"currenttier"`
	Currenttierpatched  string `json:"currenttierpatched"`
	RankingInTier       int    `json:"ranking_in_tier"`
	MmrChangeToLastGame int    `json:"mmr_change_to_last_game"`
	Elo                 int    `json:"elo"`
	Date                string `json:"date"`
	DateRaw             int64  `json:"date_raw"`
}

type SeasonMMR struct {
	Wins             int    `json:"wins"`
	NumberOfGames    int    `json:"number_of_games"`
	FinalRank        int    `json:"final_rank"`
	FinalRankPatched string `json:"final_rank_patched"`
	ActRankWins      []struct {
		PatchedTier string `json:"patched_tier"`
		Tier        int    `json:"tier"`
	} `json:"act_rank_wins"`
}

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

type SimplePlayer struct {
	DisplayName string `json:"display_name"`
	Team        string `json:"team"`
}

type Player struct {
	Puuid              string `json:"puuid"`
	Name               string `json:"name"`
	Tag                string `json:"tag"`
	Team               string `json:"team"`
	Character          string `json:"character"`
	Currenttier        int    `json:"currenttier"`
	CurrenttierPatched string `json:"currenttier_patched"`
	PlayerCard         string `json:"player_card"`
	PlayerTitle        string `json:"player_title"`
	Stats              struct {
		Score   int `json:"score"`
		Kills   int `json:"kills"`
		Deaths  int `json:"deaths"`
		Assists int `json:"assists"`
	} `json:"stats"`
	AbilityCasts   AbilityCasts `json:"ability_casts"`
	DamageMade     int          `json:"damage_made"`
	DamageReceived int          `json:"damage_received"`
}

type AbilityCasts struct {
	CCast int `json:"c_cast"`
	QCast int `json:"q_cast"`
	ECast int `json:"e_cast"`
	XCast int `json:"x_cast"`
}

func (ac *AbilityCasts) UnmarshalJSON(data []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(data, &m)
	if err != nil {
		return err
	}
	ac.CCast, err = safeStringInt(m["c_cast"])
	if err != nil {
		return err
	}
	ac.QCast, err = safeStringInt(m["q_cast"])
	if err != nil {
		return err
	}
	ac.ECast, err = safeStringInt(m["e_cast"])
	if err != nil {
		return err
	}
	ac.XCast, err = safeStringInt(m["x_cast"])
	if err != nil {
		return err
	}
	return nil
}

type Team struct {
	HasWon     bool `json:"has_won"`
	RoundsWon  int  `json:"rounds_won"`
	RoundsLost int  `json:"rounds_lost"`
}

type Round struct {
	WinningTeam string `json:"winning_team"`
	EndType     string `json:"end_type"`
	BombPlanted bool   `json:"bomb_planted"`
	BombDefused bool   `json:"bomb_defused"`
	PlantEvents struct {
		PlantLocation Location `json:"plant_location"`
		PlantedBy     struct {
			DisplayName string `json:"display_name"`
			Team        string `json:"team"`
		} `json:"planted_by"`
		PlantSide              string           `json:"plant_side"`
		PlantTimeInRound       int              `json:"plant_time_in_round"`
		PlayerLocationsOnPlant []PlayerLocation `json:"player_locations_on_plant"`
	} `json:"plant_events"`
	DefuseEvents struct {
		DefusedBy               SimplePlayer     `json:"defused_by"`
		DefuseLocation          Location         `json:"defuse_location"`
		DefuseTimeInRound       int              `json:"defuse_time_in_round"`
		PlayerLocationsOnDefuse []PlayerLocation `json:"player_locations_on_defuse"`
	} `json:"defuse_events"`
	PlayerStats []PlayerStatus `json:"player_stats"`
}
type PlayerStatus struct {
	AbilityCasts struct {
		CCast int `json:"c_cast"`
		QCast int `json:"q_cast"`
		ECast int `json:"e_cast"`
		YCast int `json:"y_cast"`
	} `json:"ability_casts"`
	PlayerPuuid       string        `json:"player_puuid"`
	PlayerDisplayName string        `json:"player_display_name"`
	PlayerTeam        string        `json:"player_team"`
	DamageEvents      []DamageEvent `json:"damage_events"`
	Damage            int           `json:"damage"`
	Bodyshots         int           `json:"bodyshots"`
	Headshots         int           `json:"headshots"`
	Legshots          int           `json:"legshots"`
	KillEvents        []KillEvent   `json:"kill_events"`
	Kills             int           `json:"kills"`
}

type DamageEvent struct {
	ReceiverPuuid       string `json:"receiver_puuid"`
	ReceiverDisplayName string `json:"receiver_display_name"`
	ReceiverTeam        string `json:"receiver_team"`
	Bodyshots           int    `json:"bodyshots"`
	Damage              int    `json:"damage"`
	Headshots           int    `json:"headshots"`
	Legshots            int    `json:"legshots"`
}

type KillEvent struct {
	KillTimeInRound       int              `json:"kill_time_in_round"`
	KillTimeInMatch       int              `json:"kill_time_in_match"`
	KillerPuuid           string           `json:"killer_puuid"`
	KillerDisplayName     string           `json:"killer_display_name"`
	KillerTeam            string           `json:"killer_team"`
	VictimPuuid           string           `json:"victim_puuid"`
	VictimDisplayName     string           `json:"victim_display_name"`
	VictimTeam            string           `json:"victim_team"`
	VictimDeathLocation   Location         `json:"victim_death_location"`
	DamageWeaponID        string           `json:"damage_weapon_id"`
	SecondaryFireMode     bool             `json:"secondary_fire_mode"`
	PlayerLocationsOnKill []PlayerLocation `json:"player_locations_on_kill"`
	Assistants            []Assistant      `json:"assistants"`
}

type Assistant struct {
	AssistantPuuid       string `json:"assistant_puuid"`
	AssistantDisplayName string `json:"assistant_display_name"`
	AssistantTeam        string `json:"assistant_team"`
}

type PlayerLocation struct {
	Location          Location `json:"location"`
	PlayerPuuid       string   `json:"player_puuid"`
	PlayerDisplayName string   `json:"player_display_name"`
	PlayerTeam        string   `json:"player_team"`
}

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

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
