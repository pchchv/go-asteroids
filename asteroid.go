package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Small AsteroidSize = iota
	Large
	Medium
)

var asteroidRec rl.Rectangle

// Enum for storing the size of the asteroid
type AsteroidSize int

type Asteroid struct {
	size         rl.Vector2
	speed        rl.Vector2
	position     rl.Vector2
	asteroidSize AsteroidSize
}

// Draw draws the asteroid to the screen.
func (a *Asteroid) Draw() {
	destTexture := rl.Rectangle{X: a.position.X, Y: a.position.Y, Width: a.size.X, Height: a.size.Y}
	rl.DrawTexturePro(
		texTiles,
		asteroidRec,
		destTexture,
		rl.Vector2{X: a.size.X / 2, Y: a.size.Y / 2},
		0.0,
		rl.White,
	)
}

func (a *Asteroid) Update() {
	// move the asteroid in its direction
	a.position = rl.Vector2Add(a.position, a.speed)

	// wrap the position, so they are always on screen
	wrapPosition(&a.position, a.size.X)
}

func createAsteroid(asteroidSize AsteroidSize, position, speed rl.Vector2) Asteroid {
	// scale the image of the asteroid based on the asteroidSize
	var size rl.Vector2
	switch asteroidSize {
	case Small:
		size = rl.Vector2{X: tileSize * 0.4, Y: tileSize * 0.4}
	case Medium:
		size = rl.Vector2{X: tileSize * 0.7, Y: tileSize * 0.7}
	case Large:
		size = rl.Vector2{X: tileSize * 1.0, Y: tileSize * 1.0}
	}

	return Asteroid{
		position:     position,
		speed:        speed,
		size:         size,
		asteroidSize: asteroidSize,
	}
}

func splitAsteroid(asteroid Asteroid) {
	// do nothing for small
	if asteroid.asteroidSize == Small {
		return
	}

	// work out how many splits to do
	var split int
	var newSize AsteroidSize
	if asteroid.asteroidSize == Large {
		split, newSize = 2, Medium
	} else {
		split, newSize = 4, Small
	}

	// create the new smaller asteroids
	for range split {
		angle := float64(rl.GetRandomValue(0, 360))
		direction := getDirectionVector(float32(angle))
		speed := rl.Vector2Scale(direction, 2.0)
		newAsteroid := createAsteroid(newSize, asteroid.position, speed)
		asteroids = append(asteroids, newAsteroid)
	}
}

func createLargeAsteroid() Asteroid {
	// generate a random position on screen
	randomX := float32(rl.GetRandomValue(0, screenWidth))
	randomY := float32(rl.GetRandomValue(0, screenHeight))

	// generate a random edge of the screen to spawn
	var position rl.Vector2
	randomEdge := rl.GetRandomValue(0, 3)
	switch randomEdge {
	case 0:
		position = rl.Vector2{X: randomX, Y: +tileSize}
	case 1:
		position = rl.Vector2{X: screenWidth + tileSize, Y: randomY}
	case 2:
		position = rl.Vector2{X: randomX, Y: screenHeight + tileSize}
	case 3:
		position = rl.Vector2{X: -tileSize, Y: randomY}
	}

	// generate a random speed and direction for the asteroid
	speed := rl.Vector2{
		X: float32(rl.GetRandomValue(-10, 10)) / 10,
		Y: float32(rl.GetRandomValue(-10, 10)) / 10,
	}

	// create the large asteroid
	return createAsteroid(Large, position, speed)
}
