package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/globalsign/mgo"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func main() {
	mem.SwapDevices()
	ps, err := process.Processes()
	must(err)
	cnt := 0
	for _, p := range ps {
		cnt++
		name, err := p.Name()
		if err != nil {
			fmt.Println(p.Pid, " error ", err)
		}
		if strings.Contains(name, "chrome") {
			fmt.Println(name, cnt)
			info, err := p.MemoryInfo()
			if err != nil {
				fmt.Println(name, "MemoryInfo() error ", err)
			}
			fmt.Println(info.String())
			fmt.Println(info.Data)
			fmt.Println(info.HWM)
			fmt.Println(info.Locked)
			fmt.Println(info.RSS)
			fmt.Println(info.Stack)
			fmt.Println(info.Swap)
			fmt.Println(info.VMS)
		}
	}

	URL := "mongodb://user:password@11.11.11.11,22.22.22.22/database?replicaSet=rs0"
	mgo.Dial(URL)
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
