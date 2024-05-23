package repositories

import (
	"database/sql"
	"main.go/models"
)

type ProductRepo struct {
	db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{db}
}

func (r *ProductRepo) GetProductById(id int) (*models.ProductRepresentation, error) {

	return nil, nil
}
func (r *ProductRepo) GetProducts(sellingUserId int, limit int, offset int) ([]*models.ProductRepresentation, error) {

	return nil, nil
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
