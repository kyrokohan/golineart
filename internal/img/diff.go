package img

import (
	"errors"
	"image"
	"image/color"

	"github.com/kyrokohan/golineart/internal/types"
	"github.com/kyrokohan/golineart/internal/utils"
)

func checkCompat(a, b *image.RGBA, line types.Line) error {
	x0, y0, x1, y1 := line.X0, line.Y0, line.X1, line.Y1
	r := a.Bounds()

	if !a.Rect.Eq(b.Rect) {
		return errors.New("images must have identical bounds")
	}

	if x0 < r.Min.X || x1 >= r.Max.X || y0 < r.Min.Y || y1 >= r.Max.Y {
		return errors.New("line has out of bounds pixels")
	}

	return nil
}

func LineDiff(a, b *image.RGBA, line types.Line, c color.RGBA, alpha uint, blend bool) (uint64, error) {
	if err := checkCompat(a, b, line); err != nil {
		return 0, err
	}

	var sum uint64
	for i := range len(line.Pixels) {
		x, y := line.Pixels[i].X, line.Pixels[i].Y
		ai := a.PixOffset(x, y)
		bi := b.PixOffset(x, y)

		// to blend or not to blend
		var r1, g1, b1 int
		if blend {
			r1, g1, b1 = utils.ApplyAlpha(a, c, alpha, ai)
		} else {
			r1, g1, b1 = int(a.Pix[ai]), int(a.Pix[ai+1]), int(a.Pix[ai+2])
		}

		// get MSE
		dr := r1 - int(b.Pix[bi])
		dg := g1 - int(b.Pix[bi+1])
		db := b1 - int(b.Pix[bi+2])

		sum += uint64(dr*dr + dg*dg + db*db)
	}
	return sum, nil
}
