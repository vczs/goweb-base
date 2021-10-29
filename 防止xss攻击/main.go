package main

import (
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9003", nil)
}

//防止xss攻击：只转换html内容 不转换<script>内容
func handler(w http.ResponseWriter, r *http.Request) {
	//定义模板
	t := template.New("model.tmpl")
	//给模板文件添加自定义函数
	t.Funcs(template.FuncMap{"safe": func(url string) template.HTML { return template.HTML(url) }})
	//解析模板
	t, _ = t.ParseFiles("./view/c3.tmpl")
	//渲染 执行
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='https://www.baidu.com'>百度</a>"
	t.Execute(w, map[string]string{"s1": str1, "s2": str2})
}
