package main

import (
	"fmt"
	hc "hangmanweb/hangman_classic"
	"html/template"
	"net/http"
)

type hang struct {
	Word   string
	Letter string
}

func main() {

	hc.AccentChecker("eee")

	http.HandleFunc("/", Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	//
	fs := http.FileServer(http.Dir("./static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//
	fmt.Print("Le Serveur démarre sur le port 8080\n")
	http.HandleFunc("/hangman", Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	switch r.Method {
	case "GET":
		fmt.Println("GET")
	case "POST": //
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	}
	variable := r.Form.Get("input")

	data := hang{
		Word:   "ça teste fort",
		Letter: variable,
	}

	tmpl.Execute(w, data)
	print(data.Letter, "\n")
}
