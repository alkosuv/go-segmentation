package convert

import "image/color"

func RGBA32toRGBA8(pixel color.Color) color.RGBA {
	r, g, b, a := pixel.RGBA()

	var c color.RGBA
	c.R = uint8(r >> 8)
	c.G = uint8(g >> 8)
	c.B = uint8(b >> 8)
	c.A = uint8(a >> 8)

	return c
}
