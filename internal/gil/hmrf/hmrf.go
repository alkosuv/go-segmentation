package hmrf

import (
	"image"
	"urban-image-segmentation/internal/gil"
)

type HMRF struct {
	img     image.Image
	width   int
	height  int
	regions int // count region
}

func NewHMRF(img image.Image, regions int) *HMRF {
	h := new(HMRF)
	h.img = img
	h.width = img.Bounds().Max.X
	h.height = img.Bounds().Max.Y
	h.regions = regions
	return h
}

func (h *HMRF) Predict() image.Image {
	newImg := gil.NewImage(h.width, h.height)
	return newImg
}
