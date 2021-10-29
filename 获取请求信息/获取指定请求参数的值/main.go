package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

//快速获取请求参数
func Handler(w http.ResponseWriter, r *http.Request) {
	//根据请求参数名获取其值
	//r.FormValue同时获取表单和URL请求参数名的值 如果重名会获取表单的值
	//r.PostFormValue获取表单请求参数名的值
	fmt.Println("URL的请求参数user的值：", r.FormValue("password"))
	fmt.Println("表单的请求参数username的值：", r.PostFormValue("username"))
	fmt.Println("表单的请求参数password的值：", r.PostFormValue("password"))
	fmt.Fprintln(w, "操作成功，请查看控制台！ 来自：", r.URL.Path)
}
