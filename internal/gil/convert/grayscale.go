package convert

import (
	"image"
	"image/color"
	"urban-image-segmentation/internal/gil"
)

func Grayscale(img image.Image) image.Image {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	grayImg := gil.NewImage(width, height)

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			r, g, b, _ := img.At(i, j).RGBA()
			buff := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			pixel := color.Gray16{uint16(buff)}
			grayImg.(*image.RGBA).Set(i, j, pixel)
		}
	}
	return grayImg
}
