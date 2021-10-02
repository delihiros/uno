package databank

import (
	"fmt"

	"github.com/delihiros/uno/pkg/analysis/maps"
	"github.com/delihiros/uno/pkg/entities"
	"github.com/delihiros/uno/pkg/view"
)

// TODO: grid size
func (db *Databank) generateGrids() (map[string]*maps.HeatMap, error) {
	matches, err := db.p.ListSavedMatch()
	if err != nil {
		panic(err)
	}

	mapLists := map[string][]*entities.Match{}
	heatmaps := map[string]*maps.HeatMap{}
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

					// save weights and edges to hm
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

		heatmaps[mapName] = hm

	}
	return heatmaps, nil
}

func (db *Databank) Grid(mapName string, x int, y int) (*entities.Grid, error) {
	hm, ok := db.heatmaps[mapName]
	if !ok {
		return nil, fmt.Errorf("map %v not found", mapName)
	}
	gx, gy := hm.BelongsTo(float64(x), float64(y))
	return hm.Get(gx, gy), nil
}
