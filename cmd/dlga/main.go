package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	"log"
	"os"
)

const (
	OUTPUT_DIRECTORY      = "out"
	OUTPUT_FILE_NAME      = "out"
	OUTPUT_FILE_EXTENSION = "png"
)

func main() {

	// set up flags
	imgPath := flag.String("image", "", "Input base image")
	flag.Parse()

	if *imgPath == "" {
		log.Fatal("missing image path flag!")
	}

	// open the image
	f, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal("error opening image file from path!")
	}
	defer f.Close()

	// decode the image
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("error while decoding image: %s", err)
	}

	// create the "out" directory
	if err := os.MkdirAll(OUTPUT_DIRECTORY, 0755); err != nil {
		log.Fatalf("error creating the \"%s\" directory!", OUTPUT_DIRECTORY)
	}

	// create and open the output file
	out, err := os.Create(fmt.Sprintf("%s/%s.%s", OUTPUT_DIRECTORY, OUTPUT_FILE_NAME, OUTPUT_FILE_EXTENSION))
	if err != nil {
		log.Fatal("error while writing output file!")
	}
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		log.Fatal("error while encoding image!")
	}

	// Turn image into grayscale and store image data

	// Create a blank canvas the same size as the image

	// Draw 1000 random lines on the canvas

	// Save the canvas as a png

}
