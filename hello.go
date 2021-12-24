// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
//  "net/http"
// )

// func main() {
// PORT := "127.0.0.1:8080"
// log.Print("Running server on " + PORT)
// http.HandleFunc("/complete/", CompleteTaskFunc)
// http.HandleFunc("/delete/", DeleteTaskFunc)
// http.HandleFunc("/deleted/", ShowTrashTaskFunc)
// http.HandleFunc("/trash/", TrashTaskFunc)
// http.HandleFunc("/edit/", EditTaskFunc)
// http.HandleFunc("/completed/", ShowCompleteTasksFunc)
// http.HandleFunc("/restore/", RestoreTaskFunc)
// http.HandleFunc("/add/", AddTaskFunc)
// http.HandleFunc("/update/", UpdateTaskFunc)
// http.HandleFunc("/search/", SearchTaskFunc)
// http.HandleFunc("/login", GetLogin)
// http.HandleFunc("/register", PostRegister)
// http.HandleFunc("/admin", HandleAdmin)
// http.HandleFunc("/add_user", PostAddUser)
// http.HandleFunc("/change", PostChange)
// http.HandleFunc("/logout", HandleLogout)
// http.HandleFunc("/", ShowAllTasksFunc)

// http.Handle("/static/", http.FileServer(http.Dir("public")))
// log.Print("running on port 8080")

// log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(r.URL.Path))
// }

// func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func EditTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request) {}

// func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func AddTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func SearchTaskFunc(w http.ResponseWriter, r *http.Request) {}

// func GetLogin(w http.ResponseWriter, r *http.Request) {}

// func PostRegister(w http.ResponseWriter, r *http.Request) {}

// func HandleAdmin(w http.ResponseWriter, r *http.Request) {}

// func PostAddUser(w http.ResponseWriter, r *http.Request) {}

// func PostChange(w http.ResponseWriter, r *http.Request) {}

// func HandleLogout(w http.ResponseWriter, r *http.Request) {}

// func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
// 	var message string
// 	if r.Method == "GET" {
// 		message = "all pending tasks GET"
// 	} else {
// 		message = "all pending tasks POST"
// 	}
// 	w.Write([]byte(message))
// }