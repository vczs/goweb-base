package main

import (
	"html/template"
	"net/http"
	"os"
)

//给客户端响应一个模板文件
func handler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/view/index.html")) //Must()会让go自动帮我们处理异常
	//执行  data是给客户端响应的内容
	a := 5
	t.Execute(w, a > 8)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
