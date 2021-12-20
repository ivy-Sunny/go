package main

import "fmt"

var i int = initVar()

func init() {
	fmt.Println("init...")
}
func initVar() int {
	fmt.Println("initVar...")
	return 100
}
func main() {
	fmt.Println("main...")
}
