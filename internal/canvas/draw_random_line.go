package canvas

import (
	"image"
	"image/color"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func DrawLine(dst *image.RGBA, x0, y0, x1, y1 int, c color.Color) {
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

}
