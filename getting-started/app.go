package main

import (
	"fmt"
)

func find(x int) (int,bool) {
	sq := x*x
	if sq > 100 {
		return sq,true
	}
	return sq,false
}

type testType interface {

}

func main() {

	var num = 12
	sq,isGreater := find(num)


	if isGreater {
		fmt.Print("Square is ",sq)
	}

}
