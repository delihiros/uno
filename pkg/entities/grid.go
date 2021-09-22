package entities

type Grid struct {
	X     int
	Y     int
	Value float64
	To    []*Edge
}

func NewGrid(x, y int) *Grid {
	return &Grid{
		X:     x,
		Y:     y,
		Value: 0,
		To:    []*Edge{},
	}
}

type Edge struct {
	Weight float64
	To     *Grid
}

func (g *Grid) SetEdge(e *Edge) {
	for _, t := range g.To {
		if e.To == t.To {
			t.Weight = e.Weight
			return
		}
	}
	g.To = append(g.To, e)
}
