package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/seedovan19/Agregator/pkg/database"
)

var posts = []database.Warehouse{}

func main() {
	database.Open_connection()

	// делаем миграцию в базу данных (создаем таблицы, если они еще не были созданы)
	database.Migrate()

	database.Close()

	PORT := "8083"
	http.HandleFunc("/", home_page)
	http.HandleFunc("/form", warehouse_form_page)
	http.HandleFunc("/add_warehouse", add_warehouse)

	log.Print("Running server on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func home_page(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/home_page.gohtml")
	if err != nil {
		log.Print(err.Error())
	}

	database.Open_connection()

	posts, err = database.Show_warehouse_records()
	if err != nil {
		log.Print(err.Error())
	}
	database.Close()
	tpl.Execute(w, posts)
}

func warehouse_form_page(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/form.gohtml")
	if err != nil {
		log.Print(err.Error())
	}
	tpl.Execute(w, nil)
}

func add_warehouse(w http.ResponseWriter, r *http.Request) {

	adress := r.FormValue("adress")
	name := r.FormValue("name")
	square, _ := strconv.ParseInt(r.FormValue("square")[0:], 10, 64)
	class := r.FormValue("class")
	age_of_construction, _ := time.Parse(r.FormValue("age_of_construction"), "2019-07-10")
	shelf_cost, _ := strconv.ParseInt(r.FormValue("shelf_cost")[0:], 10, 64)
	floor_cost, _ := strconv.ParseInt(r.FormValue("floor_cost")[0:], 10, 64)
	image := r.FormValue("image")
	description := r.FormValue("description")

	if adress == "" || name == "" || class == "" {
		fmt.Fprintf(w, "Не все данные заполнены")
	} else {
		database.Open_connection()

		warehouse := database.Warehouse{
			Name:               name,
			Square:             square,
			Adress:             adress,
			Shelf_storage_cost: shelf_cost,
			Floor_storage_cost: floor_cost,
			Description:        description,
			Image:              image,
		}

		building := database.Building{
			Warehouse_class:      class,
			Year_of_construction: age_of_construction,
		}

		database.Add_warehouse_record(warehouse, building)

		database.Close()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
