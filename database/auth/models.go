package auth

type DbUser struct {
	Id           int    `db:"id"`
	Email        string `db:"email"`
	Name         string `db:"name"`
	DisplayName  string `db:"display_name"`
	PasswordHash string `db:"password"`
}

type DbStoreUserRequest struct {
	Email       string
	Name        string
	DisplayName string
	Password    string
}
