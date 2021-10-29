package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

//给客户端响应一个HTML页面
func Handler(w http.ResponseWriter, r *http.Request) {
	//设置响应报文头内容   Location：重定向地址
	w.Header().Set("Location", "https://www.baidu.com")
	//设置响应状态码给客户端  302:重定向状态码
	w.WriteHeader(302)
}
