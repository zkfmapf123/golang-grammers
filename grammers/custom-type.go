package main

import "fmt"

// 1. type
var (
	floatVar32 float32 = 0.1 // less memory
	floatVar64 float64 = 0.2
	name       string  = "foo"
	intvar32   int32   = 1 // less memory
	intvar64   int64   = 48484
	intvar     int     = -1
	uintvar    uint    = 1                    // only x > 0
	uint32var  uint32  = 1000000000           // byte 2^11 (digit)
	uint64var  uint64  = 10000000000000000000 // byte 2^.. (digit)
	unit8var   uint8   = 0x01                 // byte 2^3 (digit)
	bytevar    byte    = 0x02
	runvar     rune    = 'a'
)

// 2. custom-type (strcut)
type Player struct {
	name        string
	attackPower float64
	health      int
}

// 3. simple func patter
func getHealth(player Player) int {
	return player.health
}

func getAttachPower(player Player) float64 {
	return player.attackPower
}

func getName(player Player) string {
	return player.name
}

// 4. Receiver
func (p Player) getReceiverHealth() int {
	return p.health
}

func (p Player) getReceiverAttachPower() float64 {
	return p.attackPower
}

func (p Player) getReceiverName() string {
	return p.name
}

func main() {
	p1 := Player{
		name:        "leedonggyu",
		attackPower: 30.3,
		health:      100,
	}

	fmt.Printf("health : %v\n", getHealth(p1))
	fmt.Printf("attachPower : %v\n", getAttachPower(p1))
	fmt.Printf("name : %v\n", getName(p1))

	// 5. map
	// emptyMap := map[string]int{} empry map
	users := map[string]int{
		"name": 30,
	}

	users["foo"] = 10
	users["bar"] = 11

	delete(users, "foo") // delete map

	for key, value := range users {
		fmt.Println(key, value)
	}

	age1, ok1 := users["foo"]       // int, true
	age2, ok2 := users["undefined"] // 0, false

	if ok2 {
		fmt.Println("undeined is ok")
	}

	fmt.Println(age1, ok1)
	fmt.Println(age2, ok2)

	fmt.Printf("%v\n", users)

	a := map[string]int{}     // 간단하게 map 사용
	b := make(map[string]int) // make는 heap space에 올라가는 명령어임 + 초기용량을 지정할수 있음
	fmt.Println(a, b)

	// 6. slice
	makeNum := make([]int, 5)
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(makeNum, len(makeNum), cap(makeNum))
	fmt.Println(numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 10, 20, 30, 40) // append

	fmt.Println(numbers)

	// 7. custom type
	type Weapon string

	var axe Weapon
	axe = "axe"

	getWeapon := func(w Weapon) string {
		return string(w)
	}

	fmt.Println(getWeapon(axe))

}
