package math

import (
	"image/color"
	"math"
)

// Функция рассчитывает расстояние Евклида
func EuclideanDistance(point1, point2 color.RGBA) float64 {
	r := float64(point1.R - point2.R)
	g := float64(point1.G - point2.G)
	b := float64(point1.B - point2.B)

	return math.Sqrt(math.Pow(r, 2) + math.Pow(g, 2) + math.Pow(b, 2))
}
