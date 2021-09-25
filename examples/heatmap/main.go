package main

import (
	"github.com/delihiros/uno/pkg/analysis/maps"
	"github.com/delihiros/uno/pkg/database"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/view"
)

func main() {

	db, err := database.Get()
	if err != nil {
		panic(err)
	}
	matches, err := db.ListMatch()
	if err != nil {
		panic(err)
	}
	mapLists := map[string][]*entities.Match{}
	for _, m := range matches {
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
				for _, s := range r.PlayerStats {
					for _, k := range s.KillEvents {
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
					mv.DrawRectangle(float64(j.X*hm.GridSize), float64(j.Y*hm.GridSize), float64(hm.GridSize), float64(hm.GridSize), 1, 0, 0, j.Value/20, true)
				}
			}
		}
		mv.SaveImage(mapName + "-" + "heatmap.png")
	}
}
