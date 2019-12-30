package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("When %v is divided by 4, the reminder (aka modulus) is %v\n", i, i%4)
	}

	// exercise 6
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println(x)
	}

	// exercise 7

	y := "James Bond"
	if y == "Moneypenny" {
		fmt.Println(y)
	} else if y == "James Bond" {
		fmt.Println("BONDBONDBONDBOND", x)
	} else {
		fmt.Println("neither")
	}
}
