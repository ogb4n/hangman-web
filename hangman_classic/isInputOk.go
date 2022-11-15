package hangman_classic

// Fonction qui va vérifier si le mot ou la lettre est valide
func IsInputOk(letter, randomword, wordtofind string, usedletter *[]string) (string, string) {
	// S'il y a plus d'une lettre de rentrée
	if len([]rune(letter)) > 1 {
		// Si le mot rentré a le même nombre de lettres que le mot cherché
		if len([]rune(randomword)) == len([]rune(letter)) {
			// Si le mot est le bon
			if AccentChecker(randomword) == AccentChecker(letter) {
				return wordtofind, "wordgood"
				// Si le mot n'est pas le bon
			} else {
				return wordtofind, "wordwrong"
			}
			// Si le mot rentré n'a pas le même nombre de lettres que le mot cherché
		} else {
			return wordtofind, "wordinvalid"
		}
	}
	// Si on arrive ici, alors il n'y a qu'une seule lettre de rentrée
	// On vérifie si la lettre a déjà été utilisé
	for _, let := range *usedletter {
		// Si c'est la cas on le retourne
		if let == letter {
			return wordtofind, "usedletter"
		}
	}
	// Si on arrive ici, alors la lettre n'a pas encore été utilisé
	// On ajoute la lettre à la liste de lettres utilisées
	*usedletter = append(*usedletter, letter)
	// Si ce qui a été rentré par l'utilisateur est une lettre
	if AccentChecker(letter) >= "a" && AccentChecker(letter) <= "z" {
		// On récupère la liste d'accent possible pour la lettre rentrée
		rep, listletter := AccentReformat(letter, wordtofind)
		// On boucle dans cette liste
		for _, oneletter := range listletter {
			// On boucle autant de fois que la longueur du mot cherché
			for k := 0; k < len([]rune(randomword)); k++ {
				// Si la lettre rentrée est dans la liste
				if oneletter == string([]rune(randomword)[k]) {
					// On remplace la valeur dans le mot
					rep[k] = oneletter
				}
			}
		}
		word := ""
		// On transforme rep en string appeler word
		for _, let := range rep {
			word += let
		}
		// Si rien a changé entre avant et après les découvertes liés à la lettre, alors c'est un fail
		if word == wordtofind {
			return word, "fail"
		}
		// Sinon c'est bon
		return word, "good"
		// Si ce qui a été rentré par l'utilisateur n'est pas une lettre
	} else {
		return wordtofind, "error"
	}
}
