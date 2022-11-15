package main

import (
	"fmt"
	hc "hangmanweb/hangman_classic"
	"html/template"
	"net/http"
	"os"
)

type dataSt struct {
	Word       string
	wordtofind string
	usedletter []string
}

var data dataSt

func main() {

	http.HandleFunc("/", Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler

	fs := http.FileServer(http.Dir("./static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	fmt.Print("Le Serveur démarre sur le port 8080\n")
	http.HandleFunc("/hangman", Checker) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
	wordtofind = hc.CreateWord(os.Open())
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
	tmpl.Execute(w, data)
}

func Checker(w http.ResponseWriter, r *http.Request) {
	variable := r.FormValue("input")
	println(variable)
	hc.IsInputOk(variable, data.Word, data.wordtofind, &data.usedletter)
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
