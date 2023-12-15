package main

import "fmt"

// Weapon Type
// axe
// sward
// wooden stick
// Knife
type WeaponType int

const (
	Axe WeaponType = iota // increment default 0
	Sward
	WoodenStick
	Knife
)

func getDamage(w WeaponType) int {
	switch w {
	case Axe:
		return 100
	case Sward:
		return 10
	case WoodenStick:
		return 1
	case Knife:
		return 50
	default:
		panic(fmt.Sprintf("%v program stop", w))
	}
}

// better func
func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "Axe"

	case Sward:
		return "Sward"

	case WoodenStick:
		return "WoodenStick"

	case Knife:
		return "Knife"

	default:
		panic("Program stop")
	}
}

func main() {

	fmt.Println(Axe, Axe.String(), getDamage(Axe))
	fmt.Println(WoodenStick, WoodenStick.String(), getDamage(WoodenStick))
	fmt.Println(Sward, Sward.String(), getDamage(Sward))
	fmt.Println(Knife, Knife.String(), getDamage(Knife))

	fmt.Println("Unkdefined", getDamage(10000)) // panic

}
