package maps

import (
	"fmt"

	"github.com/fogleman/gg"
)

func NewBreeze() *Map {
	img, err := gg.LoadPNG("assets/img/breeze_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Breeze",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.00007,
		yMultiplier:  -0.00007,
		xScalarToAdd: 0.465123,
		yScalarToAdd: 0.833078,
		xAdjust:      -0.3,
		yAdjust:      -0.3,
		width:        wh.Max.X,
		height:       wh.Max.Y,
	}
}
