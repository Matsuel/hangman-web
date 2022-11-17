package main

import (
	"fmt"
	hw "hangmanweb"
	hc "hangmanweb/hangman-classic"
	"html/template"
	"net/http"
	"os"
)

var data hw.Hangman

func main() {
	data.MotTab, data.Mot, data.Motstr = hw.Initword(os.Args[len(os.Args)-1])
	data.Attempts = 10
	fmt.Println("Starting server on port 8080")
	http.HandleFunc("/", HandlerUser)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/hangman", Handler)
	http.ListenAndServe(":8080", nil)
	return
}

func HandlerUser(w http.ResponseWriter, r *http.Request) {
	// switch r.Method {
	// case "POST": // Gestion d'erreur
	// 	if err := r.ParseForm(); err != nil {
	// 		return
	// 	}
	// }
	// fmt.Println("HEEEEEEEEEEEEEEEERE")
	// // Récupérez votre valeur
	// username := r.FormValue("inputBox")
	// //username := r.Form.Get("inputBox")
	// fmt.Println(username)
	// //fmt.Println(variable)
	tmpl := template.Must(template.ParseFiles("./static/index.html"))
	// data.Username = username
	// fmt.Println(data.Username)
	//data.LettersUsed = append(data.LettersUsed, variable)
	//fmt.Println(data)
	tmpl.Execute(w, nil)
	return
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST": // Gestion d'erreur
		if err := r.ParseForm(); err != nil {
			return
		}
	}
	// Récupérez votre valeur
	variable := r.FormValue("input")
	fmt.Println("heere" + variable)
	//fmt.Println(variable)
	tmpl := template.Must(template.ParseFiles("./static/play.html"))
	if data.Username == "" {
		data.Username = variable
		fmt.Println(data.Username)
	} else {
		data.Letter = variable
		fmt.Println(data.Username)
	}
	//data.LettersUsed = append(data.LettersUsed, variable)
	fmt.Println(data.Motstr)
	HangmanWeb()
	tmpl.Execute(w, data)
	return
}

func HangmanWeb() {
	Motstr, State := hc.IsInputOk(data.Letter, data.Mot, data.Motstr, &data.LettersUsed)
	data.Motstr = Motstr
	if !(LetterPresent(data.Letter)) {
		data.LettersUsed = append(data.LettersUsed, data.Letter)
	}
	if State == "fail" {
		data.Attempts--
	}
	if data.Mot == data.Motstr {
		fmt.Println("gg")
		return
	}
}

func LetterPresent(letter string) bool {
	for _, ch := range data.LettersUsed {
		if data.Letter == ch {
			return true
		}
	}
	return false
}

//Pour la carte d'identité faut rediriger depuis la fonction handler vers /http.HandleFunc("/hangman", Handler)
