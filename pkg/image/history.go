package image

import (
	"image/draw"
	"sync"
)

type Storage struct {
	store *sync.Map
}

func NewStorage() *Storage {
	return &Storage{
		store: &sync.Map{},
	}
}

func (s *Storage) Put(k string, v draw.Image) {
	s.store.Store(k, v)
}

func (s *Storage) Get(k string) (draw.Image, bool) {
	v, ok := s.store.Load(k)
	if !ok {
		return nil, false
	}
	return v.(draw.Image), true
}
