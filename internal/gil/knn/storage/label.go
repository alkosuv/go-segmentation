package storage

import (
	"fmt"
	"image/color"
	"strconv"
	"strings"
)

type Label struct {
	RGBA  color.RGBA
	Index int
}

func (l Label) String() string {
	return fmt.Sprintf("%v;%d;\n", l.RGBA, l.Index)
}

func NewLabel(value string) (*Label, error) {
	buff := strings.Split(value, ";")

	strRGBA := buff[0]
	strIndex := buff[1]

	strRGBA = strRGBA[1 : len(strRGBA)-1]
	arrRGBA := strings.Split(strRGBA, " ")

	index, err := strconv.Atoi(strIndex)
	if err != nil {
		return nil, err
	}

	r, err := strconv.ParseUint(arrRGBA[0], 10, 64)
	if err != nil {
		return nil, err
	}

	g, err := strconv.ParseUint(arrRGBA[1], 10, 64)
	if err != nil {
		return nil, err
	}

	b, err := strconv.ParseUint(arrRGBA[2], 10, 64)
	if err != nil {
		return nil, err
	}

	a, err := strconv.ParseUint(arrRGBA[3], 10, 64)
	if err != nil {
		return nil, err
	}

	l := &Label{
		RGBA: color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		},
		Index: index,
	}

	return l, nil
}
