package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"main.go/models"
)

type FavoriteRepo struct {
	db *sql.DB
}

func NewFavoriteRepo(db *sql.DB) *FavoriteRepo {
	return &FavoriteRepo{db}
}
func (r *FavoriteRepo) GetFavorites(userId int, limit int, offset int) ([]*models.ProductRepresentation, error) {
	var arg1, arg2 any
	var limitNumber, offsetNumber = Number(limit), Number(offset)
	piles := NewPairOfRelatedPiles([]any{arg1, arg2}, []optional{limitNumber, offsetNumber})
	piles.MakeAssociation()

	query := customGetFavoritesQuery(limit, offset)
	rows, err1 := r.db.Query(query, userId, piles.args[0], piles.args[1])
	if err1 != nil {
		return nil, SomethingWentWrong
	}
	defer rows.Close()

	var favorites []*models.ProductRepresentation
	for rows.Next() {
		var product models.ProductRepresentation
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageUrl, &product.UserID)
		if err != nil {
			return nil, SomethingWentWrong
		}
		favorites = append(favorites, &product)
	}
	if err := rows.Err(); err != nil {
		return nil, SomethingWentWrong
	}
	if len(favorites) == 0 {
		return nil, Empty
	}
	return favorites, nil
}
func (r *FavoriteRepo) SaveFavorite(userId int, productId int) error {
	favorite, err1 := r.isFavorite(userId, productId)
	if err1 != nil {
		return SomethingWentWrong
	}
	if favorite {
		return Conflict
	}
	insertQuery := `INSERT INTO favorites (userx, product) 
					VALUES ($1, $2)`
	_, err := r.db.Exec(insertQuery, userId, productId)
	if err != nil {
		return SomethingWentWrong
	}
	return nil
}
func (r *FavoriteRepo) DeleteFavorite(userId int, productId int) error {
	deleteQuery := `DELETE FROM favorites 
					WHERE userx=$1 AND product=$2`
	result, err1 := r.db.Exec(deleteQuery, userId, productId)
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

func (r *FavoriteRepo) isFavorite(userId int, productId int) (bool, error) {
	checkQuery := `SELECT id FROM favorites WHERE userx=$1 AND product=$2`
	err := r.db.QueryRow(checkQuery, userId, productId).Scan()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
		return false, SomethingWentWrong
	}
	return true, nil
}

func customGetFavoritesQuery(limit int, offset int) string {
	var stringLimit, stringOffset string // if a string is not initialized, its value is ""
	query := `SELECT p.id, p.namex, p.description, p.imageurl, p.userx 
				FROM favorites f, products p 
            		WHERE f.userx=$1 AND f.product=p.id
            		ORDER BY p.namex DESC%s%s`
	if limit >= 0 {
		stringLimit = " LIMIT $2"
	}
	if offset >= 0 {
		stringOffset = " OFFSET $3"
	}
	customQuery := fmt.Sprintf(query, stringLimit, stringOffset)
	return customQuery
}
