package datastruct

import (
	"fmt"
	"testing"
)

// .

func TestQueue(t *testing.T) {
	q := NewQueue[int]()
	q.Push(1)
	q.Push(2)
	fmt.Println(q.arr)
	fmt.Println(q.Pop())
	fmt.Println(q.arr)
	v, err := q.Get()
	fmt.Println(v, err)
	fmt.Println(q.Length())
}
