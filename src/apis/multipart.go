package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"net/textproto"
)

func main(){
	var data bytes.Buffer
	dataWriter := multipart.NewWriter(&data)
	dataWriter.WriteField("name", "Michael Jackson")

	part := make(textproto.MIMEHeader)
	part.Set("Content-Type", "image/png")
	part.Set("Content-Disposition", 	`form-data; name="thumbnail"; filename="copy.png"`)
	fileWriter, err := dataWriter.CreatePart(part) //
	// fileWriter, err := dataWriter.CreateFormFile("thumbnail", "/copy.png") //
	if err != nil {
		panic(err)
	}

	readFile, err := os.Open("assets/test.png")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	io.Copy(fileWriter, readFile)
	dataWriter.Close()

	resp, err := http.Post("http://localhost:18888", dataWriter.FormDataContentType(), &data)
	// dataWriter.FormDataContentType returns "multipart/form-data; boundary=" + writer.Boundary()
	if err != nil {
		panic(err)
	}
	log.Println("Status", resp.Status)
}
