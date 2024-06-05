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

	templateQuery := `SELECT p.id, p.namex, p.description, p.imageurl, p.userx
					FROM favorites f, products p
            		WHERE f.userx=$1 AND f.product=p.id
            		ORDER BY p.namex DESC %s %s `

	valueAndQueryArray := NewArrayOfValuesAndQueries(valueWithQuery{Number(limit), " LIMIT %v "},
		valueWithQuery{Number(offset), " OFFSET %v "})
	valuesQueries := ArrayOfValuesAndQueries{vq: *valueAndQueryArray}
	valuesQueries.filterNonSense()
	// I was going to be a good boy and use placeholders for the SQL queries, but im in a hurry and their restrictions
	// were becoming annoying... SQL jedis can't reach the power the dark side provides
	// (I don't fear SQL injection in this context/layer, but I know this code is much more vulnerable)
	selectQuery := fmt.Sprintf(templateQuery, valuesQueries.vq[0].query, valuesQueries.vq[1].query)
	// This is done in 2 steps, because they are nested
	query, e := insertValues(selectQuery, valuesQueries.vq[0].value, valuesQueries.vq[1].value)

	rows, err1 := r.db.Query(query, userId)
	if err1 != nil || e != nil {
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
		return err1
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
	var testProductId int
	checkProductQuery := `SELECT id FROM products WHERE id=$1`
	err1 := r.db.QueryRow(checkProductQuery, productId).Scan(&testProductId)
	if err1 != nil {
		if errors.Is(err1, sql.ErrNoRows) {
			return false, NotFound
		}
		return false, SomethingWentWrong
	}

	var testId int
	checkFavoriteQuery := `SELECT id FROM favorites WHERE userx=$1 AND product=$2`
	err2 := r.db.QueryRow(checkFavoriteQuery, userId, productId).Scan(&testId)
	if err2 != nil {
		if errors.Is(err2, sql.ErrNoRows) {
			return false, nil
		}
		return false, SomethingWentWrong
	}
	return true, nil
}
