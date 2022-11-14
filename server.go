package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type hang struct {
	word   string
	letter string
}

func main() {

	http.HandleFunc("/", Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	//
	fs := http.FileServer(http.Dir("./static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))
	//

	http.HandleFunc("/hangman", Handler) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
	fmt.Print("Le Serveur démarre sur le port 8080\n")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/index.html"))

	switch r.Method {
	case "POST": //
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	}
	variable := r.Form.Get("input")

	data := hang{
		word:   "EHEH",
		letter: variable,
	}

	tmpl.Execute(w, data)
	print(data.letter)
}
