package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/xvbnm48/rest-api-golang/internal/comment"
)

//Handler - stores pointer to our service
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - a object to store responses froum our api
type Response struct {
	Message string
	Error   string
}

// newhandler - returns a pointer to a handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

//setup routes - set up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("setting up routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/api/comment", h.GetAllComment).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		
		if err := sendOkResponse(w, Response{Message: "i am Alive"}); err != nil {
			panic(err)
		}
	})
}

// GetComment - retrieving a coment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Fprint(w, "unable to parse uint from id")
	} // mengubah id uint to int

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by id", err)
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
	// fmt.Fprintf(w, "%+v", comment)

}

//get all comments  - retrieve all comments from the comment service
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	

	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "fail to retrieve all comments", err)
	}

	if err := sendOkResponse(w ,comments); err != nil {
		panic(err)
	}
}

// post comment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusOK)

	var comment comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "failed to decode JSON BODY", err)
	}

	comment, err := h.Service.PostComment(comment)

	if err != nil {
		sendErrorResponse(w, "failed to post new commnet", err)
	}

	// fmt.Fprintf(w, "%+v", comment)
	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// updateComment - update comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	
	var comment comment.Comment

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "failed to decode JSON BODY", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "failed to parse uint from ID")
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)

	if err != nil {
		sendErrorResponse(w, "failed to update comment", err)
		return
	}
	// fmt.Fprintf(w, "%+v", comment)
	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

//delete comment - deletes comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "failed to parse uint from ID", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "failed to delete comment by comment id", err)
		return
	}

	// mengunakan helper response ok
	if err = sendOkResponse(w, Response{Message: "success deleted comment"}); err != nil {
		panic(err)
	}

	// if err := json.NewEncoder(w).Encode(Response{Message: "success deleted comment"}); err != nil {
	// 	panic(err)
	// }

	// fmt.Fprint(w, "success deleted comment")

}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}

}
