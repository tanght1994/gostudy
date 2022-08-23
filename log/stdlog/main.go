package main

import "log"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println(1, 2, 3, 4)
}
