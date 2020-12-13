package kmeans

import (
	"fmt"
	"image"
	"image/color"
	"urban-image-segmentation/internal/gil"
)

type KMeans struct {
	img        image.Image
	width      int
	height     int
	countCntrd int
	cntrd      []centroid
}

type centroid struct {
	X     int
	Y     int
	Color color.RGBA
}

func NewKMeans(img image.Image, countCentroids int) *KMeans {
	k := new(KMeans)

	k.img = img
	k.width = img.Bounds().Max.X
	k.height = img.Bounds().Max.Y
	k.countCntrd = countCentroids

	return k
}

func (k *KMeans) Predict() (image.Image, error) {
	newImg := gil.NewImage(k.width, k.height)

	k.genCentroid()
	fmt.Printf("%+v\n", k.cntrd)
	fmt.Printf("%+v\n", len(k.cntrd))

	return newImg, nil
}

func (k *KMeans) genCentroid() {
	k.cntrd = make([]centroid, 0, k.countCntrd)

	size := k.width * k.height
	step := size / k.countCntrd

	for i := step; i < size; i += step {
		y := i / k.width
		x := i % k.width
		c := k.RGBA32toRGBA8(k.img.At(x, y))

		buff := centroid{
			Y:     y,
			X:     x,
			Color: c,
		}

		k.cntrd = append(k.cntrd, buff)
	}
}

func (k *KMeans) RGBA32toRGBA8(pixel color.Color) color.RGBA {
	r, g, b, a := pixel.RGBA()

	var c color.RGBA

	c.R = uint8(r >> 8)
	c.G = uint8(g >> 8)
	c.B = uint8(b >> 8)
	c.A = uint8(a >> 8)

	return c
}
