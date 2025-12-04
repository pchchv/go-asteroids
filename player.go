package main

import rl "github.com/gen2brain/raylib-go/raylib"

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
