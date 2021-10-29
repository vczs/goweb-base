package main

import (
	"fmt"
	"net/http"
	"time"
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
	//创建server并详细配置 创建路由
	server := http.Server{
		Addr:        ":8080",         //端口号
		Handler:     &myHandler,      //绑定处理器
		ReadTimeout: time.Second * 3, //最大等待时间
	}
	server.ListenAndServe() //开启路由
}
