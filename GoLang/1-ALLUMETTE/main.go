package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Choisiser le nombre d'allumette (4 minimum) : ")
	fmt.Scan(&n)

	if n < 4 {
		fmt.Println("Le nombre d'allumettes doit être au moins 4.")
		return
	}

	player := 1
	for n > 0 {
		fmt.Printf("Joueur %d choisiser votre nombre d'allumette (1 à 3) ? ", player)
		var take int
		fmt.Scan(&take)

		if take < 1 || take > 3 {
			fmt.Println("Vous devez prendre entre 1 et 3 allumettes.")
			continue
		}

		n -= take
		if n <= 0 {
			fmt.Printf("Joueur %d vous avez perdu.\n", player)
			return
		}

		player = 3 - player
	}

	fmt.Println("Votre nombre d'allumette est incorrect. ( négatif )")
}
