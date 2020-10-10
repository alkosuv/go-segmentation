//gil - Go Image Library

package gil

import (
	"image"
	"image/png"
	"os"
)

// TODO: rename OpenImagePNG
func OpenImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return png.Decode(file)
}

func SaveImage(path string, img image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func NewImage(width, height int) image.Image {
	// TODO: width, height > 0, если нет, то вернуть ошибку
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	return image.NewRGBA(image.Rectangle{upLeft, lowRight})
}
