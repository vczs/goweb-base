package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/get1", getCookieHandler_1)
	http.HandleFunc("/get2", getCookieHandler_2)
	http.ListenAndServe(":8080", nil)
}

//获取所有cookie
func getCookieHandler_1(w http.ResponseWriter, r *http.Request) {
	cookies := r.Header["Cookie"]
	fmt.Println("cookies:", cookies)
}

//获取某一个cookie
func getCookieHandler_2(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user1")
	fmt.Println("cookie:", cookie)
}
