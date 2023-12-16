package main

import "fmt"

// Comparable을 사용하여, map 내에서 일관되고 적절하게 동작하는것을 보장
// 주로 Map, Slice에서 사용
type CustomMap[T comparable, K any] struct {
	data map[T]K
}

func (m *CustomMap[T, K]) Insert(k T, v K) error {
	m.data[k] = v

	return nil
}

func NewCustomeMap[T comparable, V any]() *CustomMap[T, V] {

	return &CustomMap[T, V]{
		data: make(map[T]V),
	}
}

func foo[T any](val T) {
	fmt.Println(val)
}

func main() {

	strMap := NewCustomeMap[string, int]()
	strMap.Insert("foo", 10)
	fmt.Println(strMap.data)

	boolMap := NewCustomeMap[string, bool]()
	boolMap.Insert("foo", true)
	fmt.Println(boolMap)

	foo("hello world")
	foo(123)
	foo(true)
}
