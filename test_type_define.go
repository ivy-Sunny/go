package main

import "fmt"

func main() {
	type MyInt1 int
	var i MyInt1
	i = 100
	fmt.Printf("i: \t%T,%v\n", i, i)

	type MyInt2 = int
	var j MyInt2
	j = 100
	fmt.Printf("j: \t%T,%v\n", j, j)

}
