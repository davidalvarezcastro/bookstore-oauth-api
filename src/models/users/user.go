package users

// User stores user info
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// UserLoginRequest stores login request
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
