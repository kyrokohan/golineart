package img

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"os"
)

func ToGrayscale(imgPath string) *image.RGBA {
	// open the image
	f, err := os.Open(imgPath)
	if err != nil {
		log.Fatal("error opening image file from path!")
	}
	defer f.Close()

	// decode the image
	src, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("error while decoding image: %s", err)
	}

	// create blank canvas
	srcBounds := src.Bounds()
	srcRGBA := image.NewRGBA(srcBounds)
	draw.Draw(srcRGBA, srcBounds, src, srcBounds.Min, draw.Src)

	// grayscale
	for y := srcBounds.Min.Y; y < srcBounds.Max.Y; y++ {
		for x := srcBounds.Min.X; x < srcBounds.Max.X; x++ {
			srcRGBA.Set(x, y, color.Gray16Model.Convert(srcRGBA.At(x, y)))
		}
	}

	return srcRGBA
}
