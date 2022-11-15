package main

import (
	"bufio"
	"fmt"
	"hangman_classic"
	"math/rand"
	"os"
	"time"
)

func main() {
	// On récupère les arguments
	args := os.Args[1:]
	// S'il n'y a pas qu'un seul argument, on arrête le programme
	if len(args) != 1 {
		os.Exit(0)
	}
	// Sinon on ouvre le fichier correspondant au nom passé en argument
	f, _ := os.OpenFile(args[0], os.O_RDWR, 0644)

	scanner := bufio.NewScanner(f)

	// On déclare notre liste de mots
	wordlist := []string{}

	// On ajoute tous les mots à notre liste de mots
	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}
	// On initialise le random
	rand.Seed(time.Now().UnixNano())
	// On génère un nombre aléatoire qui va récupérer un mot grâce à son index dans notre liste de mots
	randomword := wordlist[rand.Intn(len(wordlist))]

	// On ouvre le fichier dans lequel se trouvent les positions du hangman
	j, _ := os.OpenFile("hangman.txt", os.O_RDWR, 0644)

	scanner = bufio.NewScanner(j)

	// On crée une liste qui va contenir nos 10 positions possibles pour le pendu
	hang := make([]string, 10)
	count := 0
	index := 0
	for scanner.Scan() {
		count++
		// On crée nos positions de pendu en gardant des groupes de 7 lignes dans le fichier hangman.txt
		hang[index] += scanner.Text() + "\n"
		if count == 7 {
			count = 0
			index++
		}
	}
	choose := ""
	// On crée notre mot avec quelques lettres d'afficher
	randomwordhide := hangman_classic.CreateWord(randomword)
	// On déclare nos variables qui vont servir à l'avancée de notre jeu
	state := ""
	essaie := 10
	usedletter := []string{}
	// On clear le terminal
	hangman_classic.Clear()
	// On affiche l'état de la partie
	fmt.Printf("La partie commence, tu possèdes actuellement %v essais !\n", essaie)
	fmt.Println(randomwordhide)
	// On fait une boucle infinie que l'on stoppera lorsque le joueur aura gagné ou perdu
	for true {
		// Le joueur choisit une lettre ou un mot
		fmt.Print("Choose: ")
		fmt.Scanln(&choose)
		// On vérifie si ce qu'il a marqué est valide
		randomwordhide, state = hangman_classic.IsInputOk(choose, randomword, randomwordhide, &usedletter)
		// On clear le terminal
		hangman_classic.Clear()
		// Si le joueur a déjà fait une erreur
		if essaie != 10 {
			// On affiche l'état du pendu
			fmt.Print(hang[9-essaie])
		}
		// Si ce que le joueur a marqué n'est pas valide
		if state == "fail" {
			// On diminue le nombre d'essai restant du joueur
			essaie--
			// On affiche l'état de la partie
			fmt.Print(hang[9-essaie])
			fmt.Printf("La lettre %v n'est pas comprise dans le mot, il ne te reste plus que : %v essais\n", choose, essaie)
			// Si la lettre a déjà été utilisé
		} else if state == "usedletter" {
			// On affiche le message correspondant
			fmt.Printf("Lettre déjà utiliser\n")
			// Si la lettre est valide
		} else if state == "good" {
			// On affiche le message correspondant
			fmt.Printf("La lettre %v est bien comprise dans le mot\n", choose)
			// Si le mot rentré n'est pas de la bonne taille
		} else if state == "wordinvalid" {
			// On affiche le message correspondant
			fmt.Printf("Le format n'est pas valide, veuillez rentrer une lettre ou un mot de bonne taille\n")
			// Si le mot est le bon
		} else if state == "wordgood" {
			// On affiche le message correspondant
			fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", essaie, randomword)
			// On arrête le programme
			os.Exit(0)
			// Si l'input n'est pas une lettre
		} else if state == "error" {
			// On affiche le message correspondant
			fmt.Println("La lettre est invalide, veuillez recommencer")
			// Si la mot n'est pas le bon
		} else if state == "wordwrong" {
			// On retire 2 essais au lieu de 1
			essaie -= 2
			// On affiche le message correspondant
			fmt.Printf("Le mot proposé n'est pas le bon, il te reste %v essais\n", essaie)
		}
		// Si le joueur n'a plus d'essais
		if essaie <= 0 {
			// On clear le terminal
			hangman_classic.Clear()
			// On affiche l'état de la partie
			fmt.Print(hang[9])
			fmt.Printf("Tu as perdu, le mot était : %v", randomword)
			// On arrête le programme
			os.Exit(0)
		}
		fmt.Println(randomwordhide)
		// Si le mot a été totalement découvert
		if randomwordhide == randomword {
			// On clear le terminal
			hangman_classic.Clear()
			// On affiche le message correspondant
			fmt.Printf("Tu as trouvé, il te restait %v essai(s), le mot est : %v", essaie, randomword)
			// On arrête le programme
			os.Exit(0)
		}
	}
}
