package handlers

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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

type Response struct {
	Id   int16
	Name string
	Hash string
}

// UploadHandler for route upload file
// TODO: 多个文件片段拼接
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Fprintf(w, "Error method")
	} else {
		// save the file either to memory or
		// to a temporary file(if size bigger than maxMemory passed to ParseMultipartForm)
		r.ParseMultipartForm(32 << 20)
		// get the file handle and save to the file system
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		h := sha1.New()
		if _, err := io.Copy(h, file); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%x", h.Sum(nil))
		res := Response{
			Id:   1,
			Name: handler.Filename,
			Hash: string(h.Sum(nil)),
		}
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(res)
		f, err := os.OpenFile("./upload-files/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}
