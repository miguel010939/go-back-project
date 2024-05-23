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

func (r *UserRepo) UserSignUp(signUpForm *models.UserSignUpForm, auth *AuthRepo) (string, error) {
	if !signUpForm.IsValid() {
		return "", InvalidInput
	}
	// TODO Hash the passwords!!! This makes CleanPassword() redundant. I don't think im gonna bother salting them tho, we'll see
	insertQuery := `INSERT INTO users (username, email, password)
					VALUES ($1, $2, $3) 
					RETURNING id`
	var id int
	var token string
	err := r.db.QueryRow(insertQuery, signUpForm.Username, signUpForm.Email, signUpForm.Password).Scan(&id)
	if err != nil {
		return "", SomethingWentWrong
	}
	token, err2 := auth.GetToken(id)
	if err2 != nil {
		return "", err2
	}
	return token, nil
}
func (r *UserRepo) UserLogIn(logInForm *models.UserLogInForm, auth *AuthRepo) (string, error) {
	if !logInForm.IsValid() {
		return "", InvalidInput
	}
	// TODO Hash the passwords!!! This makes CleanPassword() redundant. I don't think im gonna bother salting them tho, we'll see
	selectQuery := `SELECT id FROM users 
					WHERE username = $1 AND password = $2`
	var id int
	var token string
	err2 := r.db.QueryRow(selectQuery, logInForm.Username, logInForm.Password).Scan(&id)
	if err2 != nil {
		return "", SomethingWentWrong
	}
	token, err3 := auth.GetToken(id)
	if err3 != nil {
		return "", err3
	}
	return token, nil
}

func (r *UserRepo) UserLogOut(token string) error {
	deleteQuery := `DELETE FROM sessions WHERE token = $1`
	result, err1 := r.db.Exec(deleteQuery, token)
	if err1 != nil {
		return SomethingWentWrong
	}
	rows, err2 := result.RowsAffected()
	if err2 != nil {
		return SomethingWentWrong
	}
	if rows == 0 {
		return NoPermission
	}
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
