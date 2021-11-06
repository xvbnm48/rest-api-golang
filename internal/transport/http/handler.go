package http

import (
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
		fmt.Fprintf(w, "i am alive!")
	})
}

// GetComment - retrieving a coment by id
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Fprint(w, "unable to parse uint from id")
	} // mengubah id uint to int

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		fmt.Fprint(w, "Error retrieving comment by id")
	}
	fmt.Fprintf(w, "%+v", comment)

}

//get all comments  - retrieve all comments from the comment service
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "fail to retrieve all comments")
	}

	fmt.Fprintf(w, "%+v", comments)
}

// post comment - adds a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.PostComment(comment.Comment{
		SLug: "/",
	})

	if err != nil {
		fmt.Fprintf(w, "failed to post new commnet")
	}

	fmt.Fprintf(w, "%+v", comment)
}

// updateComment - update comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	comment, err := h.Service.UpdateComment(1, comment.Comment{
		SLug: "/new",
	})
	if err != nil {
		fmt.Fprintf(w, "failed to update comment")
	}
	fmt.Fprintf(w, "%+v", comment)
}

//delete comment - deletes comment by id
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "failed to parse uint from ID")
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		fmt.Fprint(w, "failed to delete comment by comment id")
	}

	fmt.Fprint(w, "success deleted comment")

}
