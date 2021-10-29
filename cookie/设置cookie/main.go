package main

import "net/http"

func main() {
	http.HandleFunc("/set1", setCookieHandler_1)
	http.HandleFunc("/set2", setCookieHandler_2)
	http.ListenAndServe(":8080", nil)
}

//设置一个默认cookie给客户端
func setCookieHandler_1(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user2",
		Value:    "admin2",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
}

//设置一个30秒时长的cookie给客户端
func setCookieHandler_2(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "user1",
		Value:    "admin1",
		HttpOnly: true,
		MaxAge:   30, //cookie的有效时间 30秒
	}
	w.Header().Set("Set-Cookie", cookie.String())
}
