package view

import (
	"fmt"

	"github.com/delihiros/uno/pkg/entities"
	"github.com/fogleman/gg"
)

type MapVisualizer struct {
	m      *entities.Map
	dc     *gg.Context
	Width  int
	Height int
}

func NewMapVisualizer(m *entities.Map) (*MapVisualizer, error) {
	img, err := gg.LoadPNG("assets/img/ascent_cw.png")
	if err != nil {
		return nil, fmt.Errorf("could not find image")
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &MapVisualizer{
		m:      m,
		dc:     dc,
		Width:  wh.Max.X,
		Height: wh.Max.Y,
	}, nil
}

func (mv *MapVisualizer) Scale(rx, ry int) (float64, float64) {
	return mv.m.Scale(mv.Width, mv.Height, rx, ry)
}

func (mv *MapVisualizer) SaveImage(path string) error {
	return mv.dc.SavePNG(path)
}

func (mv *MapVisualizer) DrawCircle(x, y float64, r, red, green, blue float64) {
	mv.dc.SetRGB(red, green, blue)
	mv.dc.DrawCircle(x, y, r)
	mv.dc.Stroke()
}

func (mv *MapVisualizer) DrawLine(x1, y1, x2, y2, lineWidth, red, green, blue float64) {
	mv.dc.SetLineWidth(lineWidth)
	mv.dc.SetRGB(red, green, blue)
	mv.dc.DrawLine(x1, y1, x2, y2)
	mv.dc.Stroke()
}

func (mv *MapVisualizer) DrawRectangle(x1, y1, w, h, red, green, blue, alpha float64, fill bool) {
	mv.dc.SetRGBA(red, green, blue, alpha)
	mv.dc.DrawRectangle(x1, y1, w, h)
	mv.dc.Fill()
	mv.dc.Stroke()
}
