package main

import (
	"log"
	"net/http"

	"github.com/VictoriaChou/upload-server/database"
	"github.com/VictoriaChou/upload-server/handlers"
)

func main() {
	database.Init()
	defer database.MyDB.Close()
	http.HandleFunc("/user", handlers.GetUserHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:3000", nil))
}
