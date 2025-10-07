package main

import (
	"html/template"
	"log"
	"net/http"
)

type Player struct {
	Name1 string
	Name2 string
}

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func JeuFacile(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		joueur := Player{
			Name1: r.FormValue("name1"),
			Name2: r.FormValue("name2"),
		}
		tmpl, err := template.ParseFiles("./game/powerfacile.html", "./template/info_joueurs.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, joueur)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func JeuNormal(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		joueur := Player{
			Name1: r.FormValue("name1"),
			Name2: r.FormValue("name2"),
		}
		tmpl, err := template.ParseFiles("./game/powernormal.html", "./template/info_joueurs.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, joueur)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func JeuHard(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		joueur := Player{
			Name1: r.FormValue("name1"),
			Name2: r.FormValue("name2"),
		}
		tmpl, err := template.ParseFiles("./game/powerhard.html", "./template/info_joueurs.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.Execute(w, joueur)
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/powerfacile.html", JeuFacile)
	http.HandleFunc("/powernormal.html", JeuNormal)
	http.HandleFunc("/powerhard.html", JeuHard)

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	imFS := http.FileServer(http.Dir("src/"))
	http.Handle("/src/", http.StripPrefix("/src/", imFS))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
