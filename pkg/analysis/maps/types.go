package maps

import (
	"image"

	"github.com/fogleman/gg"
)

type Map struct {
	MapName      string
	mapImage     *image.Image
	dc           *gg.Context
	xMultiplier  float64
	yMultiplier  float64
	xScalarToAdd float64
	yScalarToAdd float64
	xAdjust      float64
	yAdjust      float64
	width        int
	height       int
}

func (m *Map) scale(rx, ry float64) (float64, float64) {
	x := rx*m.xMultiplier + m.xScalarToAdd + m.xAdjust
	y := 1 - (ry*m.yMultiplier + m.yScalarToAdd + m.yAdjust)
	return float64(m.width) * x, float64(m.height) * y
}

func (m *Map) SaveImage(path string) error {
	return m.dc.SavePNG(path)
}

func (m *Map) DrawCircle(rx, ry, r, red, green, blue float64) {
	x, y := m.scale(rx, ry)
	m.dc.SetRGB(red, green, blue)
	m.dc.DrawCircle(x, y, r)
	m.dc.Stroke()
}

func (m *Map) DrawLine(rx1, ry1, rx2, ry2, lineWidth, red, green, blue float64) {
	x1, y1 := m.scale(rx1, ry1)
	x2, y2 := m.scale(rx2, ry2)
	m.dc.SetLineWidth(lineWidth)
	m.dc.SetRGB(red, green, blue)
	m.dc.DrawLine(x1, y1, x2, y2)
	m.dc.Stroke()
}
