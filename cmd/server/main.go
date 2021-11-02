package main

import (
	"fmt"
	"net/http"

	"github.com/xvbnm48/rest-api-golang/internal/database"
	transportHTTP "github.com/xvbnm48/rest-api-golang/internal/transport/http"
)

//APP - the sturct which contains things like pointers
type App struct{}

// RUn - set up our application
func (app *App) Run() error {
	fmt.Println("setting up our app")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("failed  to set up server")
		return err
	}
	return nil
}

func main() {
	fmt.Println("GO Rest API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("error starting up our REST API")
		fmt.Println(err)
	}
}
