package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main(){
	client := &http.Client{}
	// http only support GET, HEAD, POST
	request, err := http.NewRequest("DELETE", "http://localhost:18888", nil)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
	values := url.Values{"test": {"value"}}
	read := strings.NewReader(values.Encode())
	log.Println(values)
	log.Println(values.Encode())
	log.Println(read)
}