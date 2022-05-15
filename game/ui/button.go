package ui

import (
	"fmt"
	"github.com/faiface/pixel"
	"golang.org/x/image/colornames"
	"online-game/game"
)

func (ui *UI) Button(pos pixel.Vec, size pixel.Vec, label string) bool {

	mousePos := ui.win.MousePosition()

	cursorOver := false
	pressedOver := false

	// Detect
	if !ui.lock {
		if (mousePos.X >= pos.X && mousePos.X <= pos.X+size.X) &&
			(mousePos.Y >= pos.Y && mousePos.Y <= pos.Y+size.Y) {
			cursorOver = true
		}

		if cursorOver && ui.win.Pressed(0) {
			pressedOver = true
		}
	}

	// Render
	color := pixel.RGB(0.24, 0.04, 0.48)

	if cursorOver {
		color = pixel.RGB(0.27, 0.07, 0.51)
	} else if pressedOver {
		color = pixel.RGB(0.30, 0.1, 0.54)
	}

	// Clear the ui and text
	ui.text.Clear()
	ui.Clear()

	// Draw the fill
	game.DrawRect(ui.imd, pos.X, pos.Y, size.X, size.Y, color)

	// Draw the outline
	game.DrawRectLines(ui.imd, 1, pos.X, pos.Y, size.X, size.Y, colornames.White)

	// Set the text
	fmt.Fprint(ui.text, label)

	// Text pos
	textPos := size.Sub(ui.text.Bounds().Size()).Scaled(0.5)

	// Draw the button
	ui.Draw()

	// Draw the text
	ui.text.Draw(ui.win, pixel.IM.Moved(pos.Add(textPos)))

	// Event
	if cursorOver && ui.win.JustReleased(0) {
		return true
	}

	return false
}
