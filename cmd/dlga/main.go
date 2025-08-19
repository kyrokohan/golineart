package main

import (
	"dark-lines/internal/img"
	"flag"
	_ "image/jpeg"
	"log"
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

	// process image (turn to grayscale)
	grayscaleImage := img.ProcessImage(*imgPath)

	// save image
	img.SaveImage(grayscaleImage, OUTPUT_DIRECTORY, OUTPUT_FILE_NAME, OUTPUT_FILE_EXTENSION)

	// Turn image into grayscale and store image data

	// Create a blank canvas the same size as the image

	// Draw 1000 random lines on the canvas

	// Save the canvas as a png

}
