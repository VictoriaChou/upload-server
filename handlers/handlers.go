package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictoriaChou/upload-server/database"
)

// GetUserHandler for route get user
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := database.User{
		Name: "Scarlett",
	}
	err := user.Get(database.MyDB)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(user)
}
