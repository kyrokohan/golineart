package img

import "image"

func CloneImage(src *image.RGBA) *image.RGBA {
	r := src.Bounds()
	dst := image.NewRGBA(r)
	w := r.Dx()

	for y := r.Min.Y; y < r.Max.Y; y++ {
		index := src.PixOffset(r.Min.X, y)
		copy(dst.Pix[index:index+4*w], src.Pix[index:index+4*w])
	}

	return dst
}
