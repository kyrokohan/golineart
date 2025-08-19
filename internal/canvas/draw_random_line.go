package canvas

import (
	"image"
	"image/color"
	"math/rand"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func generateRandomLineCoordinates(b image.Rectangle) (int, int, int, int) {
	// get widths
	dx := b.Dx()
	dy := b.Dy()

	// get random edges while making sure same edge never gets chosen
	e0 := rand.Intn(4)
	e1 := rand.Intn(3)
	if e1 >= e0 {
		e1++
	}

	// helper function to get point on an edge
	pointOnEdge := func(edge int) (int, int) {
		switch edge {
		case 0: // top
			return b.Min.X + rand.Intn(dx), b.Min.Y
		case 1: // right
			return b.Max.X - 1, b.Min.Y + rand.Intn(dy)
		case 2: // bottom
			return b.Min.X + rand.Intn(dx), b.Max.Y - 1
		default: // left
			return b.Min.X, b.Min.Y + rand.Intn(dy)
		}
	}

	x0, y0 := pointOnEdge(e0)
	x1, y1 := pointOnEdge(e1)

	return x0, y0, x1, y1
}

func DrawRandomLine(dst *image.RGBA, c color.Color) (int, int, int, int) {
	// get start and end points of the random line
	x0, y0, x1, y1 := generateRandomLineCoordinates(dst.Bounds())

	// the rest of this is using Bresenham's line drawing algorithm (the error version)
	dx := abs(x1 - x0)

	sx := -1
	if x0 < x1 {
		sx = 1
	}

	dy := -abs(y1 - y0)
	sy := -1
	if y0 < y1 {
		sy = 1
	}

	err := dx + dy

	for {
		dst.Set(x0, y0, c)
		e2 := 2 * err

		if e2 >= dy {
			if x0 == x1 {
				break
			}
			err = err + dy
			x0 = x0 + sx
		}

		if e2 <= dx {
			if y0 == y1 {
				break
			}
			err = err + dx
			y0 = y0 + sy
		}
	}

	// return the starting and ending coords
	return x0, y0, x1, y1
}
