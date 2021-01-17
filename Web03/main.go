package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "创建多路复用器")
}

func main() {
	//创建多路复用器
	mux := http.NewServeMux()

	//http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)

	//http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", mux)
}
