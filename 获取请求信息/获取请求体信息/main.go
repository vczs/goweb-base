package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
func Handler(w http.ResponseWriter, r *http.Request) {
	//获取请求体内容的长度
	lens := r.ContentLength
	//创建一个切片来接收请求体内容
	slice := make([]byte, lens)
	//将请求体内容读取到切片中并打印
	r.Body.Read(slice)
	fmt.Println(string(slice))
	fmt.Fprintln(w, "操作成功，请查看控制台！ 来自：", r.URL.Path)
}
