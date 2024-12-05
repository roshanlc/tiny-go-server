package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

// TODO: read the port from env var
const port = "8080"

func main() {

	envPort := os.Getenv("TINY_PORT")
	if envPort == "" {
		envPort = port
	}

	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request:", r.Method, " : ", r.Host, " : ", r.UserAgent())
		fmt.Fprintln(w, "Request by:", r.RemoteAddr)
		fmt.Fprintln(w, "Served by:", runtime.GOOS)
	})
	log.Println("Starting server at port:", envPort)
	log.Fatal(http.ListenAndServe(fmt.Sprint(":", envPort), nil))
}
