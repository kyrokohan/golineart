package img

import (
	"errors"
	"image"
)

func ImageDifference(a, b *image.RGBA) (float64, error) {
	if !a.Rect.Eq(b.Rect) {
		return 0, errors.New("images must have identical bounds")
	}
	r := a.Rect
	w, h := r.Dx(), r.Dy()

	// get mean squared error over RGB of the image.
	var sum uint64
	for y := r.Min.Y; y < r.Max.Y; y++ {
		ia := a.PixOffset(r.Min.X, y)
		ib := b.PixOffset(r.Min.X, y)
		for range w {
			dr := int(a.Pix[ia+0]) - int(b.Pix[ib+0])
			dg := int(a.Pix[ia+1]) - int(b.Pix[ib+1])
			db := int(a.Pix[ia+2]) - int(b.Pix[ib+2])
			sum += uint64(dr*dr + dg*dg + db*db)
			ia += 4
			ib += 4
		}
	}
	return float64(sum) / float64(3*w*h), nil
}
