package main

import (
	"log"
	"net/http"
	"rest-go-demo/controllers"
	"rest-go-demo/database"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	initDB()
	server()
}

func server() {
	log.Println(("HTTP Server started!"))
	router := mux.NewRouter().StrictSlash(true)
	initializeHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initializeHandlers(router *mux.Router) {
	router.HandleFunc("/create", controllers.CreateArticle).Methods("POST")
	router.HandleFunc("/article/{id}", controllers.GetArticleById).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.UpdateArticleById).Methods("PUT")
	router.HandleFunc("/article/{id}", controllers.DeletePersonById).Methods("DELETE")
}

func initDB() {
	config := database.Config{
		ServerName: "localhost:1444",
		User:       "root",
		Password:   "root",
		DB:         "mydb",
	}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)

	if err != nil {
		panic(err.Error())
	}
}
