package socket

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func MainJson() {
	indexFunc := func(response HttpResponse, request HttpRequest) {
		io.Copy(os.Stdout, request.Body)
		fmt.Println()
		decoder := json.NewDecoder(request.Body)
		var info map[string]interface{}
		err := decoder.Decode(&info)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(info)
	}
	http.HandleFunc("/", indexFunc)
	/*
		3.启动web服务
		http.ListenAndServer(addr,nil)
	*/
	http.ListenAndServe(":9999", nil)
}

func MainCookie() {
	indexFunc := func(response HttpResponse, request HttpRequest) {
		//读取cookie
		println(request.Header.Get("Cookie"))
		counter := 1
		counterCookie := http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter),
			HttpOnly: true,
		}
		http.SetCookie(response, &counterCookie)
	}
	http.HandleFunc("/", indexFunc)
	/*
		3.启动web服务
		http.ListenAndServer(addr,nil)
	*/
	http.ListenAndServe(":9999", nil)
}

func MainRedirect() {
	indexFunc := func(response HttpResponse, request HttpRequest) {
		http.Redirect(response, request, "/login", 302)
		fmt.Fprint(response, "首页")
	}
	loginFunc := func(response HttpResponse, request HttpRequest) {
		fmt.Fprint(response, "登录页面")
	}
	http.HandleFunc("/", indexFunc)
	http.HandleFunc("/login", loginFunc)

	/*
		3.启动web服务
		http.ListenAndServer(addr,nil)
	*/
	http.ListenAndServe(":9999", nil)
}
