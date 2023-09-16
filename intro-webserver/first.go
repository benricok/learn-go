package main

import "net/http"

func main() {
	http.HandleFunc("/hello-world", func (w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}