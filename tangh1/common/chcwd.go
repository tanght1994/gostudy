package common

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

func chcwd() error {
	// 将当前工作目录切换到此程序所在的目录
	p, e := exec.LookPath(os.Args[0])
	if e != nil {
		return e
	}
	p = strings.ReplaceAll(p, "\\", "/")
	if i := strings.LastIndex(p, "/"); i == -1 {
		return errors.New("can't find /")
	} else {
		p = p[:i+1]
	}
	if e := os.Chdir(p); e != nil {
		return e
	}
	return nil
}
