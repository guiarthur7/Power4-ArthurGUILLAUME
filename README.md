# 🎮 Power 4 Web - Arthur GUILLAUME

## 📋 Description

Power 4 Web est un jeu de Puissance 4 multijoueur local développé en Go avec une interface web. Deux joueurs peuvent s'affronter sur le même ordinateur à travers une interface graphique élégante.

## ✨ Fonctionnalités

### 🎯 Fonctionnalités Principales

- **3 modes de difficulté** :

  - 🟢 **Facile** : Grille 6x7 (classique)
  - 🟡 **Normal** : Grille 6x9 
  - 🔴 **Difficile** : Grille 7x8

### Installation

1. **Cloner le repository**

```bash
git clone https://github.com/votre-username/Power4-ArthurGUILLAUME.git
cd Power4-ArthurGUILLAUME
```

2. **Lancer le serveur Go**

```bash
go run server.go
```

3. **Ouvrir le navigateur**
```
http://localhost:8080
```

## 🎮 Comment Jouer

1. **Page d'accueil** : Entrez les pseudos des deux joueurs
2. **Choix du mode** : Sélectionnez la difficulté souhaitée
3. **Partie** : Cliquez sur les numéros de colonnes pour placer vos jetons
4. **Victoire** : Le premier à aligner 4 jetons gagne !
5. **Nouvelle partie** : Cliquez sur "🔄 Nouvelle partie" pour rejouer

## 🛠️ Technologies Utilisées

- **Backend** : Go (Golang)
- **Frontend** : HTML5, CSS3 
- **Template** : Go templates (`html/template`)
- **Serveur** : `net/http` (package standard Go)

## 📝 Règles du Jeu

- Chaque joueur joue à tour de rôle
- Les jetons tombent par gravité (colonne la plus basse disponible)
- **Victoire** : Aligner 4 jetons horizontalement, verticalement ou en diagonale
- **Match nul** : Grille complètement remplie sans alignement de 4 jetonsS

## 🌟 Fonctionnalités Avancées

### Génération Dynamique de la Grille
- Utilisation de boucles Go templates pour générer automatiquement les cellules
- Adaptation automatique selon le mode de jeu

### Gestion de Session
- Conservation des pseudos entre les parties
- Pas besoin de ressaisir les noms pour rejouer

## 👨‍💻 Auteur

**Arthur GUILLAUME** - Ynov Campus Lyon B1 Informatique

---

🎮 **Bon jeu !** 🎮