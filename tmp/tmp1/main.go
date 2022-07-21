package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := Server{}
	http.ListenAndServe("127.0.0.1:8000", server)
}

type Server struct{}

func (s Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	bys, err := ioutil.ReadAll(req.Body)
	must(err)
	fmt.Println(string(bys))
	res.Write([]byte("Hello World"))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
