package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	hc "hangmanweb/hangman_classic"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type dataSt struct {
	Word       string
	UsedLetter []string
	Letter     string
	HiddenWord string
	Tries      int
}

type intrSt struct {
}

var intr intrSt
var data dataSt

func main() {
	data.Tries = 10
	rand.Seed(time.Now().UTC().UnixNano())

	f, _ := os.OpenFile("../hangman_classic/main/words.txt", os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)

	wordlist := []string{}
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	data.Word = wordlist[rand.Intn(len(wordlist))]
	data.HiddenWord = hc.CreateWord(data.Word)

	http.HandleFunc("/", Handler_index) // Ici, quand on arrive sur la racine, on appelle la fonction Handler
	http.HandleFunc("/login", Handler_login)
	http.HandleFunc("/win", Handler_win)
	http.HandleFunc("/loose", Handler_loose)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Le Serveur démarre sur le port 8080\n")
	// http.HandleFunc("/hangman", Checker) // Ici, on redirige vers /hangman pour effectuer les fonctions POST
	http.ListenAndServe(":8080", nil)
}

func Handler_login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/login")
	tmpl1 := template.Must(template.ParseFiles("./static/login.html"))
	if r.Method == "POST" {
		clicked := r.FormValue("input")
		if clicked != "" {
			file, _ := json.MarshalIndent(clicked, "", "")

			_ = ioutil.WriteFile("clients.json", file, 0644)
			http.Redirect(w, r, "/", 301)
		}
	}
	tmpl1.Execute(w, data)
}

func Handler_index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/index")
	tmpl2 := template.Must(template.ParseFiles("./static/index.html"))
	data.Letter = r.FormValue("input")
	fmt.Print("DATA LETTER", data.Letter)
	var state string

	data.HiddenWord, state = hc.IsInputOk(data.Letter, data.Word, data.HiddenWord, &data.UsedLetter)

	if state == "fail" {
		data.Tries--
		if data.Tries == 0 {
			fmt.Print("You've lost!")
			http.Redirect(w, r, "/loose", 301)
		} else {
			fmt.Println("This letter is not in the word, you've got ", data.Tries, "tries left")
		}
	} else if state == "good" {
		if data.Word == data.HiddenWord {
			fmt.Println("Congrats you've won!")
			http.Redirect(w, r, "./win", 301)
		} else {
			fmt.Println("This letter is in the word")
		}
	} else if state == "usedletter" {
		fmt.Println("You've already used this letter, try again.")
	} else if state == "wordwrong" {
		data.Tries--
		data.Tries--
		if data.Tries == 0 {
			http.Redirect(w, r, "/loose", 301)
			fmt.Print("You've lost!")
		} else {
			fmt.Println("Wrong word you've got ", data.Tries, "tries left")
		}
	} else if state == "wordgood" {
		fmt.Println("Congrats you've won!")
		http.Redirect(w, r, "/win", 301)
	} else if state == "error" {
		fmt.Println("Invalid letter use another one.")
	} else if state == "wordinvalid" {
		fmt.Print("This word is invalid, try again")
	}

	switch r.Method {
	case "POST": //
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	}
	tmpl2.Execute(w, data)
}

func Handler_win(w http.ResponseWriter, r *http.Request) {
	tmpl_win := template.Must(template.ParseFiles("./static/win.html"))
	tmpl_win.Execute(w, intr)
}

func Handler_loose(w http.ResponseWriter, r *http.Request) {
	tmpl_loose := template.Must(template.ParseFiles("./static/loose.html"))
	tmpl_loose.Execute(w, intr)
}
