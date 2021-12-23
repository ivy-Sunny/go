package goroutine

import (
	"fmt"
	"time"
)

func MainGoroutine() {
	var a [10]int
	for i := 0; i < 100; i++ {
		go func() {
			for {
				fmt.Println("hello from goruntine", i)
			}
		}()
	}
	//i最后会变10，并跳出循环 race conditon
	time.Sleep(time.Second)

	fmt.Println(a)
}
