package main

import (
	"log"

	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/view"
	"github.com/thoas/go-funk"

	"github.com/delihiros/uno/pkg/entities"
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

	db, err := database.New()
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
					vx, vy := m.Scale(mv.Width, mv.Height, victimLocation.X, victimLocation.Y)
					killerLocation := event.FindKillerLocation()
					mv.DrawCircle(vx, vy, 3, 1, 0, 0)
					if killerLocation != nil {
						kx, ky := m.Scale(mv.Width, mv.Height, killerLocation.X, killerLocation.Y)
						mv.DrawCircle(kx, ky, 3, 0, 0, 1)
						mv.DrawLine(vx, vy, kx, ky, 2, 0, 0.5, 0.5)
					}
				}
			}
		}
	}
	mv.SaveImage("death.png")
}
