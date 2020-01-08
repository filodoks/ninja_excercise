package main

import "fmt"

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		"James",

		map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		[]string{
			"Martini",
			"Water",
		},
	}
	fmt.Println(s)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}
