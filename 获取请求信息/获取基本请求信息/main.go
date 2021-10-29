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
	fmt.Fprintln(w, "您发送的请求地址：", r.URL.Path)
	//改行会将请求地址后的?后面的内容打印出来
	fmt.Fprintln(w, "您发送的请求地址后的查询字符串：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中所有信息：", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值：", r.Header.Get("Accept-Encoding"))
}
