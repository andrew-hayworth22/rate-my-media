package movies

import "time"

type AppMovie struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	ReleaseDate    time.Time `json:"release_date"`
	RuntimeMinutes int       `json:"runtime_minutes"`
}
