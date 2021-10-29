package main

import "net/http"

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

//给客户端响应一个HTML页面
func Handler(w http.ResponseWriter, r *http.Request) {
	html := `
	<html>
		<head>
			<title>Hello Client！</title>
			<meta charset="UTF-8">
		<head>
		<body>
			网页回应：您的请求已收到！
		</body>
	<html>
`
	w.Write([]byte(html))
}
