package main

import "fmt"

type Color int

func (c Color) String() string {
	switch c {
	case Colorblue:
		return "blue"
	case ColorYellow:
		return "yellow"
	case ColorBlack:
		return "black"
	case ColorPink:
		return "pink"
	default:
		panic("program exit")
	}
}

const (
	Colorblue Color = iota << 1
	ColorBlack
	ColorYellow
	ColorPink
)

func main() {

	fmt.Println("Color is Blue : ", Colorblue.String())
}
