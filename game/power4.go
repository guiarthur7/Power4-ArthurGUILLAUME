package power4

type Game struct {
	Tour  int
	Board [][]string
}

func InitGameFacile() *Game {
	jeu := &Game{
		Tour: 1,
		Board: [][]string{
			{"", "", "", "", "", "", ""},
			{"", "", "", "", "", "", ""},
			{"", "", "", "", "", "", ""},
			{"", "", "", "", "", "", ""},
			{"", "", "", "", "", "", ""},
			{"", "", "", "", "", "", ""},
		},
	}
	return jeu
}

func InitGameMoyen() *Game {
	jeu := &Game{
		Tour: 1,
		Board: [][]string{
			{"", "", "", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", "", "", ""},
		},
	}
	return jeu
}

func InitGameHard() *Game {
	jeu := &Game{
		Tour: 1,
		Board: [][]string{
			{"", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", ""},
			{"", "", "", "", "", "", "", ""},
		},
	}
	return jeu
}

func PlacerJeton(jeu *Game, col int) bool {
	if ColonnePleine(jeu, col) == true {
		return false
	}

	for ligne := len(jeu.Board) - 1; ligne >= 0; ligne-- {
		if jeu.Board[ligne][col] == "" {
			if jeu.Tour == 1 {
				jeu.Board[ligne][col] = "X"
			} else {
				jeu.Board[ligne][col] = "O"
			}
			SwitchPlayer(jeu)
			return true
		}
	}

	return false
}

func SwitchPlayer(jeu *Game) {
	if jeu.Tour == 1 {
		jeu.Tour = 2
	} else if jeu.Tour == 2 {
		jeu.Tour = 1
	}
}

func GrillePleineVerif(jeu *Game) bool {
	for col := 0; col < len(jeu.Board[0]); col++ {
		if jeu.Board[0][col] == "" {
			return false
		}
	}
	return true
}

func ColonnePleine(jeu *Game, col int) bool {
	if jeu.Board[0][col] == "" {
		return false
	}
	return true
}

func VerifVictoire(jeu *Game, joueur int) bool {

	// Vérification horizontale
	for ligne := 0; ligne < len(jeu.Board); ligne++ {
		for col := 0; col < len(jeu.Board[0])-3; col++ {
			if jeu.Board[ligne][col] != "" &&
				jeu.Board[ligne][col] == jeu.Board[ligne][col+1] &&
				jeu.Board[ligne][col] == jeu.Board[ligne][col+2] &&
				jeu.Board[ligne][col] == jeu.Board[ligne][col+3] {
				return true
			}
		}
	}

	// Vérification verticale
	for ligne := 0; ligne < len(jeu.Board)-3; ligne++ {
		for col := 0; col < len(jeu.Board[0]); col++ {
			if jeu.Board[ligne][col] != "" &&
				jeu.Board[ligne][col] == jeu.Board[ligne+1][col] &&
				jeu.Board[ligne][col] == jeu.Board[ligne+2][col] &&
				jeu.Board[ligne][col] == jeu.Board[ligne+3][col] {
				return true
			}
		}
	}

	// Vérification diago (haut-gauche vers bas-droite)
	for ligne := 0; ligne < len(jeu.Board)-3; ligne++ {
		for col := 0; col < len(jeu.Board[0])-3; col++ {
			if jeu.Board[ligne][col] != "" &&
				jeu.Board[ligne][col] == jeu.Board[ligne+1][col+1] &&
				jeu.Board[ligne][col] == jeu.Board[ligne+2][col+2] &&
				jeu.Board[ligne][col] == jeu.Board[ligne+3][col+3] {
				return true
			}
		}
	}

	// Vérification diago (bas-gauche vers haut-droite)
	for ligne := 3; ligne < len(jeu.Board); ligne++ {
		for col := 0; col < len(jeu.Board[0])-3; col++ {
			if jeu.Board[ligne][col] != "" &&
				jeu.Board[ligne][col] == jeu.Board[ligne-1][col+1] &&
				jeu.Board[ligne][col] == jeu.Board[ligne-2][col+2] &&
				jeu.Board[ligne][col] == jeu.Board[ligne-3][col+3] {
				return true
			}
		}
	}

	return false
}
