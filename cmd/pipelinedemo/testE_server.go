package main

import (
	"fmt"
	"github.com/golang/net/websocket"
	"net/http"
	"time"
)

func main() {
	fmt.Println("启动时间")
	fmt.Println(time.Now())

	http.HandleFunc("/", h_index)
	//绑定socket方法
	http.Handle("/webSocket", websocket.Handler(h_webSocket))
	//开始监听
	http.ListenAndServe(":666", nil)

}

func h_index(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func h_webSocket(ws *websocket.Conn) {
	var str string

	for {
		err := websocket.Message.Receive(ws, &str)
		if err != nil {
			panic(err)
		}
		fmt.Println("接收的數據: ", str)
		err = websocket.Message.Send(ws, time.Now().Format("2006-01-02 15:04:05")+" 來自服務器的回應: "+str)
		if err != nil {
			panic(err)
		}
	}

}
