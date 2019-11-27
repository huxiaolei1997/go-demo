package main

import (
	"net/http"
)

func main() {
	// 实现一个简单的HTTP服务器，在浏览器里输入http://127.0.0.1:8080即可浏览文件，
	// 这些文件正是当前目录在HTTP服务器上的映射目录。
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8087", nil)
}
