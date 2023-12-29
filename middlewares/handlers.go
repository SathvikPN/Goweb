package middlewares

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/SathvikPN/Goweb/models"
	"github.com/SathvikPN/Goweb/services"
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
