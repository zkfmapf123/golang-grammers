package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

// extends Entity
// 같은 값을 가지고있을 경우 => 메모리 낭비 => Pointer로 공유
type SpecialEntity struct {
	Entity       // entity *Entity 도 가능
	specialField float64
}

// +v  => key, value 다같이 나옴
// v => value만 나옴
func main() {
	// Position
	p := &Position{
		x: 10,
		y: 10,
	}

	// entity
	e := &Entity{
		name:     "my entity",
		id:       "1",
		version:  "1.0.0",
		Position: *p,
	}

	// specialEntity
	se := &SpecialEntity{
		specialField: 88.88,
		Entity:       *e,
	}

	p.x = 200
	p.y = 200

	fmt.Println("%+v\n", p) // 값 바뀜
	fmt.Println("%+v\n", e) // 값 안바뀜 (pointer로 주고받은게 아니니까...)
	fmt.Printf("%+v\n", se) // 값 안바뀜 (pointer로 주고받은게 아니니까...)
}
