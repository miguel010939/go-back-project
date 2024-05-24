package repositories

import (
	"database/sql"
	"errors"
	"main.go/models"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) GetProductById(id int) (*models.ProductRepresentation, error) {
	var product models.ProductRepresentation
	selectQuery := `SELECT id, name, description, imageurl, userx 
					FROM products WHERE id = $1`
	row := r.db.QueryRow(selectQuery, id)
	err := row.Scan(&product.ID, &product.Name, &product.Description, &product.ImageUrl, &product.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, NotFound
		}
		return nil, SomethingWentWrong
	}
	return &product, nil
}

// TODO MAKE THE ARGUMENTS OPTIONAL, in this repo method and others similar
func (r *ProductRepo) GetProducts(sellingUserId int, limit int, offset int) ([]*models.ProductRepresentation, error) {
	selectQuery := `SELECT id, name, description, imageurl, userx 
					FROM products WHERE userx = $1
					ORDER BY name LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(selectQuery, sellingUserId, limit, offset)
	if err != nil {
		return nil, SomethingWentWrong
	}
	defer rows.Close()

	var products []*models.ProductRepresentation
	for rows.Next() {
		var product models.ProductRepresentation
		err2 := rows.Scan(&product.ID, &product.Name, &product.Description, &product.ImageUrl, &product.UserID)
		if err2 != nil {
			return nil, SomethingWentWrong
		}
		products = append(products, &product)
	}
	if err3 := rows.Err(); err3 != nil {
		return nil, SomethingWentWrong
	}
	if len(products) == 0 {
		return nil, Empty
	}
	return products, nil
}
func (r *ProductRepo) SaveProduct(userId int, product *models.ProductForm) (int, error) {
	if !product.IsValiD() {
		return -1, InvalidInput
	}
	insertQuery := `INSERT INTO products (namex, description, imageurl, userx) 
					VALUES ($1, $2, $3, $4)`
	result, err := r.db.Exec(insertQuery)
	if err != nil {
		return -1, SomethingWentWrong
	}
	insertId, err2 := result.LastInsertId()
	if err2 != nil {
		return -1, SomethingWentWrong
	}
	return int(insertId), nil
}
func (r *ProductRepo) DeleteProduct(id int, deletingUserId int) error { //TODO remember to get this user id from the token
	deleteQuery := `DELETE FROM products WHERE id = $1 AND userx = $2`
	result, err1 := r.db.Exec(deleteQuery, id, deletingUserId)
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
