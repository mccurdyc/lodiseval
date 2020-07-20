package mapstore

import (
	"context"

	"github.com/mccurdyc/lodiseval/store"
)

func Factory(ctx context.Context, conf *store.Config) (store.Store, error) {
	return newMapStore(), nil
}

type Map struct {
	m map[interface{}]interface{}
}

func newMapStore() *Map {
	return &Map{
		m: make(map[interface{}]interface{}),
	}
}

func (s *Map) Set(k, v interface{}) error {
	s.m[k] = v

	return nil
}

func (s *Map) Get(k interface{}) (interface{}, error) {
	return s.m[k], nil
}

func (s *Map) GetState() (interface{}, error) {
	return s.m, nil
}
