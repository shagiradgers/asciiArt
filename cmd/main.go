package main

import (
	"asciiArt/pkg/asciiConverter"
	"fmt"
	"image"
	"image/png"
	"os"
)

func main() {
	var imageName string
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	fmt.Println("Enter image name (png)")
	_, err := fmt.Scan(&imageName)
	if err != nil {
		fmt.Println("Error, then getting image name")
	}
	file, err := os.Open(imageName)

	if err != nil {
		fmt.Println("Error: File could not be opened")
		os.Exit(1)
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			return
		}
	}(file)

	pixels, err := asciiConverter.GetPixels(file)

	if err != nil {
		fmt.Println("Error: Image could not be decoded")
		os.Exit(1)
	}

	asciiConverter.DrawAscii(pixels)
}
