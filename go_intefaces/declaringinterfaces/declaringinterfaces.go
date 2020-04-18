package main

import (
	"fmt"
)

type Car interface {
	Color() string
	Speed() string
}

func main() {
	var car Car
	fmt.Println("Value-",car)
	fmt.Printf("Type- %T",car)
}