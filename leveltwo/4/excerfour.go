package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)

	b := a << 1
	fmt.Printf("%d\t%b\t%#x\n", b, b, b)

	c := `here is something
	a
	sdasd
	daff
	aggggggg
	another thing`

	fmt.Println(c)
}