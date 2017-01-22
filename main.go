package main

import (
	"html/template"
	"net/http"
	"log"
	"github.com/BrandonEchols/02_SpellBookApp/models"
	"fmt"
	"os"
	"strings"
)

var tpl *template.Template
var SBM map[string]models.SpellBook
var port string
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
	http.HandleFunc("/spellpage/", spellpage)

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	if port != "" {
		fmt.Println("Server listening on port", port)
		http.ListenAndServe(":" + port, nil)
	} else {
		fmt.Println("Server listening on port 8080")
		http.ListenAndServe(":8080", nil)
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
	var pd pageData
	if strings.Contains(req.URL.Path, models.CLERIC){
		pd = pageData{Title: "Cleric Spellpage", SpellBook: SBM[models.CLERIC]}
	} else if strings.Contains(req.URL.Path, models.BARD){
		pd = pageData{Title: "Bard Spellpage", SpellBook: SBM[models.BARD]}
	} else if strings.Contains(req.URL.Path, models.DRUID){
		pd = pageData{Title: "Druid Spellpage", SpellBook: SBM[models.DRUID]}
	} else if strings.Contains(req.URL.Path, models.WIZARD){
		pd = pageData{Title: "Wizard Spellpage", SpellBook: SBM[models.WIZARD]}
	} else if strings.Contains(req.URL.Path, models.RANGER){
		pd = pageData{Title: "Ranger Spellpage", SpellBook: SBM[models.RANGER]}
	} else if strings.Contains(req.URL.Path, models.PALADIN){
		pd = pageData{Title: "Paladin Spellpage", SpellBook: SBM[models.PALADIN]}
	}

	err := tpl.ExecuteTemplate(w, "spellpage.gohtml", pd)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
