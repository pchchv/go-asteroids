package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
	// If it goes off the right side of the screen
	if pos.X < -objectSize {
		pos.X = screenWidth + objectSize
	}
	// If it goes off the bottom of the screen
	if pos.Y > screenHeight+objectSize {
		pos.Y = -objectSize
	}
	// If it goes off the top of the screen
	if pos.Y < -objectSize {
		pos.Y = screenHeight + objectSize
	}
}
