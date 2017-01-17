package main

import (
	"fmt"
	"github.com/xjchan/gChat/client"
	"github.com/xjchan/gChat/lib/config"
	myError "github.com/xjchan/gChat/lib/error"
	"github.com/xjchan/gChat/test"
	"net"
	"net/http"
)

func main() {
	configFolder := "./configs/"
	config.SetConfigFolder(configFolder) //初始配置目录

	t := test.TestConfig{}
	err := config.GetConfig("test.yaml", &t)
	myError.CheckError(err)

	fmt.Println(t)

	fmt.Println("Hello Wrold")

	var clientMap map[net.Conn]client.Client

	local := http.ListenAndServe(":8888", func() {

	})

}
