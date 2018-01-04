package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", SayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("error: ", err)
	}

}

func SayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****************************")
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("schem: ", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("value:", strings.Join(v, " "))
	}
	fmt.Fprint(w, "hello, akeyzhang!")

}
