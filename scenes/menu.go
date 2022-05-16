package scenes

import (
	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
	"online-game/game"
)

type MenuScene struct {
	game.Scene

	Window  *pixelgl.Window
	UI      *pixelui.UI
	UIStack game.UILayerStack
}

func NewMenuScene(win *pixelgl.Window, ui *pixelui.UI) *MenuScene {
	s := &MenuScene{
		UI:     ui,
		Window: win,
	}

	// Set integrated scene
	s.Scene = game.NewScene("menu-scene")

	s.UIStack = game.NewUILayerStack(s.UI)

	return s
}

func (menu *MenuScene) GetName() string {
	return menu.Name
}

func (menu *MenuScene) GetID() int {
	return menu.ID
}

func (menu *MenuScene) SetID(id int) {
	menu.ID = id
}

func (menu *MenuScene) Load() bool {

	// String
	joinIpString := new(string)
	joinPassString := new(string)

	// Layers
	mainMenuLayer := func(ui *pixelui.UI) {

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize

		if imgui.BeginV("Main Menu", nil, windowFlags) {

			margin := imgui.CursorPos().X + 20

			buttonSize := imgui.WindowWidth() - margin*2
			buttonPos := imgui.Vec2{X: margin, Y: imgui.CursorPos().Y}

			imgui.SetCursorPos(buttonPos)

			if imgui.ButtonV("Join", imgui.Vec2{X: buttonSize}) {
				menu.UIStack.SetSetting("join-menu")
			}

			buttonPos.Y = imgui.CursorPos().Y
			imgui.SetCursorPos(buttonPos)

			if imgui.ButtonV("Host", imgui.Vec2{X: buttonSize}) {
				menu.UIStack.SetSetting("host-menu")
			}

			buttonPos.Y = imgui.CursorPos().Y
			imgui.SetCursorPos(buttonPos)

			if imgui.ButtonV("Quit", imgui.Vec2{X: buttonSize}) {
				menu.Window.SetClosed(true)
			}

			imgui.End()
		}

	}
	menu.UIStack.PushLayer(mainMenuLayer)

	joinMenuStack := func(ui *pixelui.UI) {

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize

		if imgui.BeginV("Join Game", nil, windowFlags) {

			var margin float32 = 15.0

			buttonSize := imgui.WindowWidth()/float32(2) - margin

			imgui.InputText("Host IP", joinIpString)

			imgui.InputText("Password", joinPassString)

			if imgui.ButtonV("Join", imgui.Vec2{X: buttonSize, Y: 0}) {

			}

			imgui.SameLineV(0, margin)

			if imgui.ButtonV("Cancel", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.UIStack.SetSetting("main-menu")
			}

			imgui.End()
		}
	}
	menu.UIStack.PushLayer(joinMenuStack)

	hostMenuLayer := func(ui *pixelui.UI) {

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize

		if imgui.BeginV("Host Game", nil, windowFlags) {

			var margin float32 = 15.0

			buttonSize := imgui.WindowWidth()/float32(2) - margin

			imgui.InputText("Host IP", joinIpString)

			imgui.InputText("Password", joinPassString)

			if imgui.ButtonV("Host", imgui.Vec2{X: buttonSize, Y: 0}) {

			}

			imgui.SameLineV(0, margin)

			if imgui.ButtonV("Cancel", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.UIStack.SetSetting("main-menu")
			}

			imgui.End()
		}
	}
	menu.UIStack.PushLayer(hostMenuLayer)

	// Settings
	mainMenuSetting := game.UISetting{
		Render: []int{0, 0},
	}
	menu.UIStack.AddSetting("main-menu", mainMenuSetting)

	joinMenuSetting := game.UISetting{
		Render: []int{1, 1},
	}
	menu.UIStack.AddSetting("join-menu", joinMenuSetting)

	hostMenuSetting := game.UISetting{
		Render: []int{2, 2},
	}
	menu.UIStack.AddSetting("host-menu", hostMenuSetting)

	// Set the active settings
	menu.UIStack.SetSetting("main-menu")

	return true
}

func (menu *MenuScene) UnLoad() {

}

func (menu *MenuScene) Update(dt float64) game.SceneState {

	// Update and Render UIStack
	menu.UIStack.Update()

	return menu.ReturnState
}
