package main

import (
	"fmt"
	"net/http"
)

func startWebServer() {
	http.HandleFunc("/", func(result http.ResponseWriter, request *http.Request) {
		fmt.Println("hi ", request.Host)
	})

	http.ListenAndServe("127.0.0.1:9000", nil)
}
