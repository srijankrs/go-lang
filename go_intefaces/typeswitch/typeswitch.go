package main

import (
	"fmt"
)

func CustomPrint(i interface{}) { //empty interface
	fmt.Printf("Value - %v | Type - %T",i,i)
}

//func CustomPrint(i interface{}) { //empty interface
//	switch i.(type) {
//	case string:
//		fmt.Printf("Value - %v | Type is a String/Line",i)
//	default:
//		fmt.Printf("Value - %v | Type - %T",i,i)
//	}
//}


func main() {
	CustomPrint("Test")
	fmt.Println()
	CustomPrint(12.9)
}