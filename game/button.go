package game

import (
	"golang.org/x/image/colornames"
)

func GameButton(g *Graphics, text string, x, y, w, h float64) bool {

	cursorPos := g.Window.MousePosition()

	// Detection
	pressedOn := false
	cursorOn := false
	trigger := false

	if (cursorPos.X > x && cursorPos.X < x+w) && (cursorPos.Y > y && cursorPos.Y < y+h) {
		cursorOn = true
	}

	if cursorOn && g.Window.JustPressed(0) {
		pressedOn = true
	} else if cursorOn && g.Window.JustReleased(0) {
		trigger = true
	}

	// Draw
	g.IMD.Clear()

	if pressedOn {
		g.DrawRectLines(2, x+1, y+1, w-2, h-2, colornames.White)
		g.DrawText(text, x, y, colornames.White)
	} else if cursorOn {
		g.DrawRectLines(2, x, y, w, h, colornames.Whitesmoke)
		g.DrawText(text, x, y, colornames.White)
	} else {
		g.DrawRectLines(2, x, y, w, h, colornames.White)
		g.DrawText(text, x, y, colornames.White)
	}

	g.IMD.Draw(g.Window)

	// Event
	if trigger {
		return true
	}
	return false
}
