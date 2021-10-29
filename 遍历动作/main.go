package main

import (
	"html/template"
	"net/http"
	"os"
)

type people struct {
	id   int
	Name string
}

//给客户端响应一个模板文件
func handler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/view/index.html")) //Must()会让go自动帮我们处理异常
	p1 := &people{11, "十一"}
	p2 := &people{21, "二十一"}
	p3 := &people{31, "三十一"}
	//创建一个切片 将上面的people实例添加进去
	var peopleSlice []*people
	peopleSlice = append(peopleSlice, p1)
	peopleSlice = append(peopleSlice, p2)
	peopleSlice = append(peopleSlice, p3)
	t.Execute(w, peopleSlice)
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
