package main

import (
	"rest-go-demo/database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
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
