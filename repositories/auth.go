package repositories

import (
	"database/sql"
	"errors"
	"fmt"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (r *AuthRepo) GetID(token string) (int, error) {
	var id int
	// TODO clean token string
	query := fmt.Sprintf("SELECT userx FROM sessions WHERE token='%s'", token)
	row := r.db.QueryRow(query)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}
func (r *AuthRepo) GetToken(userId int) (string, error) {
	if r.hasTokenAlready(userId) {
		return "", errors.New("409") //TODO error 409
	}
	var token = generateRandomToken(userId)
	_, err := r.db.Exec("INSERT INTO sessions (userx, token) VALUES ($1, $2)", userId, token)
	if err != nil {
		return "", err
	}
	return token, nil
}
func (r *AuthRepo) hasTokenAlready(userId int) bool {
	query := fmt.Sprintf("SELECT userx FROM sessions WHERE userx=%d", userId)
	row := r.db.QueryRow(query)
	err := row.Scan()
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return true
}

func generateRandomToken(userIdSeed int) string {
	return ""
}
