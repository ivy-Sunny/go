package test

import (
	"fmt"
	"sync"
)

var wp sync.WaitGroup

func showMSG(i int) {
	// defer wp.Add(-1) 等价于 defer wp.Wait()
	defer wp.Wait()
	fmt.Printf("i: %v\n", i)
}

func MainWaitgroup() {
	for i := 0; i < 10; i++ {
		//启动一个协程来执行
		go showMSG(i)
		wp.Add(1)
	}
	wp.Wait()
	//主协程
	fmt.Println("end...")
}
