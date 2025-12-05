package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	tileSize     = 64
	screenWidth  = 800
	screenHeight = 400
)

var (
	player        Player
	texTiles      rl.Texture2D
	texBackground rl.Texture2D
	spriteRec     rl.Rectangle
)

func init() {
	// builtin go function which runs before main()
	// setup the raylib window
	rl.InitWindow(screenWidth, screenHeight, "Asteroids")
	rl.SetTargetFPS(60)

	// load textures
	texTiles = rl.LoadTexture("resources/tilesheet.png")
	texBackground = rl.LoadTexture("resources/space_background.png")
	spriteRec = rl.Rectangle{X: tileSize * 0, Y: tileSize * 2, Width: tileSize, Height: tileSize}

	initGame()
}

func deinit() {
	rl.CloseWindow()
	// unload textures when the game closes
	rl.UnloadTexture(texTiles)
	rl.UnloadTexture(texBackground)
}

func draw() {
	rl.BeginDrawing()
	// set the background to a nebula
	bgDest := rl.Rectangle{X: 0, Y: 0, Width: screenWidth, Height: screenHeight}
	bgSource := rl.Rectangle{X: 0, Y: 0, Width: float32(texBackground.Width), Height: float32(texBackground.Height)}
	rl.DrawTexturePro(texBackground, bgSource, bgDest, rl.Vector2{X: 0, Y: 0}, 0, rl.White)

	// draw the player
	player.Draw()

	// draw the score to the screen
	rl.DrawText("Score 0", 10, 10, 20, rl.Gray)
	rl.EndDrawing()

}

// TODO: update the state
func update() {}

func initGame() {
	player = Player{
		position:     rl.Vector2{X: 400, Y: 200},
		speed:        rl.Vector2{X: 0.0, Y: 0.0},
		size:         rl.Vector2{X: tileSize, Y: tileSize},
		rotation:     0.0,
		acceleration: 0.0,
		isBoosting:   false,
	}
}

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
