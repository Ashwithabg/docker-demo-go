package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "hello\n")
}


func healthCheck(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintf(writer, "Running\n")
}

func main() {
	address := ":8080"
	http.HandleFunc("/", hello)
	http.HandleFunc("/healthcheck", healthCheck)
	log.Println("listen on", address)
	log.Fatal(http.ListenAndServe(address, nil))
}
