package main

import (
	"dark-lines/internal/canvas"
	"dark-lines/internal/img"
	"flag"
	"fmt"
	"image/color"
	_ "image/jpeg"
	"log"
	"math/rand/v2"
	"strconv"
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

	// generate a bunch of lines to get a good result
	for i := range 10000 {
		bow := rand.IntN(2)

		// choose randomly between black or white line
		c := color.RGBA{255, 255, 255, 255}
		if bow == 1 {
			c = color.RGBA{0, 0, 0, 255}
		}

		// choose and draw the best out of N lines
		canvas.DrawBestOfNLines(cnv, grayscaleImage, 100, c, 0)

		// save img and log progress periodically
		if i%100 == 0 {
			img.SaveImage(cnv, OUTPUT_DIRECTORY, strconv.Itoa(i), OUTPUT_FILE_EXTENSION)
		}
		fmt.Printf("Generated %d/%d\r", i+1, 10000)
	}

	// save the final image
	img.SaveImage(cnv, OUTPUT_DIRECTORY, "final", OUTPUT_FILE_EXTENSION)
	fmt.Println("\nDone!")
}
