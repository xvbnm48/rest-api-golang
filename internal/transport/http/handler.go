package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Handler - stores pointer to our service
type Handler struct {
	Router *mux.Router
}

// newhandler - returns a pointer to a handler
func NewHandler() *Handler {
	return &Handler{}
}

//setup routes - set up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("setting up routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "i am alive!")
	})
}
