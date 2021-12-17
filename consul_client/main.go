package main

import (
	"fmt"

	consulapi "github.com/hashicorp/consul/api"
)

func main() {
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		fmt.Println("1", err)
	}
	services, err := client.Agent().Services()
	if err != nil {
		fmt.Println("2", err)
	}

	if _, found := services["tanght"]; !found {
		fmt.Println("3")
	}
}
