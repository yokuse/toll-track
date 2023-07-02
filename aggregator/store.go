package main

import (
	"fmt"
	"toll-calculator/types"
)

// put in some db
type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (s *MemoryStore) Insert(d types.Distance) error {
	s.data[d.OBUID] += d.Value
	return nil
}

func (s *MemoryStore) Read(obuId int) (float64, error) {
	dist, ok := s.data[obuId]
	if !ok {
		return 0.0, fmt.Errorf("could not find distance for obu if %d", obuId)
	}
	return dist, nil	
}