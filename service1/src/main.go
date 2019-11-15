package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello, " + message
	fmt.Fprintf(w, message)
}

func main() {

	http.HandleFunc("/", sayHello)
	fmt.Println("Listening on port 8881")
	log.Fatal(http.ListenAndServe(":8881", nil))

}
