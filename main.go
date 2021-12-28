package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/seedovan19/Agregator/pkg/database"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {
	database.Init()

	// делаем миграцию в базу данных (создаем таблицы, если они еще не были созданы)
	database.Migrate()

	PORT := "127.0.0.1:8083"
	http.HandleFunc("/", CompleteTaskFunc)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Print("Running server on " + PORT)
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}
