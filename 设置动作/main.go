package main

import (
	"html/template"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	path, _ := os.Getwd()
	t := template.Must(template.ParseFiles(path + "/view/index.html"))
	//执行
	t.Execute(w, "狸猫")
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
