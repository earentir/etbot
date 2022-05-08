package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func HTTPCheckResponse(uri string) bool {

	good, err := url.ParseRequestURI(uri)
	if err != nil {
		fmt.Printf("err: %e\n", err)
		return false
	} else {
		responce, err := http.Get(good.String())
		if err != nil {
			fmt.Print(err)
			return false
		} else {
			if responce.StatusCode >= 200 && responce.StatusCode <= 299 {
				return true
			} else {
				return false
			}
		}
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
		fmt.Printf("HTTP Error Code: %v", responce.StatusCode)
		return ""
	}

}
