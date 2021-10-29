package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

//给客户端响应多个模板文件
func handler(w http.ResponseWriter, r *http.Request) {
	//解析多个模板
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path+"/view/index1.html", path+"/view/index2.html")) //让go自己帮我们处理异常
	//执行  选择要响应给前端的模板
	t.ExecuteTemplate(w, "index2.html", "Hello World")
}
