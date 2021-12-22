package main

import "fmt"

type Pet interface {
	eat(string) string
}
type Dog struct {
	name string
}

func (dog *Dog) eat(food string) string {
	dog.name = "花花……"
	return dog.name + "吃的" + food
}

func main() {
	var pet Pet = &Dog{
		name: "余晖",
	}
	s := pet.eat("火鸡")
	fmt.Println(s)
	fmt.Println(pet)
}
