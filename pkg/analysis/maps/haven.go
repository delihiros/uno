package maps

import (
	"fmt"

	"github.com/fogleman/gg"
)

func NewHaven() *Map {
	img, err := gg.LoadPNG("assets/img/haven_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Haven",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.000075,
		yMultiplier:  -0.000075,
		xScalarToAdd: 1.09345,
		yScalarToAdd: 0.642728,
		xAdjust:      -0.734,
		yAdjust:      -0.736,
		width:        wh.Max.X,
		height:       wh.Max.Y,
	}
}
