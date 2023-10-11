package datastruct

import (
	"errors"
	"sync"
)

// .

type queue[T any] struct {
	sync.RWMutex
	arr []T
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{}
}

func (q *queue[T]) Push(e T) {
	q.Lock()
	defer q.Unlock()
	q.arr = append(q.arr, e)
}

func (q *queue[T]) Pop() (T, error) {
	q.Lock()
	defer q.Unlock()
	var e T
	if len(q.arr) <= 0 {
		return e, errors.New("empty queue")
	}
	e = q.arr[0]
	q.arr = q.arr[1:len(q.arr)]
	return e, nil
}

func (q *queue[T]) Get() (T, error) {
	if len(q.arr) <= 0 {
		var e T
		return e, errors.New("empty queue")
	} else {
		return (q.arr)[0], nil
	}
}

func (q *queue[T]) Length() int {
	return len(q.arr)
}
