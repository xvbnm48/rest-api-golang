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
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("GET")
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
	fmt.Fprint(w, "%+v", comment)

}

//get all comments  - retrieve all comments from the comment service
func (h *Handler) GetAllComment(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		fmt.Fprintf(w, "fail to retrieve all comments")
	}

	fmt.Fprint(w, "%+v", comments)
}
