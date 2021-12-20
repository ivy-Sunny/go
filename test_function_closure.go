package main

import "fmt"

func cal(base int) (func(int) int, func(int) int) {
	add := func(a int) int {
		base += a
		return base
	}
	sub := func(a int) int {
		base -= a
		return base
	}
	return add, sub
}

func add() func(y int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	f := add()
	r := f(10)
	fmt.Println(r)
	r = f(20)
	fmt.Println(r)
	r = f(30)
	fmt.Println(r)

	add, sub := cal(100)
	r = add(100)
	fmt.Println(r)
	r = sub(50)
	fmt.Println(r)
}
