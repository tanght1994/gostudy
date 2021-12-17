package common

import (
	"fmt"
	"os"
)

var Log *logger

func init() {
	// 修改当前工作目录为此程序所在的目录
	e := chcwd()
	if e != nil {
		fmt.Println("chcwd error, ", e)
		os.Exit(1)
	}

	// 创建全局Logger
	Log, e = newLogger()
	if e != nil {
		fmt.Println("newLogger error, ", e)
		os.Exit(1)
	}
}
