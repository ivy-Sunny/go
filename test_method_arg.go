package main

import "fmt"

func main() {
	type Person struct {
		name string
	}
	p1 := Person{"tom"}
	p2 := &Person{"tom"}
	fmt.Printf("p1: %T\n", p1)
	fmt.Printf("p2: %T\n", p2)

}
