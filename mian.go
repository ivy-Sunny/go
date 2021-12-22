package main

import (
	"fmt"
	"pro2/test"
)

func main() {
	/*	go test.ResponseSize("https://baidu.com")
		go test.ResponseSize("https://jd.com")
		go test.ResponseSize("https://duoke360.com")

		time.Sleep(time.Second * 10)*/
	defer close(test.Values)
	go test.Send()
	fmt.Println("wait...")
	value := <-test.Values
	fmt.Printf("receive: %v\n", value)
	fmt.Println("end...")
}
