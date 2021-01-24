package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

func handler(w http.ResponseWriter, r *http.Request){
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	log.Printf("URL: %s\n", r.URL.String())
	log.Printf("Query: %v\n", r.URL.Query())
	// to remove character %, add \n at the end
	fmt.Fprint(w, "<html><body>hello</body><html>\n")
}

func handlerDigest(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %s\n", r.URL.String())
	fmt.Printf("Query: %v\n", r.URL.Query())
	fmt.Printf("Proto: %2\n", r.Proto)
	fmt.Printf("Method: %2\n", r.Method)
	fmt.Printf("Header: %v\n", r.Header)
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("--body--\n%s\n", string(body))
	if _, ok := r.Header["Authorization"]; !ok {
		w.Header().Add("WWW-Authenticate", `Digest realm="Secret Zone", nonce="TgLc25U2BQA=f510a2780473e18e6587be702c2e67fe2b04afd", algorithm=MD5, qop="auth"`)
		w.WriteHeader(http.StatusUnauthorized)
	} else {
		fmt.Fprintf(w, "<html><body>secret page</body></html>\n")
	}
}

func handlerCookie(w http.ResponseWriter, r *http.Request){
	c := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
	if _, ok := r.Header["Cookie"]; ok {
		fmt.Fprintf(w, "<html><body> 두번째 이후</body></html>\n")
	} else {
		fmt.Fprintf(w, "<html><body>첫방문</body></html>\n")
	}
}

func main(){
	var httpServer http.Server
	http.HandleFunc("/", handler)
	http.HandleFunc("/digest", handlerDigest)
	http.HandleFunc("/cookie", handlerCookie)
	log.Println("start http listening :18888")
	httpServer.Addr = ":18888"
	log.Println(httpServer.ListenAndServe())
}