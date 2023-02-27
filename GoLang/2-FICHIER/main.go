package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var choix int

	for {
		fmt.Println("Menu:")
		fmt.Println("1. Récupérer tout le texte contenu dans un fichier")
		fmt.Println("2. Ajouter du texte dans un fichier")
		fmt.Println("3. Supprimer tout le contenu d'un fichier")
		fmt.Println("4. Remplacer le contenu d'un fichier par du texte donné par l'utilisateur")
		fmt.Println("5. Quitter")

		fmt.Print("Entrez votre choix: ")
		fmt.Scan(&choix)

		switch choix {
		case 1:
			lireFichier()
		case 2:
			ajouterFichier()
		case 3:
			supprimerFichier()
		case 4:
			remplacerFichier()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Choix invalide")
		}
	}
}

func lireFichier() {
	var nomFichier string

	fmt.Print("Entrez le nom du fichier: ")
	fmt.Scan(&nomFichier)

	contenu, err := ioutil.ReadFile(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier:", err)
		return
	}

	fmt.Println("Contenu du fichier:")
	fmt.Println(string(contenu))
}

func ajouterFichier() {
	var nomFichier string
	var texte string

	fmt.Print("Entrez le nom du fichier: ")
	fmt.Scan(&nomFichier)

	fmt.Print("Entrez le texte à ajouter: ")
	fmt.Scan(&texte)

	fichier, err := os.OpenFile(nomFichier, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier:", err)
		return
	}

	defer fichier.Close()

	_, err = fmt.Fprintln(fichier, texte)
	if err != nil {
		fmt.Println("Erreur lors de l'écriture dans le fichier:", err)
		return
	}

	fmt.Println("Texte ajouté avec succès dans le fichier")
}

func supprimerFichier() {
	var nomFichier string

	fmt.Print("Entrez le nom du fichier: ")
	fmt.Scan(&nomFichier)

	err := os.Remove(nomFichier)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du fichier:", err)
		return
	}

	fmt.Println("Fichier supprimé avec succès")
}

func remplacerFichier() {
	var nomFichier string
	var texte string

	fmt.Print("Entrez le nom du fichier: ")
	fmt.Scan(&nomFichier)

	fmt.Print("Entrez le texte de remplacement: ")
	fmt.Scan(&texte)

	err := ioutil.WriteFile(nomFichier, []byte(texte), 0644)
	if err != nil {
		fmt.Println("Erreur lors de la modification du fichier:", err)
		return
	}

	fmt.Println("Contenu du fichier remplacé avec succès")
}
