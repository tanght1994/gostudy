package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
)

// fsnotify库可以监控文件夹or文件的变化
// 如果有变化, 用户可以从fsnotify的chann中读取到事件

func main() {
	// 创建监视器
	watcher, err := fsnotify.NewWatcher()
	must(err)

	// 监控f1.txt
	err = watcher.Add("./file/f1.txt")
	must(err)
	// 监控f2.txt
	err = watcher.Add("./file/f2.txt")
	must(err)
	// 监控f3.txt
	err = watcher.Add("./file/f3.txt")
	must(err)
	// 监控不存在的文件.txt
	err = watcher.Add("./file/不存在的文件.txt")
	fmt.Println(err)
	// 监控hello文件夹
	err = watcher.Add("./file/hello")
	must(err)

	// 取消对f2.txt的监控
	watcher.Remove("./file/f2.txt")

	// 从监控器的Events管道中获取被监控文件的事件
	for {
		select {
		case e, ok := <-watcher.Events:
			if !ok {
				return
			}
			fmt.Println(e.Name) // 被监控文件的文件名
			fmt.Println(e.Op)   // 发生的事件
		case e, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println(e)
		}
	}
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
