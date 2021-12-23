package channel

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(id, w)
	return w
}

func doWorker(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d Received %c \n", id, n)
		w.wg.Done()
	}
}

func channelClose(wg *sync.WaitGroup) {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(0, w)
	//c <- 'a'
	//time.Sleep(time.Second)
	//c <- 'b'
	//time.Sleep(time.Second)
	//c <- 'c'
	//time.Sleep(time.Second)
	//close(c)
}
func bufferedChannel(wg *sync.WaitGroup) {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWorker(0, w)
	//c <- 'd'
	//c <- 'a'
	//c <- 'b'
	//c <- 'c'
	//c <- 'd'
}
func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(20)
	var workers [10]worker // c == nil
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-workers[i].done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-workers[i].done
	}
	wg.Wait()
	// wait for all of them
	/*	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}*/
}
func MainDone() {
	chanDemo()
}
