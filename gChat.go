package main

import (
	"fmt"
	"github.com/xjchan/gChat/lib/config"
	myError "github.com/xjchan/gChat/lib/error"
	"github.com/xjchan/gChat/test"
)

func main() {
	configFolder := "./configs/"
	config.SetConfigFolder(configFolder) //初始配置目录

	t := test.TestConfig{}
	err := config.GetConfig("test.yaml", &t)
	myError.CheckError(err)

	fmt.Println(t)

	fmt.Println("Hello Wrold")

}
