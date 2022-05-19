package scenes

import (
	"fmt"
	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel/pixelgl"
	"github.com/inkyblackness/imgui-go"
	"github.com/shrainu/gnet"
	"log"
	"online-game/game"
	"strconv"
)

type MenuScene struct {
	game.Scene

	Game *game.Game

	Window *pixelgl.Window

	UI      *pixelui.UI
	UIStack game.UILayerStack
}

func NewMenuScene(
	win *pixelgl.Window,
	ui *pixelui.UI,
	g *game.Game,
) *MenuScene {

	s := &MenuScene{
		UI:     ui,
		Game:   g,
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

			imgui.InputTextV("Password", &menu.Game.Player.Password, imgui.InputTextFlagsPassword, nil)

			if imgui.ButtonV("Join", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.UIStack.SetSetting("host-game-menu")

				menu.Game.Player.Type = game.PlayerTypeClient

				go menu.Game.Client.ProcessMessages()

				go func() {
					menu.Game.Client.Client.Channel = make(chan gnet.Message)
					if err := menu.Game.Client.Client.ConnectToServer(*joinIpString); err != nil {
						log.Fatal(err)
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

			imgui.InputTextV("Password", &menu.Game.Player.Password, imgui.InputTextFlagsPassword, nil)

			if imgui.ButtonV("Host", imgui.Vec2{X: buttonSize, Y: 0}) {
				menu.UIStack.SetSetting("host-game-menu")

				menu.Game.Player.Type = game.PlayerTypeServer

				go func() {
					menu.Game.Server.Server.Active = true
					if err := menu.Game.Server.Server.StartServer(*joinIpString); err != nil {
						log.Fatal(err)
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

				var otherString string
				if menu.Game.Player.Type == game.PlayerTypeClient {
					otherString = "Host"
				} else {
					otherString = "Guest"
				}

				imgui.ButtonV(otherString, imgui.Vec2{X: nameSize, Y: 0})

				imgui.SameLineV(0, pureMargin)

				imgui.PushItemFlag(imgui.ItemFlagsDisabled, true)

				if imgui.Checkbox("Ready##Guest", &menu.Game.Player.OtherReady) {

				}

				imgui.PopItemFlag()

				imgui.ButtonV("You", imgui.Vec2{X: nameSize, Y: 0})

				imgui.SameLineV(0, pureMargin)

				if imgui.Checkbox("Ready##Player", &menu.Game.Player.Ready) {
					switch menu.Game.Player.Type {
					case game.PlayerTypeServer:
						if len(menu.Game.Server.Server.Sessions) != 0 &&
							menu.Game.Server.Server.Sessions[0] != nil {

							menu.Game.Server.Server.SendMessage(
								menu.Game.Server.Server.Sessions[0],
								game.MessageTypeSetReady,
								strconv.FormatBool(menu.Game.Player.Ready),
							)
						}
						break
					case game.PlayerTypeClient:
						sent := menu.Game.Client.SendMessage(
							game.MessageTypeSetReady,
							strconv.FormatBool(menu.Game.Player.Ready),
						)
						if !sent {
							menu.Game.Player.Reset()
						}
						break
					case game.PlayerTypeUndefined:
						break
					}
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

			if menu.Game.Player.Type == game.PlayerTypeClient ||
				(!menu.Game.Player.OtherReady || !menu.Game.Player.Ready) {

				imgui.PushItemFlag(imgui.ItemFlagsDisabled, true)
				imgui.PushStyleColor(
					imgui.StyleColorButton,
					imgui.Vec4{X: 0.30, Y: 0.30, Z: 0.30, W: 1.0},
				)
			}

			if imgui.ButtonV("Start", imgui.Vec2{X: buttonSize, Y: 0}) {
				fmt.Println("PRESSED!")
			}

			if menu.Game.Player.Type == game.PlayerTypeClient ||
				(!menu.Game.Player.OtherReady || !menu.Game.Player.Ready) {

				imgui.PopStyleColor()
				imgui.PopItemFlag()
			}

			imgui.SameLineV(0, pureMargin)

			if imgui.ButtonV("Abort", imgui.Vec2{X: buttonSize, Y: 0}) {
				switch menu.Game.Player.Type {
				case game.PlayerTypeServer:
					if len(menu.Game.Server.Server.Sessions) > 0 {
						menu.Game.Server.Server.SendMessage(
							menu.Game.Server.Server.Sessions[0],
							game.MessageTypeServerDisconnect, "",
						)
					}
					menu.Game.Server.Server.CloseServer()
					menu.Game.Player.Reset()
					break
				case game.PlayerTypeClient:
					menu.Game.Client.Client.Session.Close()
					menu.Game.Player.Reset()
					break
				}
			}

			if menu.Game.Player.Type == game.PlayerTypeUndefined {
				menu.UIStack.SetSetting("main-menu")
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
