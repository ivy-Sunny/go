package socket

import (
	"learngo/socket/service"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func MainRpcServer() {
	addr := ":9999"
	rpc.Register(&service.Calculator{})
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("[-]listen on: %s\n", addr)

	for {
		accept, err := listener.Accept()
		if err != nil {
			log.Printf("[-]error client: %s\n", err.Error())
			continue
		}
		log.Printf("[-]client connected: %s\n", accept.RemoteAddr())
		go jsonrpc.ServeConn(accept)
	}
}
