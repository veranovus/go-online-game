package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"image/color"
)

func DrawRect(imd *imdraw.IMDraw, x, y, w, h float64, color color.Color) {

	imd.Color = color
	imd.Push(pixel.V(x, y))
	imd.Push(pixel.V(x+w, y))
	imd.Push(pixel.V(x+w, y+h))
	imd.Push(pixel.V(x, y+h))
	imd.Polygon(0)
}

func DrawRectLines(imd *imdraw.IMDraw, thickness, x, y, w, h float64, color color.Color) {

	imd.Color = color
	imd.EndShape = imdraw.NoEndShape
	imd.Push(pixel.V(x, y), pixel.V(x+w, y))
	imd.Push(pixel.V(x+w, y+h), pixel.V(x, y+h))
	imd.Push(pixel.V(x, y))
	imd.Line(thickness)
}
