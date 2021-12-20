package main

import "fmt"

type Person struct {
	id   int
	name string
}

func showPerson(per *Person) {
	per.id = 101
	per.name = "kite"
	fmt.Printf("per: %v\n", per)
}

func main() {
	tom := Person{
		id:   100,
		name: "tom",
	}
	fmt.Printf("tom: %v\n", tom)
	fmt.Printf("--------\n")
	showPerson(&tom)
	fmt.Printf("tom: %v\n", tom)
}
