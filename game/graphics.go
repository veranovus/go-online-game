package game

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	"log"
)

type Graphics struct {
	Text   *text.Text
	Atlas  *text.Atlas
	IMD    *imdraw.IMDraw
	Window *pixelgl.Window
}

func NewGraphics(win *pixelgl.Window) *Graphics {
	g := &Graphics{
		Atlas:  text.NewAtlas(basicfont.Face7x13, text.ASCII),
		IMD:    imdraw.New(nil),
		Window: win,
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

func (g *Graphics) DrawText(text string, x, y float64, color color.Color) {

	g.Text.Clear()

	_, err := fmt.Fprintln(g.Text, text)
	if err != nil {
		log.Fatal(err)
	}

	g.Text.Color = color
	g.Text.Draw(g.Window, pixel.IM.Moved(pixel.V(x, y)))
}
