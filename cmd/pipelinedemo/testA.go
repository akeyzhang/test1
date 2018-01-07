package main

import (
	"fmt"
	"github.com/json-iterator/go"
)

type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func main() {
	str := `{"Servers":[{"serverName":"Shanghai_VPN","ServerIP":"127.0.0.1"},{"serverName":"Beijing_VPN","ServerIP":"127.0.0.2"},{"serverName":"gaungzhou_VPN","ServerIP":"127.0.0.3"},{"serverName":"dongguan_VPN","ServerIP":"127.0.0.4"},{"serverName":"shengzhen_VPN","ServerIP":"127.0.0.5"},{"serverName":"tongling_VPN","ServerIP":"127.0.0.6"},{"serverName":"hengyang_VPN","ServerIP":"127.0.0.7"},{"serverName":"nanjing_VPN","ServerIP":"127.0.0.8"}]}`
	var s Serverslice
	jsoniter.Unmarshal([]byte(str), &s)
	//fmt.Println(s)
	for _, server := range s.Servers {
		fmt.Println("ServerName:", server.ServerName)
		fmt.Println("ServerIP:", server.ServerIP)
	}
	/*var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	s.Servers = append(s.Servers, Server{ServerName: "gaungzhou_VPN", ServerIP: "127.0.0.3"})
	s.Servers = append(s.Servers, Server{ServerName: "dongguan_VPN", ServerIP: "127.0.0.4"})
	s.Servers = append(s.Servers, Server{ServerName: "shengzhen_VPN", ServerIP: "127.0.0.5"})
	s.Servers = append(s.Servers, Server{ServerName: "tongling_VPN", ServerIP: "127.0.0.6"})
	s.Servers = append(s.Servers, Server{ServerName: "hengyang_VPN", ServerIP: "127.0.0.7"})
	s.Servers = append(s.Servers, Server{ServerName: "nanjing_VPN", ServerIP: "127.0.0.8"})
	b, err := jsoniter.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))*/

}
