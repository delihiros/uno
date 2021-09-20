package client

import "fmt"

func (e *KillEvent) FindKillerLocation() *Location {
	killer := e.KillerPuuid
	for _, loc := range e.PlayerLocationsOnKill {
		if loc.PlayerPuuid == killer {
			return &loc.Location
		}
	}
	return nil
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
