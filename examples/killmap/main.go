package main

import (
	"log"

	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/view"
	"github.com/thoas/go-funk"
)

const (
	databaseURL = "http://localhost"
	port        = 8080
)

func main() {

}

func FirstBlood() {
	const mapName = "Ascent"
	m, err := entities.NewMap(mapName)
	if err != nil {
		panic(err)
	}
	mv, err := view.NewMapVisualizer(m)
	if err != nil {
		panic(err)
	}

	db, err := database.Get(databaseURL, port)
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
	killerLocation, err := firstBlood.FindKillerLocation()
	if err == nil {
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

func AccumulateMatches(mapName string) {
	m, err := entities.NewMap(mapName)
	if err != nil {
		panic(err)
	}
	mv, err := view.NewMapVisualizer(m)
	if err != nil {
		panic(err)
	}

	db, err := database.Get(databaseURL, port)
	if err != nil {
		panic(err)
	}
	matchIDs, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	matches := []*entities.Match{}
	for _, id := range matchIDs {
		m, err := db.Match(id)
		if err != nil {
			panic(err)
		}
		matches = append(matches, m)
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
					killerLocation, err := event.FindKillerLocation()
					if err == nil {
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

func generateKillMap() {
	db, err := database.Get(databaseURL, port)
	if err != nil {
		panic(err)
	}
	matchIDs, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	mapLists := map[string][]*entities.Match{}
	for _, id := range matchIDs {
		m, err := db.Match(id)
		if err != nil {
			panic(err)
		}
		_, ok := mapLists[m.Metadata.Map]
		if !ok {
			mapLists[m.Metadata.Map] = []*entities.Match{}
		}
		mapLists[m.Metadata.Map] = append(mapLists[m.Metadata.Map], m)
	}

	for mapName, ms := range mapLists {
		if mapName == "Fracture" {
			continue
		}
		m, err := entities.NewMap(mapName)
		if err != nil {
			panic(err)
		}
		mv, err := view.NewMapVisualizer(m)
		if err != nil {
			panic(err)
		}
		for _, match := range ms {
			for i, r := range match.Rounds {
				attacker := match.AttackerOf(i)
				events, err := r.KillEvents()
				if err != nil {
					panic(err)
				}
				for _, event := range events {
					if event.KillerTeam == attacker {
						kl, err := event.FindKillerLocation()
						if err == nil {
							kx, ky := mv.Scale(kl.X, kl.Y)
							mv.DrawCircle(kx, ky, 3, 1, 0, 0)
							vx, vy := mv.Scale(event.VictimDeathLocation.X, event.VictimDeathLocation.Y)
							mv.DrawCircle(vx, vy, 3, 0, 0, 1)
							mv.DrawLine(kx, ky, vx, vy, 1, 1, 0, 1)
						}
					}
				}
			}
		}

		mv.SaveImage(mapName + "-" + "heatmap.png")
	}
}
