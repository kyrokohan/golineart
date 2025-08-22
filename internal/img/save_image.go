package img

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

func SaveImage(img image.Image, o_dir, o_file, o_ext string) {
	// create the "out" directory
	if err := os.MkdirAll(o_dir, 0755); err != nil {
		log.Fatalf("error creating the \"%s\" directory!", o_dir)
	}

	// create and open the output file
	out, err := os.Create(fmt.Sprintf("%s/%s.%s", o_dir, o_file, o_ext))
	if err != nil {
		log.Fatal("error while writing output file!")
	}
	defer out.Close()

	// choose encoder based on extension
	switch strings.ToLower(strings.TrimPrefix(o_ext, ".")) {
	case "png":
		if err := png.Encode(out, img); err != nil {
			log.Fatal("error while encoding PNG image!")
		}
	case "jpg", "jpeg":
		// use a reasonable default quality
		opt := &jpeg.Options{Quality: 95}
		if err := jpeg.Encode(out, img, opt); err != nil {
			log.Fatal("error while encoding JPEG image!")
		}
	default:
		log.Fatalf("unsupported output extension: %s (supported: png, jpg, jpeg)", o_ext)
	}
}
