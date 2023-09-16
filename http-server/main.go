package main

import (
	"main/api"
	"net/http"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe("127.0.0.1:8080", srv)
}