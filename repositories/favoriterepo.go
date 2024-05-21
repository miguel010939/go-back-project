package repositories

import (
	"database/sql"
	"main.go/models"
)

type FavoriteRepo struct {
	db *sql.DB
}

func NewFavoriteRepo(db *sql.DB) *FavoriteRepo {
	return &FavoriteRepo{db}
}
func (r *FavoriteRepo) GetFavorites(userId int, limit int, offset int) ([]*models.ProductRepresentation, error) {

	return nil, nil
}
func (r *FavoriteRepo) SaveFavorite(userId int, productId int) error {

	return nil
}
func (r *FavoriteRepo) DeleteFavorite(userId int, productId int) error {

	return nil
}
