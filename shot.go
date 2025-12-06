package main

import rl "github.com/gen2brain/raylib-go/raylib"

const shotSpeed = 8.0

type Shot struct {
	speed    rl.Vector2
	position rl.Vector2
	radius   float32
	active   bool
}

func (s *Shot) Draw() {
	if s.active {
		rl.DrawCircleV(s.position, s.radius, rl.Yellow)
	}
}

func (s *Shot) Update() {
	if s.active {
		s.position.X += s.speed.X
		s.position.Y -= s.speed.Y
		if s.position.X < 0 || s.position.X > screenWidth || s.position.Y < 0 || s.position.Y > screenHeight {
			s.active = false
		}
	}
}

func fireShot() {
	for i := range shots {
		// find the first inactive shot
		if !shots[i].active {
			// start at the players position
			shots[i].position = player.position
			shots[i].active = true

			// get the players direction
			shotDirection := getDirectionVector(player.rotation)

			// get the initial velocity
			shotVelocity := rl.Vector2Scale(shotDirection, shotSpeed)
			// account for the players speed
			playerVelocity := rl.Vector2Scale(player.speed, player.acceleration)

			// fire the shot, realative to the players speed
			shots[i].speed = rl.Vector2Add(playerVelocity, shotVelocity)

			shots[i].radius = 2
			// break after one shot
			break
		}
	}
}
