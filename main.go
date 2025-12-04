package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 400
)

func init() {
	// builtin go function which runs before main()
	// setup the raylib window
	rl.InitWindow(screenWidth, screenHeight, "Asteroids")
	rl.SetTargetFPS(60)
}

func deinit() {
	rl.CloseWindow()
}
