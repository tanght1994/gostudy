// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

// Example service program that beeps.
//
// The program demonstrates how to create Windows service and
// install / remove it on a computer. It also shows how to
// stop / start / pause / continue any service, and how to
// write to event log. It also shows how to use debug
// facilities available in debug package.
package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/shirou/gopsutil/process"
)

func main() {

	reg := regexp.MustCompile(`.*tkf-exporter.*`)

	for {
		ps, err := process.Processes()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, p := range ps {
			name, err := p.Name()
			if err != nil {
				fmt.Println(err)
				continue
			}
			if !reg.Match([]byte(name)) {
				continue
			}
			utc, err := p.CreateTime()
			fmt.Println(utc)
		}
		time.Sleep(time.Second)
	}
}
