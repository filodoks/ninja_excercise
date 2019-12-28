package main

import (
	"fmt"
)

// var x int = 42
// var y string = "James Bond"
// var z bool = true

type hotdog int

var a hotdog
var y int

func main() {
	// s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	// fmt.Println(s)

	fmt.Println(a)
	fmt.Printf("%T\n", a)
	a = 42
	fmt.Println(a)
	y = int(a)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
