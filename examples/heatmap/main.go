package main

import (
	"log"

	"github.com/delihiros/uno/pkg/analysis/maps"
	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/view"
)

const (
	databaseURL = "http://localhost"
	port        = 8080
)

func main() {
	generateMostKilledFrom()
}

func generateDefenderDeathHeatmap() {
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
		hm := maps.NewHeatMap(4, mv.Width, mv.Height)
		for _, match := range ms {
			for i, r := range match.Rounds {
				attacker := match.AttackerOf(i)
				events, err := r.KillEvents()
				if err != nil {
					panic(err)
				}
				for _, k := range events {
					if k.KillerTeam == attacker {
						x, y := m.Scale(mv.Width, mv.Height, k.VictimDeathLocation.X, k.VictimDeathLocation.Y)
						gx, gy := hm.BelongsTo(x, y)
						g := hm.Get(gx, gy)
						g.Value++
					}
				}
			}
		}

		for _, i := range hm.Grids {
			for _, j := range i {
				if j.Value > 0 {
					mv.DrawRectangle(float64(j.X*hm.GridSize), float64(j.Y*hm.GridSize), float64(hm.GridSize), float64(hm.GridSize), 1, 0, 0, j.Value/10, true)
				}
			}
		}
		mv.SaveImage(mapName + "-" + "defenderDeathHeatmap.png")
	}
}

func generateAttackerDeathHeatmap() {
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
		hm := maps.NewHeatMap(4, mv.Width, mv.Height)
		for _, match := range ms {
			for i, r := range match.Rounds {
				attacker := match.AttackerOf(i)
				events, err := r.KillEvents()
				if err != nil {
					panic(err)
				}
				for _, k := range events {
					if k.VictimTeam == attacker {
						x, y := m.Scale(mv.Width, mv.Height, k.VictimDeathLocation.X, k.VictimDeathLocation.Y)
						gx, gy := hm.BelongsTo(x, y)
						g := hm.Get(gx, gy)
						g.Value++
					}
				}
			}
		}

		for _, i := range hm.Grids {
			for _, j := range i {
				if j.Value > 0 {
					mv.DrawRectangle(float64(j.X*hm.GridSize), float64(j.Y*hm.GridSize), float64(hm.GridSize), float64(hm.GridSize), 1, 0, 0, j.Value/10, true)
				}
			}
		}
		mv.SaveImage(mapName + "-" + "attackerDeathHeatmap.png")
	}
}

func generateMostKilledFrom() {
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
		hm := maps.NewHeatMap(4, mv.Width, mv.Height)
		for _, match := range ms {
			for _, r := range match.Rounds {
				events, err := r.KillEvents()
				if err != nil {
					panic(err)
				}
				for _, event := range events {
					kl, err := event.FindKillerLocation()
					if err != nil {
						continue
					}
					// calculate x, y of killer/victim's location
					vx, vy := m.Scale(mv.Width, mv.Height, event.VictimDeathLocation.X, event.VictimDeathLocation.Y)
					kx, ky := m.Scale(mv.Width, mv.Height, kl.X, kl.Y)
					gvx, gvy := hm.BelongsTo(vx, vy)
					gkx, gky := hm.BelongsTo(kx, ky)

					vgrid := hm.Get(gvx, gvy)
					kgrid := hm.Get(gkx, gky)
					toKiller := vgrid.GetEdge(gkx, gky)
					if toKiller == nil {
						toKiller = &entities.Edge{
							To: kgrid,
						}
					}
					toKiller.Weight++
					vgrid.SetEdge(toKiller)
					vgrid.Value++
				}
			}
		}

		maxWeightGrid := hm.Grids[0][0]
		maxWeight := 0.0
		for _, i := range hm.Grids {
			for _, j := range i {
				weight := 0.0
				for _, e := range j.To {
					weight += e.Weight
				}
				if maxWeight < weight {
					maxWeight = weight
					maxWeightGrid = j
				}
			}
		}

		log.Println(maxWeightGrid.X, maxWeightGrid.Y, len(maxWeightGrid.To), maxWeight)

		for _, edge := range maxWeightGrid.To {
			mv.DrawRectangle(float64(edge.To.X*hm.GridSize), float64(edge.To.Y*hm.GridSize), float64(hm.GridSize), float64(hm.GridSize), 0, 0, 1, edge.Weight/5, true)
		}
		mv.DrawRectangle(float64(maxWeightGrid.X*hm.GridSize), float64(maxWeightGrid.Y*hm.GridSize), float64(hm.GridSize), float64(hm.GridSize), 1, 0, 0, 1, true)

		mv.SaveImage(mapName + "-" + "mostKilledFrom.png")
	}
}
