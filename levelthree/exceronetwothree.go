package main

import "fmt"

func main() {
	// for i := 0; i <= 10000; i++ {
	// 	fmt.Println(i)
	// }

	// excercise TWO:
	// for i := 65; i <= 90; i++ {
	// 	fmt.Println(i)
	// 	for j := 0; j < 3; j++ {
	// 		fmt.Printf("\t%#U\n", i)
	// 	}
	// }

	// excercise THREE:

	bd := 1985
	for bd <= 2017 {
		fmt.Println(bd)
		bd++
	}
}
