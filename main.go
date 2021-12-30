package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/seedovan19/Agregator/pkg/database"
)

var err error

func main() {
	database.Open_connection()

	// делаем миграцию в базу данных (создаем таблицы, если они еще не были созданы)
	database.Migrate()

	database.Close()

	PORT := "8083"
	http.HandleFunc("/", home_page)
	http.HandleFunc("/form", warehouse_form)
	http.HandleFunc("/add_warehouse", add_warehouse)

	// http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Print("Running server on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func home_page(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/home_page.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tpl.Execute(w, nil)
}

func warehouse_form(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	tpl.Execute(w, nil)
}

func add_warehouse(w http.ResponseWriter, r *http.Request) {
	adress := r.FormValue("adress")
	name := r.FormValue("name")
	square := r.FormValue("square")
	class := r.FormValue("class")
	// age_of_construction := r.FormValue("age_of_construction")
	shelf_cost := r.FormValue("shelf_cost")
	floor_cost := r.FormValue("floor_cost")
	// description := r.FormValue("description")

	if adress == "" || name == "" || square == "" || class == "" || shelf_cost == "" || floor_cost == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		database.Open_connection()
		database.Close()
	}
}
