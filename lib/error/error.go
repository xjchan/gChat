package errror

import (
	"fmt"
)

//CheckError 通用错误处理
func CheckError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}

}
