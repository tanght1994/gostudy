package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

func main() {
	reporter := zipkinhttp.NewReporter("http://localhost:12345/")
	zkTracer, _ := zipkin.NewTracer(reporter)
	fmt.Println(zkTracer)
	port := 8080
	addr := fmt.Sprintf(":%d", port)
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/home", homeHandler)
	mux.HandleFunc("/async", serviceHandler)
	mux.HandleFunc("/service", serviceHandler)
	mux.HandleFunc("/db", dbHandler)
	fmt.Printf("Go to http://localhost:%d/home to start a request!\n", port)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`<a href="/home"> Click here to start a request </a>`))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	span := opentracing.StartSpan("/home")
	defer span.Finish()
	w.Write([]byte("Request started"))
	go func() {
		http.Get("http://localhost:8080/async")
	}()
	http.Get("http://localhost:8080/service")
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
	w.Write([]byte("Request done!"))
}

func serviceHandler(w http.ResponseWriter, r *http.Request) {
	http.Get("http://localhost:8080/db")
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
}
