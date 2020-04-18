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
type CarInterface3 interface {
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
	var car CarInterface1= Honda{"Red",240}
	carObj1 := car.(Honda)

	fmt.Println(carObj1.Color())
	fmt.Println(carObj1.Speed())


	//var car CarInterface1= Honda{"Red",240}
	//
	//carObj1, err := car.(CarInterface2)
	//fmt.Println("Value ",carObj1," | error ",err)
	//
	//carObj2, err := car.(CarInterface3)
	//fmt.Println("Value ",carObj2," | error ",err)
}

