package main

import (
	"fmt"
	"modtest2/retriever/mock"
	real2 "modtest2/retriever/real"
	"time"
)

type Retriever interface {
	Get(string) string
}

func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake imooc.com"}
	inspect(r)
	r = &real2.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	//Type assertion
	realRetriever := r.(*real2.Retriever)
	fmt.Println(realRetriever.TimeOut)
	//fmt.Printf(download(r))
}

func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
}
