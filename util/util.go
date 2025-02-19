package util

import (
	"image"
	"image/color"
)

func GetImageColors(image image.Image) map[color.Color]struct{} {
	colors := make(map[color.Color]struct{})
	var empty struct{}

	for y := 0; y < image.Bounds().Dy(); y++ {
		for x := 0; x < image.Bounds().Dx(); x++ {
			colors[image.At(x, y)] = empty
		}
	}
	return colors
}
