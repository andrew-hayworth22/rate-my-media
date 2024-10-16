package auth

type AppUser struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
}
