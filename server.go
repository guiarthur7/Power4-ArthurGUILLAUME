package main

import (
	"html/template"
	"log"
	"net/http"
	power4 "power4/game"
	"strconv"
)

type Player struct {
	Name1 string
	Name2 string
}

type GameData struct {
	Player  Player
	Game    *power4.Game
	Gagnant int // 0 = partie en cours, 1 = joueur 1 gagne, 2 = joueur 2 gagne, 3 = match nul
}

var jeuEnCours *power4.Game
var joueurs Player

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func JeuFacile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		joueurs = Player{
			Name1: r.FormValue("name1"),
			Name2: r.FormValue("name2"),
		}
		jeuEnCours = power4.InitGameFacile()

		data := GameData{
			Player:  joueurs,
			Game:    jeuEnCours,
			Gagnant: 0, // Nouvelle partie
		}

		tmpl, err := template.ParseFiles("./game/powerfacile.html", "./template/info_joueurs.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func JeuNormal(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		joueurs = Player{
			Name1: r.FormValue("name1"),
			Name2: r.FormValue("name2"),
		}
		jeuEnCours = power4.InitGameMoyen()

		data := GameData{
			Player:  joueurs,
			Game:    jeuEnCours,
			Gagnant: 0, // ← AJOUTE ÇA !
		}

		tmpl, err := template.ParseFiles("./game/powernormal.html", "./template/info_joueurs.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func JeuHard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		joueurs = Player{
			Name1: r.FormValue("name1"),
			Name2: r.FormValue("name2"),
		}
		jeuEnCours = power4.InitGameHard()

		data := GameData{
			Player:  joueurs,
			Game:    jeuEnCours,
			Gagnant: 0, // ← AJOUTE JUSTE CETTE LIGNE !
		}

		tmpl, err := template.ParseFiles("./game/powerhard.html", "./template/info_joueurs.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, data)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func JouerCoup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		colStr := r.FormValue("colonne")

		col, err := strconv.Atoi(colStr)
		if err != nil || col < 0 {
			http.Error(w, "Colonne invalide", http.StatusBadRequest)
			return
		}

		joueurActuel := jeuEnCours.Tour
		gagnant := 0

		if power4.PlacerJeton(jeuEnCours, col) {

			if power4.VerifVictoire(jeuEnCours, joueurActuel) {
				gagnant = joueurActuel
			} else if power4.GrillePleineVerif(jeuEnCours) {
				gagnant = 3
			}
		}

		data := GameData{
			Player:  joueurs,
			Game:    jeuEnCours,
			Gagnant: gagnant,
		}

		mode := r.FormValue("mode")
		var tmpl *template.Template

		if mode == "powerfacile" {
			tmpl, err = template.ParseFiles("./game/powerfacile.html", "./template/info_joueurs.html")
		} else if mode == "powernormal" {
			tmpl, err = template.ParseFiles("./game/powernormal.html", "./template/info_joueurs.html")
		} else if mode == "powerhard" {
			tmpl, err = template.ParseFiles("./game/powerhard.html", "./template/info_joueurs.html")
		}

		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, data)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/powerfacile.html", JeuFacile)
	http.HandleFunc("/powernormal.html", JeuNormal)
	http.HandleFunc("/powerhard.html", JeuHard)
	http.HandleFunc("/jouer", JouerCoup)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	imFS := http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", imFS))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
