package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
)

func main() {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	writer.WriteField("name", "Michael Jackson")

	// fileWriter, err := writer.CreateFormFile("thumbnail", "resume.png")
	// if err != nil {
	// 	panic(err)
	// }
	// readFile, err := os.Open("resume.png")
	// if err != nil {
	// 	// fault to read file
	// 	panic(err)
	// }
	// defer readFile.Close()
	// io.Copy(fileWriter, readFile)

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/jpeg")
	part.Set("Content-Disposition", `form-data; name="thumbnail"; filename="resume.png"`)
	fileWriter, err := writer.CreatePart(part)
	if err != nil {
		panic(err)
	}
	readFile, err := os.Open("resume.png")
	if err != nil {
		panic(err)
	}
	io.Copy(fileWriter, readFile)
	writer.Close()

	resp, err := http.Post("http://localhost:18888", writer.FormDataContentType(), &buffer)
	if err != nil {
		// fault to POST
		panic(err)
	}
	log.Println("Status: ", resp.Status)
}
