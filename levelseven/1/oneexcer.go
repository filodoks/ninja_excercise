package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
	//p := 42
	//fmt.Println(p)
	//fmt.Println(&p)
}

func changeMe(p *person) {
	p.name = "Miss Moneypenny"
	// (*p).name = "Miss Moneyp"
}
