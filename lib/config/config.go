package config

import (
	// "errors"
	// "fmt"
	// "gopkg.in/yaml.v2"
	myError "github.com/xjchan/gChat/lib/error"
	"io/ioutil"
	"os"
	"reflect"
)

//Resolver 构造配置(config)
type Resolver interface {
	GetConfig([]byte) (interface{}, error)
}

var configFolder = "./"                       // 配置文件夹,全路径或者入口文件的相对路s
var configs = make(map[string]interface{}, 1) // 配置变量

//SetConfigFolder 设置全局的配置文件夹,单例
func SetConfigFolder(folder string) {

	configFolder = folder
}

//GetConfigFunc 获取
func GetConfigFunc(key string, resolver func(buf []byte) (interface{}, error)) (interface{}, error) {
	if configs[key] != nil { //已经初始化，直接返回
		return configs[key], nil
	}

	configFile := configFolder + key

	f, err := os.Open(configFile)
	myError.CheckError(err)

	buf, err := ioutil.ReadAll(f)
	myError.CheckError(err)

	r, err := resolver(buf)
	if err == nil {
		configs[key] = r
	}
	return r, err
}

//GetConfig 获取
func GetConfig(key string, resolver Resolver) error {
	var r interface{}
	var err error
	if configs[key] != nil { //已经初始化，直接返回
		r = configs[key]
	} else {
		configFile := configFolder + key //配置文件地址

		f, err := os.Open(configFile)
		myError.CheckError(err)

		buf, err := ioutil.ReadAll(f)
		myError.CheckError(err)

		r, err = resolver.GetConfig(buf)
		if err == nil {
			configs[key] = r
		}
	}

	e := reflect.ValueOf(resolver).Elem()
	if e.CanSet() {
		e.Set(reflect.ValueOf(r))
	}

	return err
}
