package knn

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"sort"
	"sync"
	"time"
	"urban-image-segmentation/internal/dataset/label"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/grayscale"
	"urban-image-segmentation/internal/gil/knn/storage"
)

type KNN struct {
	img    image.Image
	width  int
	height int
	labels *[]storage.Label
}

type DistanceLabel struct {
	dist  float64
	index int
}

func NewKNN(img image.Image, labels *[]storage.Label) *KNN {
	k := new(KNN)

	k.width = img.Bounds().Max.X
	k.height = img.Bounds().Max.Y
	k.img = grayscale.RGBA2GRAY(img)
	k.labels = labels

	return k
}

func (k *KNN) Predict() (image.Image, error) {
	newImg := gil.NewImage(k.width, k.height)
	var wg sync.WaitGroup

	start := time.Now()

	for x := 0; x < k.width/8; x++ {
		for y := 0; y < k.height/4; {

			count := 8
			if k.height-x < count {
				count = (k.width - x) % count
			}

			wg.Add(count)
			for c := 0; c < count; c++ {
				go func(wgx, wgy int) {
					defer wg.Done()

					p := k.RGBA32toRGBA8(k.img.At(wgx, wgy))
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

	fmt.Println(time.Since(start))
	return newImg, nil
}

func (k *KNN) freqLabels(d *[]DistanceLabel) int {
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

func (k *KNN) evolutionOfDistance(point color.RGBA) *[]DistanceLabel {
	distance := make([]DistanceLabel, 0, len(*k.labels))

	for _, l := range *k.labels {
		d := DistanceLabel{
			dist:  k.distance(point, l.RGBA),
			index: l.Index,
		}
		distance = append(distance, d)
	}
	return &distance
}

func (k *KNN) distance(point1, point2 color.RGBA) float64 {
	r := float64(point1.R - point2.R)
	g := float64(point1.G - point2.G)
	b := float64(point1.B - point2.B)

	return math.Sqrt(math.Pow(r, 2) + math.Pow(g, 2) + math.Pow(b, 2))
}

// TODO: Написать
func (k *KNN) distributionByRegion() {
}

func (k *KNN) RGBA32toRGBA8(pixel color.Color) color.RGBA {
	r, g, b, a := pixel.RGBA()

	var c color.RGBA

	c.R = uint8(r >> 8)
	c.G = uint8(g >> 8)
	c.B = uint8(b >> 8)
	c.A = uint8(a >> 8)

	return c
}
