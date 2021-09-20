package maps

import (
	"fmt"

	"github.com/fogleman/gg"
)

func NewBind() *Map {
	img, err := gg.LoadPNG("assets/img/bind_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Bind",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.000059,
		yMultiplier:  -0.000059,
		xScalarToAdd: 0.587554,
		yScalarToAdd: 1.032058,
		xAdjust:      -0.553,
		yAdjust:      -0.61,
		width:        wh.Max.X,
		height:       wh.Max.Y,
	}
}
