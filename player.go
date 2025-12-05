package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	rotationSpeed = 2.0
	playerSpeed   = 6.0
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
	rl.DrawTexturePro(
		texTiles,
		spriteRec,
		destTexture,
		rl.Vector2{X: p.size.X / 2, Y: p.size.Y / 2},
		p.rotation,
		rl.White,
	)
}

func (p *Player) Update() {
	// rotate the player with the arrow keys
	if rl.IsKeyDown(rl.KeyLeft) {
		player.rotation -= rotationSpeed
	}

	if rl.IsKeyDown(rl.KeyRight) {
		player.rotation += rotationSpeed
	}

	// accelerate the player with up
	if rl.IsKeyDown(rl.KeyUp) {
		if player.acceleration < 0.9 {
			player.acceleration += 0.1
		}
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
