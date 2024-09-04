package main

import (
	"fmt"
	"net/http"

	"github.com/andrew-hayworth22/rate-my-media/app"
	"github.com/andrew-hayworth22/rate-my-media/business"
	"github.com/andrew-hayworth22/rate-my-media/database"
)

func main() {
	db := database.NewDatabase()
	bus := business.NewBusiness(db)
	app := app.NewApp(bus)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/user", app.PostUser)

	server := http.Server{
		Addr:    app.PortAddress,
		Handler: mux,
	}

	fmt.Println("Starting server...")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("error: %s", err)
	}
}
