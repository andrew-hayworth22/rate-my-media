package app

import (
	"os"

	"github.com/andrew-hayworth22/rate-my-media/business"
	"github.com/joho/godotenv"
)

type App struct {
	PortAddress string
	jwtSecret   string
	bus         business.BusinessLayer
}

func NewApp(bus business.BusinessLayer) *App {
	godotenv.Load()

	return &App{
		PortAddress: os.Getenv("PORT_ADDRESS"),
		jwtSecret:   os.Getenv("JWT_SECRET"),
		bus:         bus,
	}
}
