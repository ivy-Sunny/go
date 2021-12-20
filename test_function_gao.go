package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello,", name)
}
func test(name string, f func(string)) {
	f(name)
}

func add(a int, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}

func cal(operator string) func(int, int) int {
	switch operator {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}
func main() {
	test("tom", sayHello)

	ff := cal("+")
	fmt.Println(ff(1, 6))

	r := func(a, b int) int {
		return a + b
	}(2, 5)

	fmt.Println(r)
}
