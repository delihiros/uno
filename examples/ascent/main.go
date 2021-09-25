package main

import (
	"log"

	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/view"
	"github.com/thoas/go-funk"
)

const (
	mapName = "Ascent"
)

func main() {
	m, err := entities.NewMap(mapName)
	if err != nil {
		panic(err)
	}
	mv, err := view.NewMapVisualizer(m)
	if err != nil {
		panic(err)
	}

	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	matches, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	match := funk.Filter(matches, func(match *entities.Match) bool {
		return match.Metadata.Map == mapName
	}).([]*entities.Match)[0]
	ke0 := match.Rounds[12]
	killEvents, err := ke0.KillEvents()
	if err != nil {
		panic(err)
	}
	firstBlood := killEvents[0]
	victimLocation := firstBlood.VictimDeathLocation
	killerLocation := firstBlood.FindKillerLocation()
	if killerLocation != nil {
		vx, vy := m.Scale(mv.Width, mv.Height, victimLocation.X, victimLocation.Y)
		kx, ky := m.Scale(mv.Width, mv.Height, killerLocation.X, killerLocation.Y)
		if firstBlood.KillerTeam == "Red" {
			mv.DrawCircle(kx, ky, 3, 1, 0, 0)
		}
		if firstBlood.KillerTeam == "Blue" {
			mv.DrawCircle(kx, ky, 3, 0, 0, 1)
		}
		mv.DrawCircle(vx, vy, 3, 1, 1, 1)
		mv.DrawLine(vx, vy, kx, ky, 2, 0, 0.5, 0.5)
	}
	for _, pLoc := range firstBlood.PlayerLocationsOnKill {
		x, y := m.Scale(mv.Width, mv.Height, pLoc.Location.X, pLoc.Location.Y)
		if pLoc.PlayerTeam == "Red" {
			mv.DrawCircle(x, y, 3, 1, 0, 0)
		}
		if pLoc.PlayerTeam == "Blue" {
			mv.DrawCircle(x, y, 3, 0, 0, 1)
		}
	}
	mv.SaveImage("death.png")
}

func AccumulateMatches() {
	m, err := entities.NewMap(mapName)
	if err != nil {
		panic(err)
	}
	mv, err := view.NewMapVisualizer(m)
	if err != nil {
		panic(err)
	}

	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	matches, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	matches = funk.Filter(matches, func(match *entities.Match) bool {
		return match.Metadata.Map == mapName
	}).([]*entities.Match)

	log.Println(len(matches))

	for _, match := range matches {
		for _, round := range match.Rounds {
			for _, status := range round.PlayerStats {
				for _, event := range status.KillEvents {
					victimLocation := event.VictimDeathLocation
					killerLocation := event.FindKillerLocation()
					if killerLocation != nil {
						vx, vy := m.Scale(mv.Width, mv.Height, victimLocation.X, victimLocation.Y)
						if event.KillerTeam == "Red" {
							mv.DrawCircle(vx, vy, 3, 1, 0, 0)
						}
						if event.KillerTeam == "Blue" {
							mv.DrawCircle(vx, vy, 3, 0, 0, 1)
						}
						kx, ky := m.Scale(mv.Width, mv.Height, killerLocation.X, killerLocation.Y)
						mv.DrawCircle(kx, ky, 3, 1, 1, 1)
						mv.DrawLine(vx, vy, kx, ky, 2, 0, 0.5, 0.5)
					}
				}
			}
		}
	}

	mv.SaveImage("death.png")
}
