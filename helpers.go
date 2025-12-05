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
