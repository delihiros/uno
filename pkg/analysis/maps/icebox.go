package maps

import (
	"fmt"

	"github.com/fogleman/gg"
)

func NewIcebox() *Map {
	img, err := gg.LoadPNG("assets/img/icebox_cw.png")
	if err != nil {
		panic(fmt.Errorf("could not find image"))
	}
	dc := gg.NewContextForImage(img)
	wh := img.Bounds()
	return &Map{
		MapName:      "Icebox",
		mapImage:     &img,
		dc:           dc,
		xMultiplier:  0.000072,
		yMultiplier:  -0.000072,
		xScalarToAdd: 0.460214,
		yScalarToAdd: 0.304687,
		xAdjust:      0.24,
		yAdjust:      0.235,
		width:        wh.Max.X,
		height:       wh.Max.Y,
	}
}
