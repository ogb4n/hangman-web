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
	fileServer := http.FileServer(http.Dir("./templates"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
