package main

import (
	"fmt"
	"net/http"
)

func main() {
	//第一个参数处理请求的地址 第二个参数是处理器
	http.HandleFunc("/http", handler) //绑定处理器
	http.ListenAndServe(":8080", nil) //创建并开启路由器
}

//创建处理器,处理器名可自定义
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!!!", r.URL.Path)
}
