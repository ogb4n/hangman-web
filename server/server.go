package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
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
	//struct where i store everything needed to play the game
	Word                 string
	UsedLetter           []string
	Letter               string
	HiddenWord           string
	Tries                int
	Difficulty           string
	Username             string
	Score                int
	Error                string
	LeaderBoardScores    []int
	LeaderBoardUsernames []string
	PreviousWord         string
}

type clients struct {
	//struct where i store what's needed to login/register
	Passwords []string
	Usernames []string
	Scores    []int
	Which     int
}

var data dataSt
var clients_data clients

func main() {
	http.HandleFunc("/", Handler_index)
	http.HandleFunc("/login", Handler_login)
	http.HandleFunc("/win", Handler_win)
	http.HandleFunc("/loose", Handler_loose)
	//url of our funcs

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Print("Le Serveur dÃ©marre sur le port 8080\n")
	http.ListenAndServe(":8080", nil)
	//listening on port 8080
}

func Handler_login(w http.ResponseWriter, r *http.Request) {
	tmpl1 := template.Must(template.ParseFiles("./static/login.html"))
	//creating template for the loging page
	leaderboard()
	data.Error = ""
	if r.Method == "POST" {
		//getting our inputs
		if len(r.FormValue("input_username")) > 12 {
			data.Error = "Username must have a maximum lenght of 12 characters."
		} else {
			username := r.FormValue("input_username")
			password := r.FormValue("input_psswd")
			radiobutton := r.FormValue("contact")
			//choosing difficulty
			if radiobutton == "easy" {
				data.Difficulty = "../hangman_classic/main/words1.txt"
				create_game()
			} else if radiobutton == "medium" {
				data.Difficulty = "../hangman_classic/main/words2.txt"
				create_game()
			} else if radiobutton == "hard" {
				data.Difficulty = "../hangman_classic/main/words3.txt"
				create_game()
			}

			var isGood bool

			if username != "" && password != "" {
				for l, i := range clients_data.Usernames {
					//checking if user with this username already exists
					if string(i) == username {
						isGood = true
						//if yes we check if password is the right one then we'll login
						if hash(password) == clients_data.Passwords[l] {
							fmt.Println("Logging in")
							fmt.Println("Welcome back", username)
							clients_data.Which = l
							data.Username = clients_data.Usernames[clients_data.Which]
							data.Score = clients_data.Scores[clients_data.Which]
							http.Redirect(w, r, "/", http.StatusSeeOther)
							//redirect to the main page
						}
					}
				}
				if isGood {
					//if the password is wrong we just send an error
					fmt.Println("Wrong password.")
					data.Error = "Wrong password."
				} else {
					//if there's no account with this username, we create one
					fmt.Println("Creating your account", username)
					clients_data.Usernames = append(clients_data.Usernames, username)
					clients_data.Passwords = append(clients_data.Passwords, hash(password))
					clients_data.Scores = append(clients_data.Scores, 0)
					var count int
					for range clients_data.Usernames {
						count++
					}
					clients_data.Which = count - 1
					data.Username = clients_data.Usernames[clients_data.Which]
					data.Score = clients_data.Scores[clients_data.Which]
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}

			} else {
				if username == "" && password == "" {
					//case if user didnt input username and password
					data.Error = "Please insert a password and an username."
					fmt.Println("Please insert a password and an username.")
				} else if username == "" {
					//case if user didnt input username
					data.Error = "Please insert an username."
					fmt.Println("Please insert an username.")
				} else {
					//case if user didnt input password
					data.Error = "Please insert a password."
					fmt.Println("Please insert a password.")
				}
			}
		}
	}
	tmpl1.Execute(w, data)
}

func Handler_index(w http.ResponseWriter, r *http.Request) {
	tmpl2 := template.Must(template.ParseFiles("./static/index.html"))
	data.PreviousWord = data.Word
	leaderboard()
	fmt.Print(data.Word)

	if r.Method == "POST" {
		indexbutton := r.FormValue("indexbutton")
		if indexbutton != "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
	//creating template for the main page

	data.Letter = r.FormValue("input")
	//getting input from hangman
	var state string
	data.HiddenWord, state = hc.IsInputOk(data.Letter, data.Word, data.HiddenWord, &data.UsedLetter)
	//HiddenWord is the word with underscore that will change throughout the game, state returns if the input
	//is in the word or the word itself
	if state == "fail" {
		data.Tries--
		if data.Tries <= 0 {
			fmt.Print("You've lost!")
			http.Redirect(w, r, "/loose", http.StatusSeeOther)
		} else {
			fmt.Println("This letter is not in the word, you've got ", data.Tries, "tries left")
		}
	} else if state == "good" {
		if data.Word == data.HiddenWord {
			fmt.Println("Congrats you've won!")
			http.Redirect(w, r, "./win", http.StatusSeeOther)
		} else {
			fmt.Println("This letter is in the word")
		}
	} else if state == "usedletter" {
		fmt.Println("You've already used this letter, try again.")
	} else if state == "wordwrong" {
		data.Tries--
		data.Tries--
		if data.Tries <= 0 {
			http.Redirect(w, r, "/loose", http.StatusSeeOther)
			fmt.Print("You've lost!")
		} else {
			fmt.Println("Wrong word you've got ", data.Tries, "tries left")
		}
	} else if state == "wordgood" {
		fmt.Println("Congrats you've won!")
		http.Redirect(w, r, "/win", http.StatusSeeOther)
	} else if state == "error" {
		fmt.Println("Invalid letter use another one.")
	} else if state == "wordinvalid" {
		fmt.Print("This word is invalid, try again")
	}

	//shows more of the hiddenword / remove tries according to the state
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
	}

	tmpl2.Execute(w, data)
}

func Handler_win(w http.ResponseWriter, r *http.Request) {
	tmpl_win := template.Must(template.ParseFiles("./static/win.html"))
	if data.Difficulty == "../hangman_classic/main/words1.txt" {
		clients_data.Scores[clients_data.Which] += 1
	}
	if data.Difficulty == "../hangman_classic/main/words2.txt" {
		clients_data.Scores[clients_data.Which] += 2
	}
	if data.Difficulty == "../hangman_classic/main/words3.txt" {
		clients_data.Scores[clients_data.Which] += 3
	}

	//gives points to the user according to the difficulty
	saveClientData()
	//save data in clients.json
	create_game()
	data.Score = clients_data.Scores[clients_data.Which]
	//reset importants data to play the game
	tmpl_win.Execute(w, data)
}

func Handler_loose(w http.ResponseWriter, r *http.Request) {
	//same thing than handler_win
	tmpl_loose := template.Must(template.ParseFiles("./static/loose.html"))
	if data.Difficulty == "../hangman_classic/main/words1.txt" {
		clients_data.Scores[clients_data.Which] = 0
	}
	if data.Difficulty == "../hangman_classic/main/words2.txt" {
		clients_data.Scores[clients_data.Which] = 0
	}
	if data.Difficulty == "../hangman_classic/main/words3.txt" {
		clients_data.Scores[clients_data.Which] = 0
	}

	saveClientData()
	create_game()
	data.Score = clients_data.Scores[clients_data.Which]
	tmpl_loose.Execute(w, data)
}

func create_game() {
	rand.Seed(time.Now().UnixNano())
	data.Tries = 10
	f, _ := os.OpenFile(data.Difficulty, os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)
	//scanning the file where words are stored to get them

	wordlist := []string{}
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	random := rand.Intn(len(wordlist))
	//getting a random word from the list
	data.UsedLetter = nil
	data.Word = wordlist[random]
	data.HiddenWord = hc.CreateWord(data.Word)
	//reset all importants data to play a new game / launch the game
}

func saveClientData() {
	file, _ := json.MarshalIndent(clients_data, "", "")
	_ = ioutil.WriteFile("clients.json", file, 0644)
	//save clients_data struct to clients.json
}

func hash(password string) string {
	hash := sha1.New()
	hashInBytes := hash.Sum([]byte(password))[:20]
	return hex.EncodeToString(hashInBytes)
	//encoding passwords in sha1
}

func leaderboard() {
	username := make([]string, len(clients_data.Usernames))
	leaderboard := make([]int, len(clients_data.Scores))
	copy(username, clients_data.Usernames)
	copy(leaderboard, clients_data.Scores)
	//making a copy cause a simple username := clients_data.Username would make a mirror
	if len(leaderboard) > 0 {
		for j := 0; j != len(leaderboard)-1; j++ {
			for i := 0; i != len(leaderboard)-1; i++ {
				if leaderboard[i] < leaderboard[i+1] {
					swap := leaderboard[i]
					leaderboard[i] = leaderboard[i+1]
					leaderboard[i+1] = swap
					swap2 := username[i]
					username[i] = username[i+1]
					username[i+1] = swap2
				}
			}
		}
	}

	//sorting in non ascending order
	if len(leaderboard) > 5 {
		data.LeaderBoardScores = nil
		data.LeaderBoardUsernames = nil
		for i := 0; i != 5; i++ {
			data.LeaderBoardScores = append(data.LeaderBoardScores, leaderboard[i])
			data.LeaderBoardUsernames = append(data.LeaderBoardUsernames, username[i])
		}
	} else {
		data.LeaderBoardScores = leaderboard
		data.LeaderBoardUsernames = username
	}
}
