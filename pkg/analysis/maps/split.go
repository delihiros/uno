package maps

import (
	"fmt"

	"github.com/fogleman/gg"
)

func NewSplit() *Map {
	img, err := gg.LoadPNG("assets/img/split_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Split",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.000078,
		yMultiplier:  -0.000078,
		xScalarToAdd: 0.842108,
		yScalarToAdd: 0.648073,
		xAdjust:      -0.54,
		yAdjust:      -0.49,
		width:        wh.Max.X,
		height:       wh.Max.Y,
	}
}
