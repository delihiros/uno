package maps

import (
	"fmt"
	_ "image/png"

	"github.com/fogleman/gg"
)

func NewAscent() *Map {
	img, err := gg.LoadPNG("assets/img/ascent_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Ascent",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.00007,
		yMultiplier:  -0.00007,
		xScalarToAdd: 0.813895,
		yScalarToAdd: 0.573242,
		xAdjust:      -0.39,
		yAdjust:      -0.39,
		width:        wh.Max.X,
		height:       wh.Max.Y,
	}
}
