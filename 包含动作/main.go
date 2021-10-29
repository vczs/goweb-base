package main

import (
	"html/template"
	"net/http"
	"os"
)

//给客户端响应多个模板文件
func handler(w http.ResponseWriter, r *http.Request) {
	//解析多个模板
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path+"/view/index1.html", path+"/view/index2.html")) //让go自己帮我们处理异常
	//执行  选择响应的模板
	t.Execute(w, "Hello GO!")
}
func main() {
	http.HandleFunc("/template", handler)
	http.ListenAndServe(":8080", nil)
}
