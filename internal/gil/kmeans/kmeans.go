package kmeans

import (
	"fmt"
	"image"
	"image/color"
	"sort"
	"urban-image-segmentation/internal/gil"
	"urban-image-segmentation/internal/gil/convert"
	"urban-image-segmentation/internal/gil/math"
)

type KMeans struct {
	pixels    []pixel
	width     int
	height    int
	centroids []pixel
	count     int // count centroids
}

type pixel struct {
	x        int
	y        int
	clr      color.RGBA
	centroid int // index centroid
}

func NewKMeans(img image.Image, count int) *KMeans {
	pixels := imgToPixel(img)

	k := new(KMeans)
	k.width = img.Bounds().Max.X
	k.height = img.Bounds().Max.Y
	k.centroids = genCentroid(pixels, count)
	k.count = count
	k.pixels = pixels

	return k
}

func (k *KMeans) Predict() (image.Image, error) {
	for i := 0; i < 10; i++ {
		k.evaluateCentroids()
		k.newCentroids()
	}

	return k.image(), nil
}

func (k *KMeans) evaluateCentroids() {
	for i, p := range k.pixels {
		buff := make([]float64, 0, k.count)

		for _, c := range k.centroids {
			buff = append(buff, math.EuclideanDistance(p.clr, c.clr))
		}

		k.pixels[i].centroid = math.SliceIndexMin(buff)
	}
}

func (k *KMeans) newCentroids() {
	type dist struct {
		index int
		dist  float64
	}
	buff := make([][]dist, k.count)
	for i := range buff {
		buff[i] = make([]dist, 0)
	}

	color := k.avgColor()
	for _, c := range color {
		fmt.Println(c)
	}

	for i, p := range k.pixels {
		d := dist{
			index: i,
			dist:  math.EuclideanDistance(p.clr, color[p.centroid]),
		}

		buff[p.centroid] = append(buff[p.centroid], d)
	}

	for i := 0; i < k.count; i++ {
		s := buff[i]
		if len(s) == 0 {
			continue
		}
		sort.Slice(s, func(k, l int) bool { return s[k].dist < s[l].dist })
		k.centroids[i] = k.pixels[s[0].index]
	}
}

func (k *KMeans) avgColor() []color.RGBA {
	result := make([]color.RGBA, 0, k.count)
	r := make([]float64, k.count)
	g := make([]float64, k.count)
	b := make([]float64, k.count)
	count := make([]float64, k.count)

	for _, k := range k.pixels {
		r[k.centroid] += float64(k.clr.R)
		g[k.centroid] += float64(k.clr.G)
		b[k.centroid] += float64(k.clr.B)
		count[k.centroid]++
	}

	for i := 0; i < k.count; i++ {
		var c color.RGBA
		c.R = uint8(r[i] / count[i])
		c.G = uint8(g[i] / count[i])
		c.B = uint8(b[i] / count[i])
		result = append(result, c)
	}

	return result
}

func (k *KMeans) image() image.Image {
	img := gil.NewImage(k.width, k.height)

	for _, p := range k.pixels {
		c := color.RGBA{
			R: k.centroids[p.centroid].clr.R,
			G: k.centroids[p.centroid].clr.R,
			B: k.centroids[p.centroid].clr.R,
			A: 255,
		}
		img.(*image.RGBA).Set(p.x, p.y, c)
	}

	return img
}

func imgToPixel(img image.Image) []pixel {
	pixels := make([]pixel, 0, img.Bounds().Max.X*img.Bounds().Max.Y)

	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			p := pixel{
				x:        x,
				y:        y,
				clr:      convert.RGBA32toRGBA8(img.At(x, y)),
				centroid: -1,
			}
			pixels = append(pixels, p)
		}
	}

	return pixels
}

func genCentroid(pixels []pixel, count int) []pixel {
	centroids := make([]pixel, 0, count)

	size := len(pixels)
	step := size / count

	index := 0
	for i := step; i < size; i += step {
		p := pixel{
			x:        pixels[i].x,
			y:        pixels[i].y,
			clr:      pixels[i].clr,
			centroid: index,
		}

		centroids = append(centroids, p)
		index++
	}

	return centroids
}
