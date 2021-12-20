package main

import "fmt"

/*func sum(a int, b int) (ret int) {
	ret = a + b
	return ret
}

func comp(a, b int) (max int) {
	if a > b {
		max = a
	} else {
		max = b
	}
	return
}*/

func f3(args ...int) {
	for _, v := range args {
		fmt.Println("v: ", v)
	}
}

func main() {
	f3(1, 5, 3, 2, 5, 1)
}
