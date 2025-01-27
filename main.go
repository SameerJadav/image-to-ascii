package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	"strings"
)

func main() {
	imgPath := flag.String("i", "", "image path")

	flag.Parse()

	if *imgPath == "" {
		log.Fatal("image path not specified")
	}

	file, err := os.Open(*imgPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	asciiChars := [10]string{" ", ".", ":", "i", "1", "C", "G", "8", "%", "@"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		var builder strings.Builder
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			luma := (0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)) / 65535.0
			quantizedLuma := uint8(luma * 9.0)
			builder.WriteString(asciiChars[quantizedLuma])
			builder.WriteString(asciiChars[quantizedLuma])
		}
		fmt.Println(builder.String())
	}
}
