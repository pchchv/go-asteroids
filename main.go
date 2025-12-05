package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	tileSize      = 64
	screenWidth   = 800
	screenHeight  = 400
	rotationSpeed = 2.0
	playerSpeed   = 6.0
)

var (
	player        Player
	texTiles      rl.Texture2D
	texBackground rl.Texture2D
	spriteRec     rl.Rectangle
	boostRec      rl.Rectangle
)

type Player struct {
	isBoosting   bool
	rotation     float32
	acceleration float32
	position     rl.Vector2
	speed        rl.Vector2
	size         rl.Vector2
}

func (p *Player) Draw() {
	destTexture := rl.Rectangle{X: p.position.X, Y: p.position.Y, Width: p.size.X, Height: p.size.Y}
	if p.isBoosting {
		rl.DrawTexturePro(
			texTiles,
			boostRec,
			destTexture,
			rl.Vector2{X: p.size.X / 2, Y: p.size.Y/2 - 40},
			p.rotation,
			rl.White,
		)
	} else {
		rl.DrawTexturePro(
			texTiles,
			spriteRec,
			destTexture,
			rl.Vector2{X: p.size.X / 2, Y: p.size.Y / 2},
			p.rotation,
			rl.White,
		)
	}
}
func (p *Player) Update() {
	// rotate the player with the arrow keys
	if rl.IsKeyDown(rl.KeyLeft) {
		player.rotation -= rotationSpeed
	}

	if rl.IsKeyDown(rl.KeyRight) {
		player.rotation += rotationSpeed
	}
	// default to not boosting
	player.isBoosting = false

	// accelerate the player with up
	if rl.IsKeyDown(rl.KeyUp) {
		if player.acceleration < 0.9 {
			player.acceleration += 0.1
		}
		player.isBoosting = true
	}

	// decellerate the player with down
	if rl.IsKeyDown(rl.KeyDown) {
		if player.acceleration > 0 {
			player.acceleration -= 0.05
		}

		if player.acceleration < 0 {
			player.acceleration = 0

		}
	}

	// get the direction the sprite is pointing
	direction := getDirectionVector(player.rotation)

	// start to move to the direction
	player.speed = rl.Vector2Scale(direction, playerSpeed)

	// accelerate in that direction
	player.position.X += player.speed.X * player.acceleration
	player.position.Y -= player.speed.Y * player.acceleration

	// to void losing our ship, wrap around the screen
	wrapPosition(&p.position, tileSize)
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

func update() {
	player.Update()
}

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

func getDirectionVector(rotation float32) rl.Vector2 {
	// convert the rotation to radians
	radians := float64(rotation) * rl.Deg2rad

	// return the vector of the direction we are pointing at
	return rl.Vector2{
		X: float32(math.Sin(radians)),
		Y: float32(math.Cos(radians)),
	}
}

func wrapPosition(pos *rl.Vector2, objectSize float32) {
	// if it goes off the left side of the screen
	if pos.X > screenWidth+objectSize {
		pos.X = -objectSize
	}
	// if it goes off the right side of the screen
	if pos.X < -objectSize {
		pos.X = screenWidth + objectSize
	}
	// if it goes off the bottom of the screen
	if pos.Y > screenHeight+objectSize {
		pos.Y = -objectSize
	}
	// if it goes off the top of the screen
	if pos.Y < -objectSize {
		pos.Y = screenHeight + objectSize
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
