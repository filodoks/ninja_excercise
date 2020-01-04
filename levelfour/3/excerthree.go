package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
	fmt.Println(x)

	y := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y = append(y, 52)
	fmt.Println(y)
	y = append(y, 53, 54, 55)
	fmt.Println(y)

	z := []int{56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println(y)

	// excercize 5
	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(a)
	a = append(a[:3], a[6:]...)
	fmt.Println(a)
}
