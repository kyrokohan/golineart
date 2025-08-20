package img

import (
	"errors"
	"golineart/internal/types"
	"image"
)

func LineDiff(a, b *image.RGBA, line types.Line) (uint64, error) {
	x0, y0, x1, y1 := line.X0, line.Y0, line.X1, line.Y1
	r := a.Bounds()

	if !a.Rect.Eq(b.Rect) {
		return 0, errors.New("images must have identical bounds")
	}

	if x0 < r.Min.X || x1 >= r.Max.X || y0 < r.Min.Y || y1 >= r.Max.Y {
		return 0, errors.New("line has out of bounds pixels")
	}

	var sum uint64
	for i := range len(line.Pixels) {
		// extract the x and y
		x, y := line.Pixels[i].X, line.Pixels[i].Y

		// get the colors of canvas and src
		ar, ag, ab, _ := a.At(x, y).RGBA()
		br, bg, bb, _ := b.At(x, y).RGBA()

		// get MSE of the RGB
		dr, dg, db := int(ar)-int(br), int(ag)-int(bg), int(ab)-int(bb)

		sum += uint64(dr*dr + dg*dg + db*db)
	}

	return sum, nil
}
