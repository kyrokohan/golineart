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
	"os"
	"strconv"
)

const (
	OUTPUT_DIRECTORY = "out"
	OUTPUT_FILE      = "final"
	OUTPUT_EXTENSION = "png"
)

func main() {

	// get image path from last argument
	imgPath := os.Args[len(os.Args)-1]

	if imgPath == "" || len(os.Args) < 2 {
		log.Fatal("missing image path!")
	}

	// set up and parse flags
	rounds := flag.Int("rounds", 10000, "total number of rounds (lines) to generate")
	linesPerRound := flag.Int("lines", 100, "total number of lines to choose from per round")
	outputDir := flag.String("odir", OUTPUT_DIRECTORY, "output directory")
	outputFile := flag.String("ofile", OUTPUT_FILE, "name of the final output file")
	outputExt := flag.String("oext", OUTPUT_EXTENSION, "output extension")
	saveFrequency := flag.Int("sfreq", 0, "how many generated lines before saving the image")
	alpha := flag.Uint("alpha", 51, "opacity of the lines generated between [0, 255]")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  %s [...flags] [image path]\n\nFlags:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	if *alpha > 255 {
		log.Fatal("alpha must be between 0 and 255 (inclusive)")
	}

	// process image (turn to grayscale)
	grayscaleImage := img.ToGrayscale(imgPath)

	// generate same size white canvas
	cnv := canvas.GenerateWhiteCanvas(grayscaleImage.Bounds())

	// generate a bunch of lines to get a good result
	for i := range *rounds {
		bow := rand.IntN(2)

		// choose randomly between black or white line
		c := color.RGBA{255, 255, 255, 255}
		if bow == 1 {
			c = color.RGBA{0, 0, 0, 255}
		}

		// choose and draw the best out of N lines
		canvas.DrawBestOfNLines(cnv, grayscaleImage, *linesPerRound, c, *alpha, 0)

		// save img and log progress periodically
		if *saveFrequency != 0 && i%*saveFrequency == 0 {
			img.SaveImage(cnv, *outputDir, strconv.Itoa(i), *outputExt)
		}
		fmt.Printf("Generated %d/%d\r", i+1, *rounds)
	}

	// save the final image
	img.SaveImage(cnv, *outputDir, *outputFile, *outputExt)
	fmt.Println("\nDone!")
}
