package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/seedovan19/Agregator/pkg/database"
)

func main() {
	database.Init()

	// делаем миграцию в базу данных (создаем таблицы, если они еще не были созданы)
	database.Migrate()

	PORT := "8083"
	http.HandleFunc("/", Home_page)
	http.HandleFunc("/adding_form", Adding_form)

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Print("Running server on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func Home_page(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/home_page.html")
	tpl.Execute(w, nil)
}

func Adding_form(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/form.html")
	tpl.Execute(w, nil)
}
