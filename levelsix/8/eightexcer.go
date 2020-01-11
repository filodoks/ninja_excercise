package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f())

	// ********************************

	h := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}
	x := bar(h, []int{1, 2, 3, 4, 5})
	fmt.Println(x)
}

func foo() func() int {
	return func() int {
		return 42
	}
}

func bar(g func(xi []int) int, ii []int) int {
	n := g(ii)
	n++
	return n
}
