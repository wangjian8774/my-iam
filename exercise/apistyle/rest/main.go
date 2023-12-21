package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", pong)
	log.Println("Server starting...")
	log.Fatal(http.ListenAndServe(":50052", nil))
}

func pong(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("pong"))
}
