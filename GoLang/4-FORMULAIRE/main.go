package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, `
			<!DOCTYPE html>
			<html>
			<head>
				<title>Formulaire</title>
			</head>
			<body>
				<form action="/" method="POST">
					<label for="name">Nom</label>
					<input type="text" name="name" id="name" />
					<input type="submit" value="Envoyer" />
				</form>
			</body>
			</html>
		`)
	} else if r.Method == "POST" {
		r.ParseForm()
		fmt.Fprintf(w, "Bonjour %s !", r.FormValue("name"))
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
