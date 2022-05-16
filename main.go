package main

import (
	"online-game/game"
	"online-game/scenes"
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

	// Create and add scenes
	menuScene := scenes.NewMenuScene(win, ui)
	sceneManager.AddScene(menuScene)

	gameScene := scenes.NewGameScene(win, ui)
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
