package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	fmt.Println("==> 🌍 Server ready, listening to http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
