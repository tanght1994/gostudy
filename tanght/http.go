package main

import (
	"bufio"
	"io"
	"net/http"
	"os"
)

func http_main() {
	http_server()
}

func http_header() {
	bufio.NewWriter(os.Stdout)
	header := http.Header{}
	header.Add("tanght", "hahah")
	header.Add("tanght", "xixix")
	header.Write(os.Stdout)
}

func http_server() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/hello", http_hello)
	mux.HandleFunc("/hello/", http_hello_)
	// mux.HandleFunc("/world", http_world)
	mux.HandleFunc("/world/", http_world_)
	// mux.HandleFunc("/hello/world", http_hello_world)
	mux.HandleFunc("/hello/world/", http_hello_world_)
	mux.HandleFunc("/", http_anything)
	server := http.Server{Addr: ":8000", Handler: mux}
	server.ListenAndServe()
}

func http_hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func http_hello_(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello/")
}

func http_world(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "world")
}

func http_world_(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "world/")
}

func http_hello_world(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello/world")
}

func http_hello_world_(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello/world/")
}

func http_anything(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "anything")
}
