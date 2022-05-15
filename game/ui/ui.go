package ui

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type UI struct {
	win   *pixelgl.Window
	imd   *imdraw.IMDraw
	atlas *text.Atlas
	text  *text.Text
	lock  bool
}

func (ui *UI) Draw() {
	ui.imd.Draw(ui.win)
}

func (ui *UI) Clear() {
	ui.imd.Clear()
}

func (ui *UI) Lock() {
	ui.lock = true
}

func (ui *UI) UnLock() {
	ui.lock = false
}

func NewUI(win *pixelgl.Window) *UI {

	atlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)

	return &UI{
		win:   win,
		imd:   imdraw.New(nil),
		atlas: atlas,
		text:  text.New(pixel.V(0, 0), atlas),
		lock:  false,
	}
}
