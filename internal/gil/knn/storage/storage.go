package storage

import "image/color"

type Storage struct {
	RGBA  color.RGBA
	Index int
}

type Storages []Storage

func (s *Storages) Add(rgba color.RGBA, index int) {
	*s = append(*s, Storage{rgba, index})
}
