package hangman_classic

import "strings"

// Fonction qui supprime les accents du string passé en paramètre
func AccentChecker(s string) string {
	s = strings.ToLower(s)
	// Déclaration de la variable qui va être retournée et qui va contenir le string sans les accents
	rep := ""
	// On boucle dans le string passé en paramètre
	for _, letter := range s {
		// On vérifie si la lettre est un e avec accent
		if letter == 'é' || letter == 'è' || letter == 'ê' || letter == 'ë' {
			// On enlève l'accent
			letter = 'e'
		}
		// On vérifie si la lettre est un a avec accent
		if letter == 'à' || letter == 'â' || letter == 'ä' {
			// On enlève l'accent
			letter = 'a'
		}
		// On vérifie si la lettre est un i avec accent
		if letter == 'i' || letter == 'î' || letter == 'ï' {
			// On enlève l'accent
			letter = 'i'
		}
		// On vérifie si la lettre est un u avec accent
		if letter == 'u' || letter == 'ù' || letter == 'û' || letter == 'ü' {
			// On enlève l'accent
			letter = 'u'
		}
		// On vérifie si la lettre est un c avec accent
		if letter == 'c' || letter == 'ç' {
			// On enlève l'accent
			letter = 'c'
		}
		// On vérifie si la lettre est un y avec accent
		if letter == 'y' || letter == 'ÿ' {
			// On enlève l'accent
			letter = 'y'
		}
		// On ajoute la lettre sans l'accent s'il y en avait un
		rep += string(letter)
	}
	// On retourne le string sans les accents
	return rep
}

// Fonction qui va retourner une liste de la lettre passée en paramètre, avec tous les accents possibles
func AccentReformat(letter, wordtofind string) ([]string, []string) {
	rep := []string{}
	for _, let := range wordtofind {
		rep = append(rep, string(let))
	}
	letter = strings.ToLower(letter)
	listletter := []string{letter}
	// On vérifie si la lettre est un e avec accent
	if letter == "e" || letter == "é" || letter == "è" || letter == "ê" || letter == "ë" {
		// On change listletter avec tous les accents possibles de e
		listletter = []string{"e", "é", "è", "ê", "ë"}
	}
	// On vérifie si la lettre est un a avec accent
	if letter == "a" || letter == "à" || letter == "â" || letter == "ä" {
		// On change listletter avec tous les accents possibles de a
		listletter = []string{"a", "à", "â", "ä"}
	}
	// On vérifie si la lettre est un i avec accent
	if letter == "i" || letter == "î" || letter == "ï" {
		// On change listletter avec tous les accents possibles de i
		listletter = []string{"i", "î", "ï"}
	}
	// On vérifie si la lettre est un u avec accent
	if letter == "u" || letter == "ù" || letter == "û" || letter == "ü" {
		// On change listletter avec tous les accents possibles de u
		listletter = []string{"u", "ù", "û", "ü"}
	}
	// On vérifie si la lettre est un c avec accent
	if letter == "c" || letter == "ç" {
		// On change listletter avec tous les accents possibles de c
		listletter = []string{"c", "ç"}
	}
	// On vérifie si la lettre est un y avec accent
	if letter == "y" || letter == "ÿ" {
		// On change listletter avec tous les accents possibles de y
		listletter = []string{"y", "ÿ"}
	}
	return rep, listletter
}
