package hangman_classic

import "math/rand"

// Fonction qui va créer le mot avec quelques lettres afficher
func CreateWord(word string) string {
	// On crée une liste de string
	wordtofind := []string{}
	// On boucle pour remplir cette liste de _
	for k := 0; k < len([]rune(word)); k++ {
		if word[k] == '-' {
			wordtofind = append(wordtofind, "-")
		} else {
			wordtofind = append(wordtofind, "_")
		}
	}
	// On boucle quelques fois pour afficher des lettres random
	for i := 0; i < (len([]rune(word))/2 - 1); i++ {
		// Génère un nombre random
		tempr := rand.Intn(len([]rune(word)))
		// Si à l'emplacement random se trouve un _ on rentre dan la boucle
		if wordtofind[tempr] == "_" {
			// On remplace le _ par la lettre correspondant à l'emplacement
			wordtofind[tempr] = string([]rune(word)[tempr])
		} else {
			i--
		}
	}
	myrep := ""
	// On boucle dans notre liste de string, pour la transformer en string
	for _, letter := range wordtofind {
		myrep += letter
	}
	// On retourne notre string
	return myrep
}
