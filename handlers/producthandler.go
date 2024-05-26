package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type ProductHandler struct {
	repo repositories.ProductRepo
	auth repositories.AuthRepo
}

func NewProductHandler(db *sql.DB) *ProductHandler {
	return &ProductHandler{
		repo: *repositories.NewProductRepo(db),
		auth: *repositories.NewAuthRepo(db),
	}
}

func (ph *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	// takes path param id:int , method get
}
func (ph *ProductHandler) GetListProducts(w http.ResponseWriter, r *http.Request) {
	// takes query params user, limit, offset, method Get
}
func (ph *ProductHandler) PostNewProduct(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & JSON body productform, method Post
}
func (ph *ProductHandler) DeleteOrSellProduct(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id, Method Delete
}
