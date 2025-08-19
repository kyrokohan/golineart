package canvas

import (
	"image"
	"image/color"
)

func GenerateWhiteCanvas(bounds image.Rectangle) *image.RGBA {
	canvas := image.NewRGBA(bounds)
	canvasBounds := canvas.Bounds()

	for y := canvasBounds.Min.Y; y < canvasBounds.Max.Y; y++ {
		for x := canvasBounds.Min.X; x < canvasBounds.Max.X; x++ {
			// set pixel to white
			canvas.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	return canvas
}
