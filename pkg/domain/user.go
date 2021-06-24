package domain

type User struct {
	// TODO: Implement User model
	Id       int64
	Username string
	Password string
}

func NewUser(username, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}
