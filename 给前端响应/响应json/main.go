package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

//给客户端响应一个HTML页面
func Handler(w http.ResponseWriter, r *http.Request) {
	//设置响应报文头内容
	w.Header().Set("Content-Type", "hello golang")
	//创建一个User 用于响应给客户端
	user := &User{Id: 1, UserName: "admin", Password: "123456", Email: "admin@qq.com"}
	//将User序列化为一个json
	slice, _ := json.Marshal(user)
	//将json响应给客户端
	w.Write(slice) //传入json的byte类型切片  Write()函数会自动将byte切片转换为string并传给前端
}
