package utils

import (
	"image"
	"image/color"
)

func ApplyAlpha(dst *image.RGBA, c color.RGBA, alpha uint, idx int) (int, int, int) {
	// get effective opacity
	aMask := uint8(alpha)
	effA := int(c.A) * int(aMask) / 255
	invA := 255 - effA

	// current RGB values
	r0 := int(dst.Pix[idx])
	g0 := int(dst.Pix[idx+1])
	b0 := int(dst.Pix[idx+2])

	// blended RGB values
	r1 := (effA*int(c.R) + invA*r0 + 127) / 255
	g1 := (effA*int(c.G) + invA*g0 + 127) / 255
	b1 := (effA*int(c.B) + invA*b0 + 127) / 255

	return r1, g1, b1
}
