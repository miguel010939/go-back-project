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
	var id int

	return id, nil
}
func (r *ProductRepo) DeleteProduct(id int, deletingUserId int) error {

	return nil
}
