package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "咖啡色的羊驼"
	t := reflect.TypeOf(name)
	v := reflect.ValueOf(name)
}
