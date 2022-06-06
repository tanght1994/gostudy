package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", a)
	l, err := net.Listen("tcp", "0.0.0.0:10008")
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	http.Serve(l, mux)
}

func a(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "%s %s%s %s\n\n", req.Method, req.Host, req.RequestURI, req.Proto)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.Write([]byte(`read body error`))
	}
	for k, v := range req.Header {
		fmt.Fprint(res, k, ": ", v, "\n")
	}
	res.Write([]byte("\n"))
	res.Write(body)
}
