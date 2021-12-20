package main

import "fmt"

func f1(a int) int {
	if a == 1 {
		return 1
	} else {
		return a * f1(a-1)
	}
}
func main() {
	fmt.Println(f1(5))
}
