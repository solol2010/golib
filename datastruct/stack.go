package datastruct

import (
	"errors"
	"sync"
)

type stack[T any] struct {
	sync.RWMutex
	arr []T
}

func NewStack[T any]() *stack[T] {
	return &stack[T]{}
}

func (s *stack[T]) Push(e T) {
	s.Lock()
	defer s.Unlock()
	s.arr = append(s.arr, e)
}

func (s *stack[T]) Pop() (T, error) {
	s.Lock()
	defer s.Unlock()
	var e T
	if len(s.arr) <= 0 {
		return e, errors.New("empty stack")
	}
	e = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return e, nil
}

func (s *stack[T]) Get() (T, error) {
	if len(s.arr) <= 0 {
		var e T
		return e, errors.New("empty stack")
	} else {
		return (s.arr)[len(s.arr)-1], nil
	}
}

func (s *stack[T]) Length() int {
	return len(s.arr)
}
