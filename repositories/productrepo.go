package repositories

import (
	"database/sql"
	"errors"
	"fmt"
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
	selectQuery := `SELECT id, namex, description, imageurl, userx 
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

func (r *ProductRepo) GetProducts(sellingUserId int, limit int, offset int) ([]*models.ProductRepresentation, error) {
	templateQuery := "SELECT id, namex, description, imageurl, userx FROM products %s ORDER BY namex DESC %s %s"

	valueAndQueryArray := NewArrayOfValuesAndQueries(valueWithQuery{Number(sellingUserId), " WHERE userx = %v "},
		valueWithQuery{Number(limit), " LIMIT %v "}, valueWithQuery{Number(offset), " OFFSET %v "})
	valuesQueries := ArrayOfValuesAndQueries{vq: *valueAndQueryArray}
	valuesQueries.filterNonSense()
	// I was going to be a good boy and use placeholders for the SQL queries, but im in a hurry and their restrictions
	// were becoming annoying... SQL jedis can't reach the power the dark side provides
	// (I don't fear SQL injection in this context/layer, but I know this code is much more vulnerable)
	selectQuery := fmt.Sprintf(templateQuery, valuesQueries.vq[0].query, valuesQueries.vq[1].query, valuesQueries.vq[2].query)
	// This is done in 2 steps, because they are nested
	query, e := insertValues(selectQuery, valuesQueries.vq[0].value, valuesQueries.vq[1].value, valuesQueries.vq[2].value)

	rows, err := r.db.Query(query)
	if err != nil || e != nil {
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
					VALUES ($1, $2, $3, $4) RETURNING id`
	var insertId int
	err := r.db.QueryRow(insertQuery, product.Name, product.Description, product.ImageUrl, userId).Scan(&insertId)
	if err != nil {
		return -1, SomethingWentWrong
	}
	return insertId, nil
}
func (r *ProductRepo) DeleteProduct(id int, deletingUserId int) error {
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
