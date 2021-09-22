package view

import (
	"fmt"

	"github.com/delihiros/uno/pkg/entities"
	"github.com/fogleman/gg"
)

type MapVisualizer struct {
	m      *entities.Map
	dc     *gg.Context
	width  int
	height int
}

func NewMapVisualizer(m *entities.Map) *MapVisualizer {
	img, err := gg.LoadPNG("assets/img/ascent_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &MapVisualizer{
		m:      m,
		dc:     dc,
		width:  wh.Max.X,
		height: wh.Max.Y,
	}
}

func (mv *MapVisualizer) Scale(rx, ry float64) (float64, float64) {
	return mv.m.Scale(mv.width, mv.height, rx, ry)
}

func (mv *MapVisualizer) SaveImage(path string) error {
	return mv.dc.SavePNG(path)
}

func (mv *MapVisualizer) DrawCircle(rx, ry, r, red, green, blue float64) {
	x, y := mv.Scale(rx, ry)
	mv.dc.SetRGB(red, green, blue)
	mv.dc.DrawCircle(x, y, r)
	mv.dc.Stroke()
}

func (mv *MapVisualizer) DrawLine(rx1, ry1, rx2, ry2, lineWidth, red, green, blue float64) {
	x1, y1 := mv.Scale(rx1, ry1)
	x2, y2 := mv.Scale(rx2, ry2)
	mv.dc.SetLineWidth(lineWidth)
	mv.dc.SetRGB(red, green, blue)
	mv.dc.DrawLine(x1, y1, x2, y2)
	mv.dc.Stroke()
}
