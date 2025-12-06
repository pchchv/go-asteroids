package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screenWidth      = 800
	screenHeight     = 400
	tileSize         = 64
	maxShots         = 10
	initialAsteroids = 5
)

var (
	asteriodsDestroyed int
	gameOver           bool
	victory            bool
	paused             bool
	shots              []Shot
	asteroids          []Asteroid
	texTiles           rl.Texture2D
	texBackground      rl.Texture2D
)

func initGame() {
	// start with it not being game over, pauses or victory
	paused, victory, gameOver = false, false, false

	// reset score
	asteriodsDestroyed = 0

	// create the asteroids field
	asteroids = nil
	for range initialAsteroids {
		asteroids = append(asteroids, createLargeAsteroid())
	}

	// create the laser shots
	for i := range shots {
		shots[i].active = false
	}

	player = Player{
		position:     rl.Vector2{X: 400, Y: 200},
		speed:        rl.Vector2{X: 0.0, Y: 0.0},
		size:         rl.Vector2{X: tileSize, Y: tileSize},
		rotation:     0.0,
		acceleration: 0.0,
		isBoosting:   false,
	}
}

func init() {
	// builtin go function which runs before main()
	// setup the raylib window
	rl.InitWindow(screenWidth, screenHeight, "Asteroids")
	rl.SetTargetFPS(60)

	// load textures
	texTiles = rl.LoadTexture("assets/tilesheet.png")
	texBackground = rl.LoadTexture("assets/space_background.png")

	// sprites for the ship and it boost
	spriteRec = rl.Rectangle{X: tileSize * 0, Y: tileSize * 2, Width: tileSize, Height: tileSize}
	boostRec = rl.Rectangle{X: tileSize * 7, Y: tileSize * 5, Width: tileSize, Height: tileSize}

	// sprite for the asteroid
	asteroidRec = rl.Rectangle{X: tileSize * 1, Y: tileSize * 4, Width: tileSize, Height: tileSize}

	// create the shots
	shots = make([]Shot, maxShots)

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

	// draw the asteroid field
	for i := range asteroids {
		asteroids[i].Draw()
	}

	// draw the shots
	for i := range shots {
		shots[i].Draw()
	}

	if gameOver {
		drawCenteredText("Game over", screenHeight/2, 50, rl.Red)
		drawCenteredText("Press R to restart", screenHeight/2+60, 20, rl.DarkGray)
	}

	if victory {
		drawCenteredText("YOU WIN!", screenHeight/2, 50, rl.Gray)
		drawCenteredText("Press R to restart", screenHeight/2+60, 20, rl.RayWhite)
	}

	// draw the score to the screen
	rl.DrawText(fmt.Sprintf("Score %d", asteriodsDestroyed), 10, 10, 20, rl.Gray)
	pauseTextSize := rl.MeasureText("[P]ause", 20)
	rl.DrawText("[P]ause", screenWidth-pauseTextSize-10, 10, 20, rl.Gray)

	rl.EndDrawing()

}

func update() {
	// if there are no asteroids left, we in
	if len(asteroids) == 0 {
		victory = true
	}

	// toggle paused
	if rl.IsKeyPressed('P') {
		paused = !paused
	}

	// restart the game
	if (gameOver || victory) && rl.IsKeyPressed('R') {
		initGame()
	}

	// if it is not game over, update the frame
	if !paused && !victory && !gameOver {
		player.Update()

		// update the asteroid field
		for i := range asteroids {
			asteroids[i].Update()
		}

		// update the shots
		for i := range shots {
			shots[i].Update()
		}

		checkCollisions()
	}
}

func main() {
	// when the main function ends, call the deinit() function
	defer deinit()

	// continue the loop until the window is closed or ESC is pressed
	for !rl.WindowShouldClose() {
		draw()
		update()
	}
}
