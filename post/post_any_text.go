package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	reader := strings.NewReader("Hello, world!")
	resp, err := http.Post("http://localhost:18888", "text/plain", reader)
	if err != nil {
		// fault to POST
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
