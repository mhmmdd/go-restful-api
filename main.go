package main

import (
	"fmt"
	"log"
	"net/http"
)

func homaPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	fmt.Println("listen localhost:8081")
	http.HandleFunc("/", homaPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
