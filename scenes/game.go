package scenes

import (
	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"online-game/client"
	"online-game/game"
	"online-game/server"
)

type GameScene struct {
	game.Scene

	Client *client.Client
	Server *server.Server

	Window *pixelgl.Window

	UI      *pixelui.UI
	UIStack game.UILayerStack
}

func NewGameScene(
	win *pixelgl.Window,
	ui *pixelui.UI,
	server *server.Server,
	client *client.Client,
) *GameScene {

	s := &GameScene{
		UI:     ui,
		Window: win,
		Client: client,
		Server: server,
	}

	// Set integrated scene
	s.Scene = game.NewScene("menu-scene")

	// Set-up UI Stack
	s.UIStack = game.NewUILayerStack(s.UI)

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
	tempLayer := func(ui *pixelui.UI) {

		margin := pixel.V(20, 20)

		elementSize := pixel.V(110, 30)

		elementPos := winSize.Sub(elementSize).Scaled(0.5)

		_ = elementPos.Add(margin)
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
