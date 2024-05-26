package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type ProductHandler struct {
	repo repositories.ProductRepo
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		repo: *repositories.NewProductRepo(db),
	}
}

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {

}
func (ph *ProductHandler) GetListProducts(w http.ResponseWriter, r *http.Request) {

}
func (ph *ProductHandler) PostNewProduct(w http.ResponseWriter, r *http.Request) {

}
func (ph *ProductHandler) DeleteOrSellProduct(w http.ResponseWriter, r *http.Request) {

}
