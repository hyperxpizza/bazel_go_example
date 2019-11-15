package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {
	//static files handler
	fs := http.FileServer(http.Dir("static"))

	//bind urls and functions
	http.Handle("/", fs)
	http.HandleFunc("/increment", incrementCounter)

	//run on localhost:8888
	log.Fatal(http.ListenAndServe(":8888", nil))

}
