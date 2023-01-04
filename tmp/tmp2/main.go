package main

import (
	"fmt"
	"os"
	"text/template"
)

// A ...
type A struct {
	Alerts B
}

// B ...
type B struct {
	Firing   []C
	Resolved []C
}

// C ...
type C struct {
	Annotations map[string]string
}

func main() {
	a := A{
		Alerts: B{
			Firing: []C{
				{Annotations: map[string]string{"description": "i am description 1"}},
				{Annotations: map[string]string{"description": "i am description 2"}},
				{Annotations: map[string]string{"description": "i am description 3"}},
				{Annotations: map[string]string{"description": "i am description 4"}},
			},
			Resolved: []C{
				{Annotations: map[string]string{"description": "i am description 5"}},
				{Annotations: map[string]string{"description": "i am description 6"}},
			},
		},
	}
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println(err)
		return
	}
	tmpl.Execute(os.Stdout, a)
}
