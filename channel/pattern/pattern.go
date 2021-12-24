package channel

import (
	"fmt"
	"math/rand"
	"time"
)

//生成器
func msgGen(name string) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for true {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s: message %d", name, i)
			i++
		}
	}()
	return c
}

func fanIn(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for true {
			c <- <-c1
		}
	}()
	go func() {
		for true {
			c <- <-c2
		}
	}()
	return c
}

func MainMsgGen() {
	m1 := msgGen("service1")
	m2 := msgGen("service2")
	m := fanIn(m1, m2)
	for true {
		fmt.Println(<-m)
	}
}
