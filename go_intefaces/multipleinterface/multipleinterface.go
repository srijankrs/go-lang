package main

import (
	"fmt"
	"strconv"
)

type CarInterface1 interface {
	Color() string
}
type CarInterface2 interface {
	Speed() string
}

type Honda struct {
	ColorValue string
	SpeedValue int
}

func (honda Honda) Color() string {
	return "Color of Honda Car is " +honda.ColorValue
}

func (honda Honda) Speed() string {
	return "Speed of Honda Car is "+strconv.Itoa(honda.SpeedValue)
}

func main() {
	var honda = Honda{"Red",240}
	var carObj1 CarInterface1 = honda
	var carObj2 CarInterface2 = honda
	fmt.Println(carObj1.Color())
	fmt.Println(carObj2.Speed())

	fmt.Printf("Type- %T | ",carObj1)

	fmt.Printf("Type- %T",carObj2)
}