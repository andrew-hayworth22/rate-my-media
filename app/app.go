package app

import (
	"os"
	"strings"

	"github.com/andrew-hayworth22/rate-my-media/business"
	"github.com/joho/godotenv"
)

type App struct {
	PortAddress string
	jwtSecret   string
	bus         business.BusLayer
}

func NewApp(bus business.BusLayer) *App {
	godotenv.Load()

	portAddress := os.Getenv("PORT_ADDRESS")
	portAddress = strings.Trim(portAddress, "\"")

	return &App{
		PortAddress: portAddress,
		jwtSecret:   os.Getenv("JWT_SECRET"),
		bus:         bus,
	}
}
