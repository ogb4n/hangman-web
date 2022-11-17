package main

import (
	"fmt"
	hc "hangmanweb/hangman_classic"
	"html/template"
	"net/http"
	"io/ioutil"
	"os"
	"math/rand"
	"bufio"
	"time"
)

type dataSt struct {
	Word       string
	UsedLetter []string
	Letter string
	HiddenWord string
	Tries int
}

var data dataSt

func main() {
	data.Tries = 10
	content, _ := ioutil.ReadFile("/home/alexandre/hangman-web/hangman_classic/main/words.txt")
	rand.Seed(time.Now().UTC().UnixNano())

	file, _ := os.Open("/home/alexandre/hangman-web/hangman_classic/main/words.txt")
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}

	random := rand.Intn(lineCount)
	var word string
	var list []string
	line := 1

	for _, j := range string(content) {
		if j == 13 {
			list = append(list, word)
			line++
			word = ""
		} else {
			word+= string(j)
		}
	}

	data.Word = list[random]
	data.HiddenWord = hc.CreateWord(data.Word)

	http.HandleFunc("/", Handler) // Ici, quand on arrive sur la racine, on appelle la fonction Handler

	fs := http.FileServer(http.Dir("./static/css"))
	http.Handle("/css/", http.StripPrefix("/css/", fs))

	fmt.Print("Le Serveur démarre sur le port 8080\n")
	// http.HandleFunc("/hangman", Checker) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var state string
	tmpl := template.Must(template.ParseFiles("static/index.html"))
	data.Letter = r.FormValue("input")
	
	switch r.Method {
		case "POST": //
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	}
	tmpl.Execute(w, data)
	
	data.HiddenWord, state = hc.IsInputOk(data.Letter, data.Word, data.HiddenWord, &data.UsedLetter)
	fmt.Print(data)

	if state == "fail" {
		data.Tries--
		fmt.Println("This letter is not in the word, you've got ",data.Tries, "tries left")
	} else if state == "good" {
		fmt.Println("This letter is in the word")
	} else if state == "usedletter" {
		fmt.Println("You've already used this letter, try again.")
	} else if state == "wordwrong" {
		data.Tries--
		data.Tries--
		fmt.Println("Wrong word you've got ",data.Tries, "tries left")
	} else if state == "wordgood" {
		fmt.Println("Congrats you've won!")
	} else if state == "error" {
		fmt.Println("Invalid letter use another one.")
	}
}