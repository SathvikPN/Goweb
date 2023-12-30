package middlewares

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SathvikPN/Goweb/models"
	"github.com/SathvikPN/Goweb/services"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres golang driver // NOTE: use only init functions and nothing else (_)
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// createPost create a post in postgres DB
func CreatePost(w http.ResponseWriter, r *http.Request) {
	// set W.Header()

	// create an empty post of type models.Post
	var post models.Post

	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Fatal("Unable to decode request body", err)
	}

	insertID := services.InsertPost(post)

	// format response object
	res := response{
		ID:      insertID,
		Message: "post created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// return single post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("unable to convert string into int, error:", err)
	}

	post, err := services.GetPost(int64(id))
	if err != nil {
		log.Fatal("unable to get post, error:", err)
	}

	json.NewEncoder(w).Encode(post)
}
