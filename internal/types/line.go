package types

import "image"

type Line struct {
	X0, Y0, X1, Y1 int
	Pixels         []image.Point
}
