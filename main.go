package main

import (
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"online-game/game"
	"online-game/scenes"
	"time"
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Go RCP",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	sceneManager := game.SceneManager{}

	menuScene := scenes.NewMenuScene(win)

	sceneManager.AddScene(menuScene)

	sceneManager.SetScene("menu-scene")

	last := time.Now()
	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		win.Clear(pixel.RGB(0.06, 0.01, 0.12))

		sceneManager.Update(dt)

		win.Update()
	}
}

func main() {

	fmt.Println(game.GetRelativePath("res/test.png"))

	pixelgl.Run(run)
}
