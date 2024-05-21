package repositories

import (
	"database/sql"
	"main.go/models"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) UserSignUp(signUpForm *models.UserSignUpForm) (string, error) {
	var token string

	return token, nil
}
func (r *UserRepo) UserLogIn(signUpForm *models.UserLogInForm) (string, error) {
	var token string

	return token, nil
}

func (r *UserRepo) UserLogOut(token string) error {

	return nil
}

func (r *UserRepo) GetAllUsers() ([]models.User, error) {
	rows, err := r.db.Query("SELECT id, username FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
