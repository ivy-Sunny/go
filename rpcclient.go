package main

import (
	"learngo/socket/data"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	addr := "127.0.0.1:9999"
	conn, err := jsonrpc.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	request := &data.CalculatorRequest{}
	response := &data.CalculatorRespeonse{}
	conn.Call("Calculator.Add", request, response)
}
