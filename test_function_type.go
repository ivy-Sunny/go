package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
func main() {
	type f1 func(int, int) int
	var ff f1
	ff = sum
	r := ff(2, 5)
	fmt.Println(r)
	ff = max
	r = ff(9, 1)
	fmt.Println(r)

}
