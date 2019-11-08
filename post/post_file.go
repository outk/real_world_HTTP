package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	file, err := os.Open("../simpleget/main.go")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", file)
	if err != nil {
		// fault to POST
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
