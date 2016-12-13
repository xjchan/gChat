package main

import (
	"fmt"
	"github.com/xjchan/gChat/tools"
	// "gopkg.in/yaml.v2"
	"github.com/xjchan/gChat/test"
	// "reflect"
)

func main() {
	configFolder := "./configs/"
	tools.SetConfigFolder(configFolder) //初始配置目录

	t := test.TestConfig{}
	err := tools.GetConfig("test.yaml", &t)
	tools.CheckError(err)
	// ty := reflect.TypeOf(t)
	// fmt.Println(ty.String())
	fmt.Println(t)

	fmt.Println("Hello Wrold")

}
