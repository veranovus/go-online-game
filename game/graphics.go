package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

type Graphics struct {
	Text  *text.Text
	Atlas *text.Atlas
	IMD   *imdraw.IMDraw
}

func NewGraphics() *Graphics {
	g := &Graphics{
		Atlas: text.NewAtlas(basicfont.Face7x13, text.ASCII),
		IMD:   imdraw.New(nil),
	}

	g.Text = text.New(pixel.V(0, 0), g.Atlas)

	return g
}

func (g *Graphics) DrawRect(x, y, w, h float64, color color.Color) {

	g.IMD.Color = color
	g.IMD.Push(pixel.V(x, y))
	g.IMD.Push(pixel.V(x+w, y))
	g.IMD.Push(pixel.V(x+w, y+h))
	g.IMD.Push(pixel.V(x, y+h))
	g.IMD.Polygon(0)
}

func (g *Graphics) DrawRectLines(thickness, x, y, w, h float64, color color.Color) {

	g.IMD.Color = color
	g.IMD.EndShape = imdraw.NoEndShape
	g.IMD.Push(pixel.V(x, y), pixel.V(x+w, y))
	g.IMD.Push(pixel.V(x+w, y+h), pixel.V(x, y+h))
	g.IMD.Push(pixel.V(x, y))
	g.IMD.Line(thickness)
}

func (g *Graphics) DrawText(x, y float64, color color.Color) {

}
