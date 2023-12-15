package main

import "fmt"

// ////////////////////////////// Interface ////////////////////////
type Storer interface {
	Retrieve() ([]int, error)
	Put(...int) (bool, error)
	GetData() []int
}

// ///////////////////////////// MongoDBStore ///////////////////////////////
// MongoDBStore의 함수들을 Storer의 Interface를 구현한다면 사용이 가능
type MongoDBStore struct {
	data []int
}

func New() MongoDBStore {
	return MongoDBStore{
		data: []int{},
	}
}

func (m MongoDBStore) Retrieve() ([]int, error) {
	m.data = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	return m.data, nil
}

func (m MongoDBStore) Put(numbers ...int) (bool, error) {
	m.data = append(m.data, numbers...)
	fmt.Println(m.data)
	return true, nil
}

func (m MongoDBStore) GetData() []int {
	return m.data
}

// ///////////////////////////// ApiServer ///////////////////////////////
type ApiServer struct {
	numberStore Storer
}

func main() {
	m := New()
	as := ApiServer{
		numberStore: m,
	}

	data, _ := as.numberStore.Retrieve()
	isAppend, _ := as.numberStore.Put(100, 200, 300, 400, 500)

	if isAppend {
		fmt.Println("init : ", data)
		fmt.Println("append : ", isAppend)
	}

}
