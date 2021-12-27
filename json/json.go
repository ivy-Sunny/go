package json

import (
	"encoding/json"
	"fmt"
)

type Order struct { //tag
	ID         string  `json:"id"`
	Name       string  `json:"name,omitempty"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}
type OrderItem struct {
}

func MainJson() {
	o := Order{
		ID: "1234",
		//Name:       "learn go",
		Quantity:   3,
		TotalPrice: 30,
	}
	b, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
	fmt.Printf("%+v\n", o)

}
