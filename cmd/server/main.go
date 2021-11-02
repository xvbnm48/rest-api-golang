package main

import "fmt"

//APP - the sturct which contains things like pointers
type App struct{}

// RUn - set up our application
func (app *App) Run() error {
	fmt.Println("setting up our app")
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
