package kmeans

import (
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
	imgLabel  image.Image
}

type pixel struct {
	x        int
	y        int
	RGBA     color.RGBA
	centroid int // index centroid
}

type dist struct {
	index int
	dist  float64
}

func NewKMeans(img image.Image, count int, imgLabel image.Image) *KMeans {
	pixels := imgToPixel(convert.Grayscale(img))

	k := new(KMeans)
	k.width = img.Bounds().Max.X
	k.height = img.Bounds().Max.Y
	k.centroids = genCentroid(pixels, count)
	k.imgLabel = imgLabel
	k.count = count
	k.pixels = pixels

	return k
}

func (k *KMeans) Predict(isDraw bool) (image.Image, error) {
	for i := 0; i < 32; i++ {
		k.evaluateCentroids()
		k.newCentroids()
	}

	if isDraw {
		return k.imageColor(), nil
	}
	return k.image(), nil
}

func (k *KMeans) evaluateCentroids() {
	for i, p := range k.pixels {
		buff := make([]float64, 0, k.count)

		for _, c := range k.centroids {
			buff = append(buff, math.EuclideanDistance(p.RGBA, c.RGBA))
		}

		k.pixels[i].centroid = math.SliceIndexMin(buff)
	}
}

func (k *KMeans) newCentroids() {
	buff := make([][]dist, k.count)
	for i := range buff {
		buff[i] = make([]dist, 0)
	}

	color := k.avgDistance()

	for i, p := range k.pixels {
		d := dist{
			index: i,
			dist:  math.EuclideanDistance(p.RGBA, color[p.centroid]),
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

	for _, p := range k.pixels {
		r[p.centroid] += float64(p.RGBA.R)
		g[p.centroid] += float64(p.RGBA.G)
		b[p.centroid] += float64(p.RGBA.B)
		count[p.centroid]++
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

func (k *KMeans) avgDistance() []color.RGBA {
	result := make([]color.RGBA, 0, k.count)
	x := make([]int, k.count)
	y := make([]int, k.count)
	count := make([]int, k.count)

	for _, p := range k.pixels {
		x[p.centroid] += p.x
		y[p.centroid] += p.y
		count[p.centroid]++
	}

	for i := 0; i < k.count; i++ {
		var c color.RGBA

		if x[i] == 0 || y[i] == 0 {
			c = k.centroids[i].RGBA
			result = append(result, c)
			continue
		}

		x0 := x[i] / count[i]
		y0 := y[i] / count[i]

		c = k.pixels[(x0*y0)+y0].RGBA
		result = append(result, c)
	}

	return result
}

func (k *KMeans) image() image.Image {
	img := gil.NewImage(k.width, k.height)

	for _, p := range k.pixels {
		c := color.RGBA{
			R: k.centroids[p.centroid].RGBA.R,
			G: k.centroids[p.centroid].RGBA.G,
			B: k.centroids[p.centroid].RGBA.B,
			A: 255,
		}
		img.(*image.RGBA).Set(p.x, p.y, c)
	}

	return img
}

func (k *KMeans) imageColor() image.Image {
	img := gil.NewImage(k.width, k.height)

	for _, p := range k.pixels {
		x := k.centroids[p.centroid].x
		y := k.centroids[p.centroid].y
		img.(*image.RGBA).Set(p.x, p.y, k.imgLabel.At(x, y))
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
				RGBA:     convert.RGBA32toRGBA8(img.At(x, y)),
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
	for i := 0; index < count; i += step {
		p := pixel{
			x:        pixels[i].x,
			y:        pixels[i].y,
			RGBA:     pixels[i].RGBA,
			centroid: index,
		}

		centroids = append(centroids, p)
		index++
	}

	return centroids
}
