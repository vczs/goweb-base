package main

import (
	"html/template"
	"net/http"
)

//原生方式
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9002", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	//定义模板
	t := template.New("model.tmpl")
	//给模板文件添加自定义函数
	hello := func(name string) string { return name + "666" } //创建自定义函数
	t.Funcs(template.FuncMap{"hello": hello})                 //添加到模板文件 添加操作一定要在解析模板文件之前
	//自定义模板引擎的标识符
	t.Delims("{[", "]}")
	//解析模板
	t, _ = t.ParseFiles("./view/c2.tmpl")
	//渲染 执行
	t.Execute(w, "王宝强")
}
