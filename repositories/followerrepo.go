package repositories

import (
	"database/sql"
	"main.go/models"
)

type FollowerRepo struct {
	db *sql.DB
}

func NewFollowerRepo(db *sql.DB) *FollowerRepo {
	return &FollowerRepo{db}
}

func (r *FollowerRepo) GetUsersWhomIFollow(userId int) ([]*models.UserRepresentation, error) {

	return nil, nil
}
func (r *FollowerRepo) FollowSomeone(userId int, followedUserId int) error {

	return nil
}
func (r *FollowerRepo) UnfollowSomeone(userId int, unfollowedUserId int) error {

	return nil
}
