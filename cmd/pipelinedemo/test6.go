package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
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

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("tplfile/login.gtpl")
		t.Execute(w, token)
	} else {
		//POST
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			fmt.Println("找到token!")
		} else {
			fmt.Println("token不存在!")
		}

		username := template.HTMLEscapeString(r.Form.Get("username"))
		fmt.Println("username: ", username)
		fmt.Println("password: ", r.Form["password"])
		if len(r.Form.Get("username")) == 0 {
			fmt.Println("username is empty.")
		}
		fmt.Println("性別: ", r.Form.Get("gender"))
		for k, v := range r.Form["interest"] {
			fmt.Printf("愛好%d: %s\n", k, v)
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	if r.Method == "GET" {
		curtime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(curtime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("tplfile/upload.gtpl")
		t.Execute(w, token)
	} else {
		//POST
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./test"+handler.Filename, os.O_WRONLY|os.O_CREATE, 066)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}

}
