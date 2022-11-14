package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	MaVariable string
}

func main() {
	fileServer := http.FileServer(http.Dir("./templates")) // New code
	http.Handle("/", fileServer)                           // New code

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
