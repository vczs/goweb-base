package main

import (
	"fmt"
	"net/http"
)

func main() {
	//创建一个多路复用器
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)      //绑定处理器
	http.ListenAndServe(":8080", mux) //开启路由器并传入创建的多路复用器
}

//创建处理器,处理器名可自定义
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!!", r.URL.Path)
}
