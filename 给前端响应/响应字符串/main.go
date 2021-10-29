package main

import "net/http"

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

//给客户端响应一个字符串
//http.ResponseWriter是个接口,被HTTP处理器用于给客户端构造HTTP回应
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("字符串回应：您的请求已收到！")) //w的Write()方法对客户端进行回应
}
