package test

import (
	"fmt"
	"math/rand"
	"time"
)

var Values = make(chan int)

func Send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send: %v\n", value)
	Values <- value
}
