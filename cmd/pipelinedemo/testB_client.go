package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func checkError(xh int, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%d :Fatal error: %s\n", xh, err.Error())
	}
}

func main() {
	service := "192.168.1.129:666" //baidu
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(1, err)
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)
	checkError(2, err)
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(3, err)
	result, err := ioutil.ReadAll(conn)
	checkError(4, err)
	fmt.Println(string(result))
	conn.Close()

}
