package main

import (
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path+"/view/index.html", path+"/view/index2.html")) //Must()会让go自动帮我们处理异常
	//执行  data是给客户端响应的内容
	t.Execute(w, "hello")
}

//重名的模板
func handlerPlus(w http.ResponseWriter, r *http.Request) {
	a := 13
	//解析模板
	var t *template.Template
	if a > 18 {
		path, _ := os.Getwd()
		t = template.Must(template.ParseFiles(path + "/view/index.html")) //Must()会让go自动帮我们处理异常
	} else {
		path, _ := os.Getwd()
		t = template.Must(template.ParseFiles(path+"/view/index.html", path+"/view/index1.html")) //Must()会让go自动帮我们处理异常
	}
	//执行  data是给客户端响应的内容
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/handler", handler)
	http.HandleFunc("/handlerPlus", handlerPlus)
	http.ListenAndServe(":8080", nil)
}
