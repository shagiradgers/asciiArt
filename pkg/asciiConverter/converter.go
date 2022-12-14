package asciiConverter

import (
	"fmt"
	"image"
	"io"
)

type Pixel struct {
	R int
	G int
	B int
	A int
}

func GetPixels(file io.Reader) ([][]Pixel, error) {
	img, _, err := image.Decode(file)

	if err != nil {
		return nil, err
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

func DrawAscii(pixels [][]Pixel) {
	density := "|/^~!&*-'Il "
	densityLength := len(density)

	rows := len(pixels)
	columns := len(pixels[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			var index int

			pixel := pixels[i][j]
			average := (pixel.R + pixel.G + pixel.B) / 3

			if average == 0 {
				index = densityLength - 1
			} else {
				index = average % densityLength
			}

			fmt.Printf("%c", density[index])
		}
		fmt.Println()
	}
}
