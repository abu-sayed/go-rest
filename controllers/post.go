package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/abu-sayed/go-rest/models"
	"encoding/json"
	"github.com/abu-sayed/go-rest/services"
	"strconv"
	"log"
)

func GetPost(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		log.Println("String to integer conversion fail.")
	}

	posts := services.GetPost(id)
	json.NewEncoder(res).Encode(posts);
}

func GetPosts(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	posts := models.Post{ID:1, Title:"Hello", Content:"Hello World"}
	json.NewEncoder(res).Encode(posts);
}

func CreatePost(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	posts := models.Post{ID:1, Title:"Hello", Content:"Hello World"}
	json.NewEncoder(res).Encode(posts);
}

func UpdatePost(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	posts := models.Post{ID:1, Title:"Hello", Content:"Hello World"}
	json.NewEncoder(res).Encode(posts);
}

func DeletePost(res http.ResponseWriter, req *http.Request, params httprouter.Params) {
	posts := models.Post{ID:1, Title:"Hello", Content:"Hello World"}
	json.NewEncoder(res).Encode(posts);
}
