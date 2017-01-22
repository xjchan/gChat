package main

import (
	"fmt"
	// "github.com/xjchan/gChat/client"
	"github.com/xjchan/gChat/lib/config"
	myError "github.com/xjchan/gChat/lib/error"
	"github.com/xjchan/gChat/test"
	"net"
	// "net/http"
)

func main() {
	configFolder := "./configs/"
	config.SetConfigFolder(configFolder) //初始配置目录

	t := test.TestConfig{}
	err := config.GetConfig("test.yaml", &t)
	myError.CheckError(err)

	fmt.Println(t)

	fmt.Println("Hello Wrold")

	// var clientMap map[net.Conn]client.Client

	service := ":7777"
	tcpAddr, _ := net.ResolveTCPAddr("tcp4", service)
	listener, _ := net.ListenTCP("tcp", tcpAddr)

	for {
		conn, _ := listener.Accept()
		go func() {
			defer conn.Close()
			r := make([]byte, 512)
			for {
				conn.Read(r)
				fmt.Println(string(r))
			}
		}()
	}

}
