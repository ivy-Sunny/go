package channel

import (
	"fmt"
	"math/rand"
	"time"
)

//生成器
func msgGen(name string, done chan struct{}) chan string {
	c := make(chan string)
	go func() {
		i := 0
		for true {
			select {
			case <-time.After(time.Duration(rand.Intn(2000)) * time.Millisecond):
				c <- fmt.Sprintf("service %s: message %d", name, i)
			case <-done:
				fmt.Println("cleaning up")
				time.Sleep(time.Second * 5)
				fmt.Println("cleaning done")
				done <- struct{}{}
				return
			}
			i++
		}
	}()
	return c
}

/*

 */
func fanIn(chs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range chs {
		go func(in chan string) {
			for {
				c <- <-in
			}
		}(ch)
	}
	return c
}

func fanInBySelect(c1, c2 chan string) chan string {
	c := make(chan string)
	go func() {
		for true {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

/*
非阻塞等待
*/
func nonBlockingWait(c chan string) (string, bool) {
	select {
	case m := <-c:
		return m, true
	default:
		return "", false
	}
}

/*
超时机制
*/
func timeoutWait(c chan string, timeout time.Duration) (string, bool) {
	select {
	case m := <-c:
		return m, true
	case <-time.After(timeout):
		return "", false
	}
}

func MainMsgGen() {
	done := make(chan struct{})
	m1 := msgGen("service1", done)
	/*	m2 := msgGen("service2")
		m := fanIn(m1, m2)
		m := fanInBySelect(m1, m2)
		for true {
			fmt.Println(<-m1)
			if m, ok := nonBlockingWait(m2); ok {
				fmt.Println(m)
			} else {
				fmt.Println("no message from service2")
			}
		}*/
	for i := 0; i < 5; i++ {
		if m, ok := timeoutWait(m1, time.Second); ok {
			fmt.Println(m)
		} else {
			fmt.Println("timeout")
		}
	}
	done <- struct{}{}
	<-done
}
