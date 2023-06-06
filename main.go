package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorld)
	r.HandleFunc("/monokle", Monokle)

	fmt.Printf("listening on port 8080")
	http.ListenAndServe(":8080", r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Wrold!")
}

func Monokle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Monokle 2.0 demo")
}
