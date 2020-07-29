package main

import (
	"fmt"
	"log"
	"net/http"
)

func Hello( w http.ResponseWriter, r * http.Request)  {
	fmt.Fprintf(w, "Hello Golang")
}

func ByteDance(w http.ResponseWriter, r * http.Request)  {
	// w: 跟响应相关
	// r: 跟请求有关
	fmt.Fprintf(w, "<h1>字节跳动[ByteDance]</h1>")	// 写入响应内容
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/bytedance", ByteDance)

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}


}
