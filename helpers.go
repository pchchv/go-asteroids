package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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

func getDirectionVector(rotation float32) rl.Vector2 {
	// convert the rotation to radians
	radians := float64(rotation) * rl.Deg2rad

	// return the vector of the direction we are pointing at
	return rl.Vector2{
		X: float32(math.Sin(radians)),
		Y: float32(math.Cos(radians)),
	}
}

func checkCollisions() {
	for i := len(asteroids) - 1; i >= 0; i-- {
		// check for collision between player and asteroid
		if rl.CheckCollisionCircles(
			player.position,
			player.size.X/4,
			asteroids[i].position,
			asteroids[i].size.X/4,
		) {
			gameOver = true
		}

		// check for a collision between shots and the asteroid
		for j := range shots {
			// loop through all the active shots
			// if it has collided with an asteroid
			if shots[j].active && rl.CheckCollisionCircles(shots[j].position, shots[j].radius, asteroids[i].position, asteroids[i].size.X/2) {
				// destroy the shot and split the asteroid
				shots[j].active = false

				// the asteroid shot split according to our rules
				splitAsteroid(asteroids[i])

				// remove the original asteroid from the slice
				asteroids = append(asteroids[:i], asteroids[i+1:]...)

				// increase our score
				asteriodsDestroyed++
				break
			}
		}
	}
}

func drawCenteredText(text string, y, fontSize int32, color rl.Color) {
	textWidth := rl.MeasureText(text, fontSize)
	rl.DrawText(text, screenWidth/2-textWidth/2, y, fontSize, color)
}
