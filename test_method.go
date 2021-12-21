package main

import "fmt"

type Person struct {
	name string
}

// 接收者
func (per Person) eat() {
	fmt.Println(per.name, "eat...")
}

type Customer struct {
	name string
}

func (customer Customer) login(name, pwd string) bool {
	fmt.Printf("customer.name: %v\n", customer.name)
	if name == customer.name && pwd == "123" {
		return true
	}
	return false
}

func main() {
	var per = Person{"tom"}
	per.eat()

	cus := Customer{"tom"}

	login := cus.login("tom", "123")
	fmt.Printf("%v", login)
}
