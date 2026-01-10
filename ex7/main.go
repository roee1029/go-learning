package main

import (
	"errors"
	"sync"
)

type SafeSlice struct {
	innerSlice []T
	mu sync.RWMutex
}

func (s *SafeSlice) NewSafeSlice(slice []T) *SafeSlice {
	return &SafeSlice{slice, sync.RWMutex{}}
}

func (s *SafeSlice) Append(val T){
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.innerSlice = append(s.innerSlice, val)
}

func (s *SafeSlice) Get(index int) (T, error){
	s.mu.Lock()
	defer s.mu.Unlock()
	if index < 0 || index >= len(s.innerSlice){
		return 0, errors.New("Index out of bound")
	}
	return s.innerSlice[index], nil
}

func (s *SafeSlice) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.innerSlice)
}

func main() {

}