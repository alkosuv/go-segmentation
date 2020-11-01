package knn

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"urban-image-segmentation/internal/gil/grayscale"
	"urban-image-segmentation/internal/gil/knn/storage"
)

type KNN struct {
	img        image.Image
	width      int
	height     int
	setOfClass *[]storage.Label

	distMatrix [][]float64
	region     [][]uint8
}

func NewKNN(img image.Image, setOfClass *[]storage.Label) *KNN {
	k := new(KNN)

	k.width = img.Bounds().Max.X
	k.height = img.Bounds().Max.Y
	k.img = grayscale.RGBA2GRAY(img)
	k.setOfClass = setOfClass

	k.distMatrix = make([][]float64, k.width)
	for i := range k.distMatrix {
		k.distMatrix[i] = make([]float64, k.height)
	}

	k.region = make([][]uint8, k.width)
	for i := range k.region {
		k.region[i] = make([]uint8, k.width)
	}

	return k
}

func (k *KNN) Predict() (image.Image, error) {
	// newImg := gil.NewImage(k.width, k.height)

	k.evaluationOfDistance()
	k.distributionByRegion()

	return k.img, nil
}

func (k *KNN) evaluationOfDistance() {
	for i := 0; i < k.width; i++ {
		for j := 0; j < k.height; j++ {
			k.distMatrix[i][j] = 65025
			for l := 0; l < k.height; l++ {
				pxl1 := k.RGBA32toRGBA8(k.img.At(i, j))
				pxl2 := k.RGBA32toRGBA8(k.img.At(i, l))

				dist := k.distance(pxl1, pxl2)
				if dist < k.distMatrix[i][j] && j != l {
					k.distMatrix[i][j] = dist
				}
			}
		}
	}
}

func (k *KNN) distance(point1, point2 color.RGBA) float64 {
	return math.Sqrt(math.Pow(float64(point1.R)-float64(point2.R), 2))
}

func (k *KNN) minMax() (min float64, max float64) {
	min = 255.0
	max = -1.0
	for i := 0; i < k.width; i++ {
		for j := 0; j < k.height; j++ {
			if min > k.distMatrix[i][j] {
				min = k.distMatrix[i][j]
			}
			if max < k.distMatrix[i][j] {
				max = k.distMatrix[i][j]
			}
		}
	}

	return
}

func (k *KNN) distributionByRegion() {
	min, max := k.minMax()

	test1 := make(map[float64]int)
	test2 := make(map[uint32]int)

	step := (max - min) / 16.0
	for i := 0; i < k.width; i++ {
		for j := 0; j < k.height; j++ {
			test1[k.distMatrix[i][j]]++
			r, _, _, _ := k.img.At(i, j).RGBA()
			test2[r>>8]++

			var n uint8 = 0
			for l := min; l < max; l += step {
				if k.distMatrix[i][j] <= l {
					k.region[i][j] = n
				}
				n++
			}
		}
	}

	fmt.Println(len(test1))
	fmt.Println(len(test2))
}

// dp - decimal places
func (k *KNN) round(x float64, dp int) float64 {
	pow := math.Pow10(dp)
	return math.Round(x*pow) / pow
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
