package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

const port = "8080"

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Request by:", r.RemoteAddr)
		fmt.Fprintln(w, "Served by:", runtime.GOOS)
	})
	log.Println("Starting server at port:", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}
