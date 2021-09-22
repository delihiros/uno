package entities

import "fmt"

type Map struct {
	Name         string
	ImageURL     string
	xMultiplier  float64
	yMultiplier  float64
	xScalarToAdd float64
	yScalarToAdd float64
	xAdjust      float64
	yAdjust      float64
	width        int
	height       int
}

func NewMap(name string) (*Map, error) {
	switch name {
	case "Ascent":
		return NewAscent(), nil
	case "Bind":
		return NewBind(), nil
	case "Breeze":
		return NewBreeze(), nil
	case "Fracture":
		return NewFracture(), nil
	case "Haven":
		return NewHaven(), nil
	case "Icebox":
		return NewIcebox(), nil
	case "Split":
		return NewSplit(), nil
	default:
		return nil, fmt.Errorf("map not supported: %v", name)
	}
}

func (m *Map) Scale(width, height, rx, ry int) (float64, float64) {
	x := float64(rx)*m.xMultiplier + m.xScalarToAdd + m.xAdjust
	y := 1 - (float64(ry)*m.yMultiplier + m.yScalarToAdd + m.yAdjust)
	return float64(width) * x, float64(height) * y
}

func NewAscent() *Map {
	return &Map{
		Name:         "Ascent",
		ImageURL:     "assets/img/ascent_cw.png",
		xMultiplier:  0.00007,
		yMultiplier:  -0.00007,
		xScalarToAdd: 0.813895,
		yScalarToAdd: 0.573242,
		xAdjust:      -0.39,
		yAdjust:      -0.39,
	}
}

func NewBind() *Map {
	return &Map{
		Name:         "Bind",
		ImageURL:     "assets/img/bind_cw.png",
		xMultiplier:  0.000059,
		yMultiplier:  -0.000059,
		xScalarToAdd: 0.587554,
		yScalarToAdd: 1.032058,
		xAdjust:      -0.553,
		yAdjust:      -0.61,
	}
}

func NewBreeze() *Map {
	return &Map{
		Name:         "Breeze",
		ImageURL:     "assets/img/breeze_cw.png",
		xMultiplier:  0.00007,
		yMultiplier:  -0.00007,
		xScalarToAdd: 0.465123,
		yScalarToAdd: 0.833078,
		xAdjust:      -0.3,
		yAdjust:      -0.3,
	}
}

func NewFracture() *Map {
	return &Map{
		Name:         "Fracture",
		ImageURL:     "assets/img/fracture_cw.png",
		xMultiplier:  0.000078,
		yMultiplier:  -0.000078,
		xScalarToAdd: 0.556952,
		yScalarToAdd: 1.155886,
		// TODO: unknown value
		xAdjust: -0.553,
		yAdjust: -0.61,
	}
}
func NewHaven() *Map {
	return &Map{
		Name:         "Haven",
		ImageURL:     "assets/img/haven_cw.png",
		xMultiplier:  0.000075,
		yMultiplier:  -0.000075,
		xScalarToAdd: 1.09345,
		yScalarToAdd: 0.642728,
		xAdjust:      -0.734,
		yAdjust:      -0.736,
	}
}
func NewIcebox() *Map {
	return &Map{
		Name:         "Icebox",
		ImageURL:     "assets/img/icebox_cw.png",
		xMultiplier:  0.000072,
		yMultiplier:  -0.000072,
		xScalarToAdd: 0.460214,
		yScalarToAdd: 0.304687,
		xAdjust:      0.24,
		yAdjust:      0.235,
	}
}
func NewSplit() *Map {
	return &Map{
		Name:         "Split",
		ImageURL:     "assets/img/split_cw.png",
		xMultiplier:  0.000078,
		yMultiplier:  -0.000078,
		xScalarToAdd: 0.842108,
		yScalarToAdd: 0.648073,
		xAdjust:      -0.54,
		yAdjust:      -0.49,
	}
}
