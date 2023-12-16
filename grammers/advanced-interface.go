package main

import (
	"errors"
	"fmt"
)

// 1. advanced interface machanics
// 2. typed functions

type Storage interface {
	Get(int) (any, error)
	Put(int, any) error
}

// ///////////////////////// Foo Storage ///////////////////////////
type FooStorage struct {
	data map[int]any
}

func New() *FooStorage {
	d := map[int]any{}
	return &FooStorage{
		data: d,
	}
}

// PointReceiver를 사용하여 내부 데이터 접근 허용
func (f *FooStorage) Get(id int) (any, error) {

	v, ok := f.data[id]
	if ok {
		return v, nil
	}

	return nil, errors.New("Not Exists Value")
}

func (f *FooStorage) Put(id int, val any) error {
	f.data[id] = val
	return nil
}

type Server struct {
	store Storage
}

func main() {
	s := &Server{
		store: New(),
	}

	s.store.Put(1, "foo")
	v, _ := s.store.Get(1)
	fmt.Println(v)
}
