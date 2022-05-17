package scenes

import (
	"fmt"
	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
	"log"
	"online-game/game"
)

type MenuScene struct {
	game.Scene

	Player *game.Player

	Window *pixelgl.Window

	UI      *pixelui.UI
	UIStack game.UILayerStack
}

func NewMenuScene(
	win *pixelgl.Window,
	ui *pixelui.UI,
	player *game.Player,
) *MenuScene {

	s := &MenuScene{
		UI:     ui,
		Window: win,
		Player: player,
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

	// Ready
	// TODO : Make this dependent to player, and
	// TODO : use the same layer for `host-game-menu`
	guestReady := new(bool)
	*guestReady = false

	playerReady := new(bool)
	*playerReady = false

	// Game
	gameLength := new(int32)
	*gameLength = 3

	gameTime := new(int32)
	*gameTime = 30

	// Waits
	closingServer := new(bool)
	*closingServer = false

	// Layers
	mainMenuLayer := func(ui *pixelui.UI) {

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize | imgui.WindowFlagsNoScrollWithMouse

		if imgui.BeginV("Main Menu", nil, windowFlags) {

			margin := imgui.CursorPos().X

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
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize | imgui.WindowFlagsNoScrollWithMouse

		if imgui.BeginV("Join Game", nil, windowFlags) {

			var margin float32 = 15.0

			buttonSize := imgui.WindowWidth()/float32(2) - margin

			imgui.InputText("Host IP", joinIpString)

			imgui.InputTextV("Password", joinPassString, imgui.InputTextFlagsPassword, nil)

			if imgui.ButtonV("Join", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.Player.Type = game.PlayerTypeClient

				go menu.Player.Client.StartGameLoop()

				go func() {
					if err := menu.Player.Client.ConnectToServer(*joinIpString, *joinPassString); err != nil {
						log.Fatal("[SERVER] Server error:", err)
					}
				}()
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
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize | imgui.WindowFlagsNoScrollWithMouse

		if imgui.BeginV("Host Game", nil, windowFlags) {

			var margin float32 = 15.0

			buttonSize := imgui.WindowWidth()/float32(2) - margin

			imgui.InputText("Host IP", joinIpString)

			imgui.InputTextV("Password", joinPassString, imgui.InputTextFlagsPassword, nil)

			if imgui.ButtonV("Host", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.UIStack.SetSetting("host-game-menu")

				menu.Player.Type = game.PlayerTypeServer

				go menu.Player.Server.StartGameLoop()

				go func() {
					if err := menu.Player.Server.StartServer(*joinIpString, *joinPassString); err != nil {
						log.Fatal("[SERVER] Server error:", err)
					}
				}()
			}

			imgui.SameLineV(0, margin)

			if imgui.ButtonV("Cancel", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.UIStack.SetSetting("main-menu")
			}

			imgui.End()
		}
	}
	menu.UIStack.PushLayer(hostMenuLayer)

	hostGameMenuLayer := func(ui *pixelui.UI) {

		windowFlags := imgui.WindowFlagsNoCollapse | imgui.WindowFlagsNoMove |
			imgui.WindowFlagsNoScrollbar | imgui.WindowFlagsNoResize | imgui.WindowFlagsNoScrollWithMouse

		if imgui.BeginV("Host Game##Next", nil, windowFlags) {

			var pureMargin float32 = 10
			margin := imgui.CursorPos().X

			imgui.BeginTabBar("GameBar")

			if imgui.BeginTabItem("Players") {

				imgui.SetCursorPos(imgui.CursorPos().Plus(imgui.Vec2{Y: pureMargin}))

				nameSize := imgui.WindowWidth() - (pureMargin*2 + 15 +
					imgui.CalcTextSize("Ready", false, 0).X + pureMargin)

				imgui.ButtonV("Guest", imgui.Vec2{X: nameSize, Y: 0})

				imgui.SameLineV(0, pureMargin)

				if imgui.Checkbox("Ready##Guest", guestReady) {
					*guestReady = false
				}

				imgui.ButtonV("You", imgui.Vec2{X: nameSize, Y: 0})

				imgui.SameLineV(0, pureMargin)

				if imgui.Checkbox("Ready##Player", playerReady) {

				}

				imgui.SetCursorPos(imgui.CursorPos().Plus(imgui.Vec2{Y: pureMargin}))

				imgui.EndTabItem()
			}

			if imgui.BeginTabItem("Settings") {

				imgui.SetCursorPos(imgui.CursorPos().Plus(imgui.Vec2{Y: pureMargin}))

				imgui.Text("Win Condition")

				imgui.PushItemWidth(-1)
				imgui.InputIntV("##GameLength", gameLength, 1, 100, 0)

				imgui.Text("Round Time (seconds)")

				imgui.PushItemWidth(-1)
				if imgui.InputIntV("##GameTime", gameTime, 30, 100, 0) {
					if *gameTime < 30 {
						*gameTime = 30
					} else if *gameTime > 300 {
						*gameTime = 300
					}
				}

				imgui.SetCursorPos(imgui.CursorPos().Plus(imgui.Vec2{Y: pureMargin}))

				imgui.EndTabItem()
			}

			imgui.EndTabBar()

			imgui.Separator()

			buttonSize := (imgui.WindowWidth() - margin*2 - pureMargin) / 2
			buttonPos := imgui.Vec2{X: margin, Y: imgui.CursorPos().Y + pureMargin}

			imgui.SetCursorPos(buttonPos)

			if !*playerReady || !*guestReady {
				imgui.PushItemFlag(imgui.ItemFlagsDisabled, true)
				imgui.PushStyleColor(
					imgui.StyleColorButton,
					imgui.Vec4{X: 0.30, Y: 0.30, Z: 0.30, W: 1.0},
				)
			}

			if imgui.ButtonV("Start", imgui.Vec2{X: buttonSize, Y: 0}) {
				fmt.Println("PRESSED!")
			}

			if !*playerReady || !*guestReady {
				imgui.PopStyleColor()
				imgui.PopItemFlag()
			}

			imgui.SameLineV(0, pureMargin)

			if *closingServer {
				imgui.PushItemFlag(imgui.ItemFlagsDisabled, true)
				imgui.PushStyleColor(
					imgui.StyleColorButton,
					imgui.Vec4{X: 0.30, Y: 0.30, Z: 0.30, W: 1.0},
				)
			}

			if imgui.ButtonV("Abort", imgui.Vec2{X: buttonSize, Y: 0}) {
				go func() {
					menu.Player.Type = game.PlayerTypeUndefined

					*closingServer = true
					if err := menu.Player.Server.CloseServer(); err != nil {
						log.Println("[SERVER] Error during server termination:\n", err)
					}
					menu.UIStack.SetSetting("main-menu")
					*closingServer = false
				}()
			}

			if *closingServer {
				imgui.PopStyleColor()
				imgui.PopItemFlag()
			}

			imgui.End()
		}
	}
	menu.UIStack.PushLayer(hostGameMenuLayer)

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

	hostGameMenuSetting := game.UISetting{
		Render: []int{3, 3},
	}
	menu.UIStack.AddSetting("host-game-menu", hostGameMenuSetting)

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
