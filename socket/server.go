package socket

import (
	"fmt"
	"io/ioutil"
	"learngo/utils"
	"log"
	"net/http"
	"time"
)

type HttpResponse = http.ResponseWriter
type HttpRequest = *http.Request

//处理器
//HandleFunc interface 自定义类型(结构体)
type TimeHandle struct {
}
type IndexHandle struct {
}

func (h *IndexHandle) ServeHTTP(response HttpResponse, request HttpRequest) {
	ctx, err := ioutil.ReadFile("readme.md")
	if err == nil {
		response.Write(ctx)
	} else {
		log.Fatal(err)
		fmt.Fprint(response, "欢迎")
	}
}
func (h *TimeHandle) ServeHTTP(response HttpResponse, request HttpRequest) {
	fmt.Println(request)
	now := time.Now().Format(utils.BEIJING)
	//io.WriteString(response, now)
	fmt.Fprint(response, now)
}
func MainHandle() {
	//1.定义处理器/处理器函数
	//2.绑定URL 处理器/处理器函数关系
	//3.启动web服务
	http.Handle("/time", &TimeHandle{})
	http.Handle("/", &IndexHandle{})
	http.ListenAndServe(":9998", nil)
}

func MainHandleFunc() {
	//1.定义处理器/处理器函数
	//2.绑定URL 处理器/处理器函数关系
	//3.启动web服务

	/*
		1.处理器函数
		http.ResponeseWriter, *http.Request
	*/
	timeFunc := func(response HttpResponse, request HttpRequest) {
		fmt.Println(request)
		now := time.Now().Format(utils.BEIJING)
		//io.WriteString(response, now)
		fmt.Fprint(response, now)
	}
	indexFunc := func(response HttpResponse, request HttpRequest) {
		//解析参数
		request.ParseForm()
		form := request.Form
		fmt.Println(form)
		ctx, err := ioutil.ReadFile("readme.md")
		if err == nil {
			response.Write(ctx)
		} else {
			log.Fatal(err)
			fmt.Fprint(response, "欢迎")
		}
	}
	/*
		2.绑定URL关系
		http.HandleFunc(path,处理器函数)
	*/
	http.HandleFunc("/time/", timeFunc)
	http.HandleFunc("/", indexFunc)
	/*
		3.启动web服务
		http.ListenAndServer(addr,nil)
	*/
	http.ListenAndServe(":9999", nil)
}
