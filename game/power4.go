package power4

rows := 6

col := 7

type Jeu stuct {
	Grille [][]uint8
	Player uint8
}

func Switch(game *Jeu) {
	if game.Player == 1 {
		game.Player = 2
	} 
	if game.Player == 2 {
		game.Player = 1
	}
}