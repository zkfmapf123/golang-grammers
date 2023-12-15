package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	p       *Position
}

type SpecialEntity struct {
	e             *Entity
	specialEntity string
}

func main() {
	p := Position{
		x: 10,
		y: 20,
	}

	e := Entity{
		name:    "name",
		id:      "id",
		version: "version",
		p:       &p,
	}

	se := SpecialEntity{
		e:             &e,
		specialEntity: "specialEntity",
	}

	p.x = 1000
	p.y = 1000
	fmt.Printf("%+v\n", p, &p)
	fmt.Printf("%+v\n", e.p.x, e.p.y)
	fmt.Printf("%+v\n", se.e.p.x, se.e.p.y)

}
