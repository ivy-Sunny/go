package main

import "fmt"

func main() {
	type Dog struct {
		name  string
		age   int
		color string
	}

	type Person struct {
		dog  Dog
		name string
		age  int
	}

	dog := Dog{
		name:  "小黑",
		age:   2,
		color: "black",
	}
	per := Person{
		dog:  dog,
		name: "tom",
		age:  20,
	}
	fmt.Printf("per: %v\n", per)
}