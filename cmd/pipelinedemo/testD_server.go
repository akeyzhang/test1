package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/net/websocket"
	"net/http"
	"strings"
	"time"
)

var datas Datas
var users map[*websocket.Conn]string

type UserMsg struct {
	UserName string
	Msg      string
	DataType string
}

type UserData struct {
	UserName string
}

type Datas struct {
	UserMsgs  []UserMsg
	UserDatas []UserData
}

func main() {
	fmt.Println("启动时间")
	fmt.Println(time.Now())
	//初始化
	datas = Datas{}
	users = make(map[*websocket.Conn]string)

	//绑定效果页面
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
	var userMsg UserMsg
	var data string
	for {
		//判断是否重复连接
		if _, ok := users[ws]; !ok {
			users[ws] = "匿名"
		}
		userMsgsLen := len(datas.UserMsgs)
		fmt.Println("UserMsgs", userMsgsLen, "users长度：", len(users))
		//有消息时，全部分发送数据
		if userMsgsLen > 0 {
			b, errMarshl := json.Marshal(datas)
			if errMarshl != nil {
				fmt.Println("全局消息内容异常...")
				break
			}
			for key, _ := range users {
				errMarshl = websocket.Message.Send(key, string(b))
				if errMarshl != nil {
					//移除出错的链接
					delete(users, key)
					fmt.Println("发送出错...")
					break
				}
			}
			datas.UserMsgs = make([]UserMsg, 0)
		}
		fmt.Println("开始解析数据...")
		err := websocket.Message.Receive(ws, &data)
		fmt.Println("data：", data)
		if err != nil {
			//移除出错的链接
			delete(users, ws)
			fmt.Println("接收出错...")
			break
		}
		data = strings.Replace(data, "\n", "", 0)
		err = json.Unmarshal([]byte(data), &userMsg)
		if err != nil {
			fmt.Println("解析数据异常...")
			break
		}
		fmt.Println("请求数据类型：", userMsg.DataType)
		switch userMsg.DataType {
		case "send":
			//赋值对应的昵称到ws
			if _, ok := users[ws]; ok {
				users[ws] = userMsg.UserName
				//清除连接人昵称信息
				datas.UserDatas = make([]UserData, 0)
				//重新加载当前在线连接人
				for _, item := range users {
					userData := UserData{UserName: item}
					datas.UserDatas = append(datas.UserDatas, userData)
				}
			}
			datas.UserMsgs = append(datas.UserMsgs, userMsg)
		}
	}

}
