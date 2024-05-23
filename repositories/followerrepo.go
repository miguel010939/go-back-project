package repositories

import (
	"database/sql"
	"errors"
	"main.go/models"
)

type FollowerRepo struct {
	db *sql.DB
}

func NewFollowerRepo(db *sql.DB) *FollowerRepo {
	return &FollowerRepo{db}
}

func (r *FollowerRepo) GetUsersWhomIFollow(userId int) ([]*models.UserRepresentation, error) {
	query := `SELECT u.id, u.username FROM followers f, user u 
            		WHERE f.usera=$1 AND f.userb=u.id`
	rows, err1 := r.db.Query(query, userId)
	if err1 != nil {
		return nil, SomethingWentWrong
	}
	defer rows.Close()

	var followers []*models.UserRepresentation
	for rows.Next() {
		var user models.UserRepresentation
		err := rows.Scan(&user.ID, &user.Username)
		if err != nil {
			return nil, SomethingWentWrong
		}
		followers = append(followers, &user)
	}
	if err := rows.Err(); err != nil {
		return nil, SomethingWentWrong
	}
	if len(followers) == 0 {
		return nil, Empty
	}
	return followers, nil
}
func (r *FollowerRepo) FollowSomeone(userId int, followedUserId int) error {
	following, err1 := r.isFollowing(userId, followedUserId)
	if err1 != nil {
		return SomethingWentWrong
	}
	if following {
		return Conflict
	}

	insertQuery := `INSERT INTO followers (usera, userb) 
					VALUES ($1, $2)`
	_, err2 := r.db.Exec(insertQuery, userId, followedUserId)
	if err2 != nil {
		return SomethingWentWrong
	}
	return nil
}
func (r *FollowerRepo) UnfollowSomeone(userId int, unfollowedUserId int) error {
	deleteQuery := `DELETE  FROM followers 
					WHERE usera=$1 AND userb=$2`
	result, err1 := r.db.Exec(deleteQuery, userId, unfollowedUserId)
	if err1 != nil {
		return SomethingWentWrong
	}
	rows, err2 := result.RowsAffected()
	if err2 != nil {
		return SomethingWentWrong
	}
	if rows == 0 {
		return NotFound
	}
	return nil
}

func (r *FollowerRepo) isFollowing(userId int, followedUserId int) (bool, error) {
	checkQuery := `SELECT id FROM followers WHERE usera=$1 AND userb=$2`
	err := r.db.QueryRow(checkQuery, userId, followedUserId).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, SomethingWentWrong
	}
	return true, nil
}
