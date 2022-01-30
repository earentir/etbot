package main

import (
	"fmt"
	"io"
	"net/http"
)

func HTTPCheckReponce(url string) bool {
	responce, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}

	if responce.StatusCode >= 200 && responce.StatusCode <= 299 {
		return true
	} else {
		return false
	}
}

func HTTPGetBody(url string) string {
	responce, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
	}

	if responce.StatusCode >= 200 && responce.StatusCode <= 299 {
		bdystrmsg, err := io.ReadAll(responce.Body)
		if err != nil {
			fmt.Print(err)
		}

		return string(bdystrmsg)
	} else {
		return ""
	}

}
