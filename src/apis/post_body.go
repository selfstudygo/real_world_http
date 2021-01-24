package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func main(){
	text, err := os.Open("assets/test.txt")
	// text := strings.NewReader("텍스트")
	if err != nil {
		panic(err)
	}
	resp, err := http.Post("http://localhost:18888", "text/plain", text)
	if err != nil{
		panic(err)
	}
	log.Println("Status:", resp.Status)
	log.Println(strings.NewReader("텍스트"))
}
