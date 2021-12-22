package test

import "fmt"

func showMSG(i int) {
	fmt.Printf("i: %v\n", i)
}

func MainWaitgroup() {
	for i := 0; i < 10; i++ {
		//启动一个协程来执行
		go showMSG(i)
	}

	//主协程
	fmt.Println("end...")
}
