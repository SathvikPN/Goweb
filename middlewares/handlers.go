package middlewares

import (
	"encoding/json"
	"fmt"
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
	post.CreatedAt = services.GetCurrentTime()
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

// GetAllUser will return all the users
func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	// get all the posts in the db
	posts, err := services.GetAllPosts()

	if err != nil {
		log.Fatalf("Unable to get all posts. %v", err)
	}

	if len(posts) == 0 {
		posts = []models.Post{}
	}
	// send all the users as response
	json.NewEncoder(w).Encode(posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("unable to convert string into int, error:", err)
	}

	var post models.Post

	// decode JSON Request to post struct
	err = json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := services.UpdatePost(int64(id), post)
	msg := "post update OK."
	if updatedRows == 0 {
		msg = "no updates."
	}
	msg = msg + "Total rows/record affected: " + fmt.Sprint(updatedRows)

	// format the response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("unable to convert string into int, error:", err)
	}
	deletedRows := services.DeletePost(int64(id))
	msg := "post deleted successfully. "
	if deletedRows == 0 {
		msg = "Non-existent post. Delete OK."
	}
	msg = msg + "Total rows/record affected: " + fmt.Sprint(deletedRows)
	res := response{
		ID:      int64(id),
		Message: msg,
	}
	json.NewEncoder(w).Encode(res)
}
