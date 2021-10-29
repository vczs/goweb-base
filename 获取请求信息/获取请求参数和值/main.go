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
	//解析表单 在调用r.From之前必须执行该操作
	r.ParseForm()
	//打印请求参数 (进行下面的操作时要注释掉Body read操作 否则下面的操作不能获取到请求信息内容)
	fmt.Println("URL和表单的请求信息：", r.Form)                 //同时获取表单和URL地址里的请求参数和值 如果有相同的请求参数 表单里的在前URL里的在后
	fmt.Println("enctype的属性是默认属性时表单的请求信息：", r.PostForm) //只获取表单里的请求参数和值

	//PostForm只支持enctype为默认属性  如果enctype的属性是multipart/from-data时 就需要使用MultipartFrom来获取表单里的请求参数和值
	r.ParseMultipartForm(4) //调用r.MultipartForm时需要用r.ParseMultipartForm()来解析表单
	fmt.Println("enctype的属性是multipart/from-data时表单的请求信息：", r.MultipartForm)
	fmt.Fprintln(w, "操作成功，请查看控制台！ 来自：", r.URL.Path)
}
