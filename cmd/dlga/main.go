package main

import (
	"dark-lines/internal/canvas"
	"dark-lines/internal/img"
	"flag"
	"image/color"
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
	grayscaleImage := img.ToGrayscale(*imgPath)

	// generate same size white canvas
	cnv := canvas.GenerateWhiteCanvas(grayscaleImage.Bounds())

	// bounds := cnv.Bounds()
	// randomStartX := rand.Intn(bounds.Max.X)
	// randomStartY := rand.Intn(bounds.Max.Y)

	for range 100 {
		canvas.DrawRandomLine(cnv, color.Black)
	}

	// save image
	img.SaveImage(cnv, OUTPUT_DIRECTORY, OUTPUT_FILE_NAME, OUTPUT_FILE_EXTENSION)

	// Turn image into grayscale and store image data

	// Create a blank canvas the same size as the image

	// Draw 1000 random lines on the canvas

	// Save the canvas as a png

}
