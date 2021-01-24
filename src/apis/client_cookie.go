package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func main(){
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
	}
	for i :=0; i<2; i++{
		resp, err:= client.Get("http://localhost:18888/cookie")
		if err != nil {
			panic(err)
		}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		log.Println(string(body))
	}
}

// a := Struct {
//  Member: 'Value'
//}