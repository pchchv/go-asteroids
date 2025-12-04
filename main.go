package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	screenWidth  = 800
	screenHeight = 400
)

var texBackground rl.Texture2D

func init() {
	// builtin go function which runs before main()
	// setup the raylib window
	rl.InitWindow(screenWidth, screenHeight, "Asteroids")
	rl.SetTargetFPS(60)

	// load textures
	texBackground = rl.LoadTexture("resources/space_background.png")
}

func deinit() {
	rl.CloseWindow()
	// unload textures when the game closes
	rl.UnloadTexture(texBackground)
}

func draw() {
	rl.BeginDrawing()
	// set the background to black
	rl.ClearBackground(rl.Black)
	// draw the score to the screen
	rl.DrawText("Score 0", 10, 10, 20, rl.Gray)
	rl.EndDrawing()

}

// TODO: update the state
func update() {}

func main() {
	// when the main function ends,
	// call the deinit() function
	defer deinit()

	// continue the loop until the window is closed or ESC is pressed
	for !rl.WindowShouldClose() {
		draw()
		update()
	}
}
