package main

import "fmt"

func main() {
	// exercise 8
	// switch {
	// case false:
	// 	fmt.Println("should not print")
	// case true:
	// 	fmt.Println("should print")
	// }

	// exercise 9
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool")
	case "surfing":
		fmt.Println("go to hawaii")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}
