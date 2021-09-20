package maps

import (
	"fmt"

	"github.com/fogleman/gg"
)

func NewFracture() *Map {
	img, err := gg.LoadPNG("assets/img/fracture_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Fracture",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.000078,
		yMultiplier:  -0.000078,
		xScalarToAdd: 0.556952,
		yScalarToAdd: 1.155886,
		// TODO: unknown value
		xAdjust: -0.553,
		yAdjust: -0.61,
		width:   wh.Max.X,
		height:  wh.Max.Y,
	}
}
