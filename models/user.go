package models

type UserRepresentation struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

type UserSignUpForm struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogInForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
