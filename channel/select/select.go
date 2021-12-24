package channel

import "fmt"

func MainSelect() {
	var c1, c2 chan int // c1 and c2 is nil
	/*	n1 := <-c1
		n2 := <-c2*/

	select {
	case n := <-c1:
		fmt.Println("Received from c1: ", n)
	case n := <-c2:
		fmt.Println("Received from c2: ", n)
	default:
		fmt.Println("No value received")
	}

}
