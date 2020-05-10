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
	http.HandleFunc("/api/user", handlers.GetUserHandler)
	http.HandleFunc("/api/upload", handlers.UploadHandler)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
