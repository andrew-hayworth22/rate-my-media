package auth

type DbUser struct {
	Id           int
	Email        string
	Name         string
	DisplayName  string
	PasswordHash string
}

type DbStoreUserRequest struct {
	Email       string
	Name        string
	DisplayName string
	Password    string
}
