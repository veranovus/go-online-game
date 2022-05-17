package main

import (
	"online-game/client"
	"online-game/game"
	"online-game/ncom"
	"online-game/scenes"
	"online-game/server"
	"time"

	"github.com/dusk125/pixelui"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Go RPC",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	// Create pixel UI
	ui := pixelui.NewUI(win, 0)
	defer ui.Destroy()

	// Set style
	game.SetImGUIStyle()

	// Scene manager
	sceneManager := game.SceneManager{}

	// Create Server
	s := &server.Server{
		Channel: make(chan ncom.Message),
	}
	c := &client.Client{
		Channel:       make(chan ncom.Message),
		Authenticated: false,
	}

	// Create and add scenes
	menuScene := scenes.NewMenuScene(win, ui, s, c)
	sceneManager.AddScene(menuScene)

	gameScene := scenes.NewGameScene(win, ui, s, c)
	sceneManager.AddScene(gameScene)

	// Set initial scene
	sceneManager.SetScene("menu-scene")

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		ui.NewFrame()
		win.Clear(pixel.RGB(0.03, 0.005, 0.06))

		sceneManager.Update(dt)

		ui.Draw(win)
		win.Update()
	}
}

func main() {

	pixelgl.Run(run)
}
