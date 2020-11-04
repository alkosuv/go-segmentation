package label

import "image/color"

var Labels = map[color.RGBA]int{
	color.RGBA{0, 0, 255, 255}:     0,
	color.RGBA{255, 0, 0, 255}:     1,
	color.RGBA{255, 255, 0, 255}:   2,
	color.RGBA{0, 255, 0, 255}:     3,
	color.RGBA{255, 0, 255, 255}:   4,
	color.RGBA{0, 255, 255, 255}:   5,
	color.RGBA{255, 0, 153, 255}:   6,
	color.RGBA{153, 0, 255, 255}:   7,
	color.RGBA{0, 153, 255, 255}:   8,
	color.RGBA{153, 255, 0, 255}:   9,
	color.RGBA{255, 153, 0, 255}:   10,
	color.RGBA{0, 255, 153, 255}:   11,
	color.RGBA{0, 153, 153, 255}:   12,
	color.RGBA{0, 0, 0, 255}:       13,
	color.RGBA{0, 0, 153, 255}:     14,
	color.RGBA{255, 255, 153, 255}: 15,
}

var Color = map[int]color.RGBA{
	0:  color.RGBA{0, 0, 255, 255},
	1:  color.RGBA{255, 0, 0, 255},
	2:  color.RGBA{255, 255, 0, 255},
	3:  color.RGBA{0, 255, 0, 255},
	4:  color.RGBA{255, 0, 255, 255},
	5:  color.RGBA{0, 255, 255, 255},
	6:  color.RGBA{255, 0, 153, 255},
	7:  color.RGBA{153, 0, 255, 255},
	8:  color.RGBA{0, 153, 255, 255},
	9:  color.RGBA{153, 255, 0, 255},
	10: color.RGBA{255, 153, 0, 255},
	11: color.RGBA{0, 255, 153, 255},
	12: color.RGBA{0, 153, 153, 255},
	13: color.RGBA{0, 0, 0, 255},
	14: color.RGBA{0, 0, 153, 255},
	15: color.RGBA{255, 255, 153, 255},
}
