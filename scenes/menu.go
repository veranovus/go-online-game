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

	// Layers
	mainMenuLayer := func(ui *pixelui.UI) {

		imgui.ShowDemoWindow(nil)

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize

		if imgui.BeginV("Main Menu", nil, windowFlags) {

			if imgui.Button("Join") {

			}

			if imgui.Button("Host") {

			}

			if imgui.Button("Quit") {
				menu.Window.SetClosed(true)
			}

			imgui.End()
		}

	}
	menu.UIStack.PushLayer(mainMenuLayer)

	// Settings
	mainMenuSetting := game.UISetting{
		Render: []int{0, 0},
	}
	menu.UIStack.AddSetting("main-menu", mainMenuSetting)

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
