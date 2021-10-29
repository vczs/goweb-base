package main

import (
	"fmt"
	"net/http"
)

//创建一个结构体实例
type MyHandler struct{}

//注意：方法名必须是ServeHTTP
//结构体实例的方法实现ServeHTTP()即代表MyHandler实例实现Handler接口  即创建处理器
func (myHandler *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "自己创建的处理器！！！")
}

func main() {
	myHandler := MyHandler{}
	http.Handle("/myHandler", &myHandler) //绑定处理器
	http.ListenAndServe(":8080", nil)     //创建并开启路由器
}
