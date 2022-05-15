package scenes

import (
	"online-game/game"
	"online-game/game/ui"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type GameScene struct {
	game.Scene

	Window  *pixelgl.Window
	UI      *ui.UI
	UIStack ui.UILayerStack
}

func NewGameScene(win *pixelgl.Window) *GameScene {
	s := &GameScene{
		UI:     ui.NewUI(win),
		Window: win,
	}

	// Set integrated scene
	s.Scene = game.NewScene("menu-scene")

	// Set-up UI Stack
	s.UIStack = ui.NewUILayerStack(s.UI)

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

	// Layers
	tempLayer := func(ui *ui.UI) {

		margin := pixel.V(20, 20)

		elementSize := pixel.V(110, 30)

		elementPos := winSize.Sub(elementSize).Scaled(0.5)

		_ = elementPos.Add(margin)
	}
	menu.UIStack.PushLayer(tempLayer)

	// Settings
	tempSetting := ui.UISetting{
		Render: []int{0, 0},
		Input:  []int{0, 0},
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
