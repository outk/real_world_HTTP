package main

import (
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	client := &http.Client{}
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}

	// Add Header
	request.Header.Add("Content-Type", "image/jpg")

	// Confirm BASIC
	request.SetBasicAuth("username", "password")

	// Add another cookie
	request.AddCookie(&http.Cookie{Name: "test", Value: "value"})

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}
