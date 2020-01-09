package main

import "fmt"

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1984, "Big Brother is watching"
}

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}
