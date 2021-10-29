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

//给客户端响应一个模板文件
func handler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	//t , _ := template.ParseFiles("http/给客户端响应/model.html")
	//上面代码的具体逻辑分为两步： t := template.New("xxx.html")   t , _ := t.ParseFiles("xxx.html")
	//上面代码一般会写成t:=template.Must(template.ParseFiles("xxx.html"))  具体实现在下方

	path, _ := os.Getwd()
	// Must()会在里面函数err非空时产生一个panic  Must()会让go自动帮我们处理异常
	t := template.Must(template.ParseFiles(path + "/view/index.html"))
	//执行  第二个参数是给客户端响应的内容
	t.Execute(w, "Hello Golang")
}
