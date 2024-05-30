package repositories

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"main.go/models"
	"math/rand"
	"time"
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{db}
}

func (r *AuthRepo) GetID(token string) (int, error) {
	var id int
	token, err1 := models.CleanToken(token)
	if err1 != nil {
		return -1, InvalidInput
	}
	row := r.db.QueryRow("SELECT userx FROM sessions WHERE token=$1", token)
	err := row.Scan(&id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, NoPermission
		}
		return -1, SomethingWentWrong
	}
	return id, nil
}
func (r *AuthRepo) GetToken(userId int) (string, error) {
	if r.hasTokenAlready(userId) {
		return "", Conflict
	}
	var token = generateRandomToken(userId) // TODO maybe i could check they dont repeat.. but i mean... 2^-128 aprox 10^-39 prob
	_, err := r.db.Exec("INSERT INTO sessions (userx, token) VALUES ($1, $2)", userId, token)
	if err != nil {
		return "", SomethingWentWrong
	}
	return token, nil
}
func (r *AuthRepo) hasTokenAlready(userId int) bool {
	row := r.db.QueryRow("SELECT userx FROM sessions WHERE userx=$1", userId)
	err := row.Scan()
	if errors.Is(err, sql.ErrNoRows) {
		return false
	}
	return true
}

func generateRandomToken(userIdSeed int) string {
	// let me cook
	now := time.Now()
	number := now.Nanosecond() + rand.Intn(424242)*userIdSeed
	stringToHash := fmt.Sprintf("Today, at time %s, a user thought of the number %d", now.String(), number)

	hasher := sha256.New()
	hasher.Write([]byte(stringToHash))
	hashedBytes := hasher.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes[:16])
	// ok, never let me near the kitchen again
	return hashedString
}
