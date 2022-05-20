package scenes

import (
	"fmt"
	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
	"golang.org/x/image/colornames"
	"online-game/game"
)

type GameScene struct {
	game.Scene

	Game *game.Game

	Window   *pixelgl.Window
	Graphics *game.Graphics

	UI      *pixelui.UI
	UIStack game.UILayerStack
}

func NewGameScene(
	win *pixelgl.Window,
	ui *pixelui.UI,
	g *game.Game,
) *GameScene {

	s := &GameScene{
		UI:     ui,
		Window: win,
		Game:   g,
	}

	// Set integrated scene
	s.Scene = game.NewScene("menu-scene")

	// Set-up UI Stack
	s.UIStack = game.NewUILayerStack(s.UI)

	// Graphics
	s.Graphics = game.NewGraphics(s.Window)

	return s
}

func (menu *GameScene) GetName() string {
	return menu.Name
}

func (menu *GameScene) GetID() int {
	return menu.ID
}

func (menu *GameScene) SetID(id int) {
	menu.ID = id
}

func (menu *GameScene) Load() bool {

	// Window size
	winSize := menu.Window.Bounds().Size()
	_ = winSize

	// Layers
	tempLayer := func(ui *pixelui.UI) {

		changed := false

		buttonSize := imgui.Vec2{X: 80, Y: 50}
		//buttonMargin := (winSize.X - buttonSize.X*3) / 4
		//buttonPos := pixel.V(buttonMargin, 40)

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize |
			imgui.WindowFlagsNoScrollWithMouse | imgui.WindowFlagsNoBringToFrontOnFocus

		if imgui.BeginV("##Select", nil, windowFlags) {

			menu.Graphics.DrawText(
				fmt.Sprintf("YOU: %02d", menu.Game.Player.Score),
				5, winSize.Y-17, colornames.White,
			)

			menu.Graphics.DrawText(
				fmt.Sprintf("HOST: %02d", menu.Game.Player.OtherScore),
				winSize.X-60, winSize.Y-17, colornames.White,
			)

			if menu.Game.Player.Card != game.CardTypeNone {
				imgui.PushItemFlag(imgui.ItemFlagsDisabled, true)
				imgui.PushStyleColor(
					imgui.StyleColorButton,
					imgui.Vec4{X: 0.30, Y: 0.30, Z: 0.30, W: 1.0},
				)
			}

			imguiMargin := imgui.CursorPos().X
			margin := (imgui.WindowWidth() - (imguiMargin*2 + buttonSize.X*3)) / 2

			if imgui.ButtonV("Rock", buttonSize) {

				menu.Game.Player.Card = game.CardTypeRock
				changed = true
			}

			imgui.SameLineV(0, margin)

			if imgui.ButtonV("Paper", buttonSize) {

				menu.Game.Player.Card = game.CardTypePaper
				changed = true
			}

			imgui.SameLineV(0, margin)

			if imgui.ButtonV("Scissor", buttonSize) {

				menu.Game.Player.Card = game.CardTypeScissor
				changed = true
			}

			if menu.Game.Player.Card != game.CardTypeNone && !changed {
				imgui.PopStyleColor()
				imgui.PopItemFlag()
			}
		}
		imgui.End()

		if menu.Game.Player.Card != game.CardTypeNone && changed {
			menu.Game.Client.SendMessage(
				game.MessageTypePick,
				fmt.Sprint(menu.Game.Player.Card),
			)
		}
	}
	menu.UIStack.PushLayer(tempLayer)

	// Settings
	tempSetting := game.UISetting{
		Render: []int{0, 0},
	}
	menu.UIStack.AddSetting("temp-menu", tempSetting)

	// Set the active settings
	menu.UIStack.SetSetting("temp-menu")

	return true
}

func (menu *GameScene) UnLoad() {

}

func (menu *GameScene) Update(dt float64) game.SceneState {

	// Update and Render UIStack
	menu.UIStack.Update()

	return menu.ReturnState
}
