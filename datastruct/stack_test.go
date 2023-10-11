package datastruct

import (
	"fmt"
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	fmt.Println(s.arr)
	fmt.Println(s.Pop())
	fmt.Println(s.arr)
}

func TestPop(t *testing.T) {
	s := NewStack[int]()
	s.Push(1)
	s.Push(2)
	fmt.Println(s, s.Length())
	v, err := s.Get()
	fmt.Println(v, err)
	v1, err := s.Pop()
	fmt.Println(v1, err)
	v2, err := s.Pop()
	fmt.Println(v2, err)
	v3, err := s.Pop()
	fmt.Println(v3, err)
}
