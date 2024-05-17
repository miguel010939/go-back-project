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
