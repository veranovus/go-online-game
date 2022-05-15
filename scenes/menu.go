package scenes

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"online-game/game"
	"online-game/game/ui"
)

type MenuScene struct {
	game.Scene

	Window  *pixelgl.Window
	UI      *ui.UI
	UIStack ui.UILayerStack
}

func NewMenuScene(win *pixelgl.Window) *MenuScene {
	s := &MenuScene{
		UI:     ui.NewUI(win),
		Window: win,
	}

	// Set integrated scene
	s.Scene = game.NewScene("menu-scene")

	s.UIStack = ui.NewUILayerStack(s.UI)

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

	// Window size
	winSize := menu.Window.Bounds().Size()

	// Layers
	mainMenuLayer := func(ui *ui.UI) {

		margin := pixel.V(20, 20)

		buttonSize := pixel.V(110, 30)

		buttonPos := winSize.Sub(buttonSize).Scaled(0.5)

		if ui.Button(buttonPos, buttonSize, "Play") {
			menu.ReturnState = game.SceneStateChange + SceneIndexGame
		}

		buttonPos.Y -= buttonSize.Y + margin.Y

		if ui.Button(buttonPos, buttonSize, "Quit") {
			menu.Window.SetClosed(true)
		}
	}
	menu.UIStack.PushLayer(mainMenuLayer)

	// Settings
	mainMenuSetting := ui.UISetting{
		Render: []int{0, 0},
		Input:  []int{0, 0},
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
