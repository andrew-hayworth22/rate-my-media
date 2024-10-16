package movies

import "time"

type AppMovie struct {
	Id             int
	Name           string
	Description    string
	ReleaseDate    time.Time
	RuntimeMinutes int
}
