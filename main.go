package main

import (
	"html/template"
	"net/http"
	"log"
	"github.com/BrandonEchols/02_SpellBookApp/models"
	"fmt"
	"os"
)

var tpl *template.Template
var SBM map[string]models.SpellBook
var port int
type pageData struct {
	Title string
	SpellBook models.SpellBook
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	SBM = models.GetSpellBookMap()
	port = os.Getenv("PORT")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/spellpage", spellpage)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if port != 0 {
		http.ListenAndServe(":" + port, nil)
		fmt.Println("Server listening on port", port)
	} else {
		http.ListenAndServe(":8080", nil)
		fmt.Println("Server listening on port 8080")
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "Index"}

	err := tpl.ExecuteTemplate(w, "index.gohtml", pd)                                	//Go execute this template (returns an error value
	if err != nil {
		log.Println(err)                                                        	//Print the error to server logs
		http.Error(w, "Internal server error", http.StatusInternalServerError)        	//also print some extra information to the server logs
	}
	//fmt.Println("Path: " + req.URL.Path) <<For debugging
}
func about(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "About"}

	err := tpl.ExecuteTemplate(w, "about.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
func spellpage(w http.ResponseWriter, req *http.Request) {
	pd := pageData{Title: "Spellpage", SpellBook: SBM["Cleric"]}

	err := tpl.ExecuteTemplate(w, "spellpage.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
