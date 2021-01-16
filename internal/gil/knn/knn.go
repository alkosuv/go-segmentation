package knn

import (
	"image"
	"image/color"
	"sort"
	"sync"
	"urban-image-segmentation/internal/dataset/label"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/convert"
	"urban-image-segmentation/internal/gil/knn/storage"
	"urban-image-segmentation/internal/gil/math"
)

type KNN struct {
	img    image.Image
	width  int
	height int
	labels *[]storage.Label
}

type distanceLabel struct {
	dist  float64
	index int
}

func NewKNN(img image.Image, labels *[]storage.Label) *KNN {
	k := new(KNN)

	k.width = img.Bounds().Max.X
	k.height = img.Bounds().Max.Y
	k.img = img
	k.labels = labels

	return k
}

func (k *KNN) Predict() (image.Image, error) {
	newImg := gil.NewImage(k.width, k.height)
	var wg sync.WaitGroup

	for x := 0; x < k.width; x++ {
		for y := 0; y < k.height; {

			count := 12
			if k.height-y < count {
				count = k.height - y
			}

			wg.Add(count)
			for c := 0; c < count; c++ {
				go func(wgx, wgy int) {
					defer wg.Done()

					p := convert.RGBA32toRGBA8(k.img.At(wgx, wgy))
					distance := k.evolutionOfDistance(p)
					sort.Slice(*distance, func(i, j int) bool { return (*distance)[i].dist < (*distance)[j].dist })
					*distance = (*distance)[:1000]
					l := k.freqLabels(distance)
					newImg.(*image.RGBA).Set(wgx, wgy, label.Color[l])
				}(x, y)

				y++
			}

			wg.Wait()
		}
	}

	return newImg, nil
}

func (k *KNN) freqLabels(d *[]distanceLabel) int {
	freq := make([]int, len(label.Labels))

	for _, d := range *d {
		freq[d.index]++
	}

	index := 0
	for i, v := range freq {
		if v > freq[index] {
			index = i
		}
	}

	return index
}

func (k *KNN) evolutionOfDistance(point color.RGBA) *[]distanceLabel {
	distance := make([]distanceLabel, 0, len(*k.labels))

	for _, l := range *k.labels {
		d := distanceLabel{
			dist:  math.EuclideanDistance(point, l.RGBA),
			index: l.Index,
		}
		distance = append(distance, d)
	}
	return &distance
}
