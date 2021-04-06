package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Go is neat!</h1>")
	fmt.Fprintf(w, "<h3> Expert Web Design! </h3>")
	fmt.Fprintf(w, "<h3>© Copyright Reserved</h3>")

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8923", nil)
}
