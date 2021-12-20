package main

import "fmt"

func main() {
	var tom Person
	tom = Person{
		id:    101,
		name:  "tom",
		age:   20,
		email: "tom@gmail.com",
	}
	fmt.Println(tom)

}

type Customer struct {
	id, age     int
	name, email string
}
type Person struct {
	id    int
	name  string
	age   int
	email string
}
