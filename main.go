package main

import "github.com/mccurdyc/lodiseval/pkg/replica"

func main() {
	var (
		ch1    = make(chan interface{}, 5)
		store1 = newMapStore(map[interface{}]interface{}{})

		ch2    = make(chan interface{}, 5)
		store2 = newMapStore(map[interface{}]interface{}{})

		ch3    = make(chan interface{}, 5)
		store3 = newMapStore(map[interface{}]interface{}{})
	)

	r1 := replica.New("1", store1, ch1)
	r2 := replica.New("2", store2, ch2)
	r3 := replica.New("3", store3, ch3)

	go func() {
	}

	go func() {
	}

	go func() {
	}

}

// TODO: move to a mapstore.go file in package main or something.
type mapStore struct {
	m map[interface{}]interface{}
}

func newMapStore(m map[interface{}]interface{}) *mapStore {
	return &mapStore{
		m: m,
	}
}

func (s *mapStore) Set(k, v interface{}) {
	s.m[k] = v
}

func (s *mapStore) Get(k interface{}) interface{} {
	return s.m[k]
}

func (s *mapStore) GetState() interface{} {
	return s.m
}
