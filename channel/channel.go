package channel

import (
	"fmt"
	"time"
)

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func worker(id int, c chan int) {
	for {
		fmt.Printf("%v\n", time.Now())
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d Received %c \n", id, n)
	}
}

func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	time.Sleep(time.Second)
	c <- 'b'
	time.Sleep(time.Second)
	c <- 'c'
	time.Sleep(time.Second)
	close(c)
}
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'd'
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}
func chanDemo() {
	var channels [10]chan<- int // c == nil
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
}
func MainChannel() {
	//chanDemo()
	bufferedChannel()
	//channelClose()
	time.Sleep(time.Second * 100)
}
