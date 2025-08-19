package img

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
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

	if err := png.Encode(out, img); err != nil {
		log.Fatal("error while encoding image!")
	}
}
