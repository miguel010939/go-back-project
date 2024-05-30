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
	var arg1, arg2, arg3 any
	var userIdNumber, limitNumber, offsetNumber = Number(sellingUserId), Number(limit), Number(offset)
	piles := NewPairOfRelatedPiles([]any{arg1, arg2, arg3}, []optional{userIdNumber, limitNumber, offsetNumber})
	piles.MakeAssociation()

	selectQuery := customGetProductsQuery(sellingUserId, limit, offset)
	// Query ignores extra args
	rows, err := r.db.Query(selectQuery, piles.args[0], piles.args[1], piles.args[2])
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
					VALUES ($1, $2, $3, $4) RETURNING id`
	var insertId int
	err := r.db.QueryRow(insertQuery, product.Name, product.Description, product.ImageUrl, userId).Scan(&insertId)
	if err != nil {
		return -1, SomethingWentWrong
	}
	return insertId, nil
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

func customGetProductsQuery(sellingUserId int, limit int, offset int) string {
	var stringUserId, stringLimit, stringOffset string // if a string is not initialized, its value is ""
	query := "SELECT id, namex, description, imageurl, userx FROM products %sORDER BY namex%s%s"
	if sellingUserId >= 0 {
		stringUserId = "WHERE userx = $1 "
	}
	if limit >= 0 {
		stringLimit = " LIMIT $2"
	}
	if offset >= 0 {
		stringOffset = " OFFSET $3"
	}
	customQuery := fmt.Sprintf(query, stringUserId, stringLimit, stringOffset)
	// inserts the strings into the string query if the related value "makes sense"
	return customQuery
}

// the idea is to assign to the arguments the values in the slice of values (if they make sense, contextually) in order
type RelatedPiles struct {
	args   []any
	values []optional
	// All of this so i dont have to code the conditionals 1 by 1
	// The cleaner way would be to pair this with the string parts of the query in a slightly more complex struct... This should work for now
}

func NewPairOfRelatedPiles(args []any, values []optional) *RelatedPiles {
	return &RelatedPiles{
		args:   args,
		values: values,
	}
}

type optional interface {
	IsItThere() bool
}

func (piles *RelatedPiles) MakeAssociation() {
	var i int // = 0
	for _, v := range piles.values {
		if v.IsItThere() {
			piles.args[i] = v
			i++
		}
	}
}

// Go doesnt let me implement an interface for int, maybe because its a built-in type
type Number int

func (n Number) IsItThere() bool {
	if n >= 0 {
		return true
	}
	return false
}
