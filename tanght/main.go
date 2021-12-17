package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	s := `{"name": "tanght", "age": {"name": "haha", "a": 1111}}`
	d := make(map[string]interface{})
	if e := json.Unmarshal([]byte(s), &d); e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
}
