package handlers

import (
	"database/sql"
	"encoding/json"
	"main.go/models"
	"main.go/repositories"
	"net/http"
	"strconv"
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
	id, err := ParseIntPathParam(r.URL.Path, "products/")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	}
	product, err := ph.repo.GetProductById(id)
	if err != nil {
		errorDispatch(w, r, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(*product)
	if err2 != nil {
		//http.Error(w, err2.Error(), http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	}
	// TODO log success
}
func (ph *ProductHandler) GetListProducts(w http.ResponseWriter, r *http.Request) {
	// takes query params user, limit, offset, method Get
	userStr := r.URL.Query().Get("user")
	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")
	// if the query param is not found, the var is assigned a nonsensical value, so it is ignored by the repo method
	if userStr == "" {
		userStr = "-1"
	}
	if limitStr == "" {
		limitStr = "-1"
	}
	if offsetStr == "" {
		offsetStr = "-1"
	}
	user, err1 := strconv.Atoi(userStr)
	limit, err2 := strconv.Atoi(limitStr)
	offset, err3 := strconv.Atoi(offsetStr)
	if err1 != nil || err2 != nil || err3 != nil {
		//http.Error(w, "invalid query params", http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	}

	products, err := ph.repo.GetProducts(user, limit, offset)
	if err != nil {
		return
	} // TODO maybe later simplify this so that the repo method itself returns an array of values, not pointers
	var productArray []models.ProductRepresentation
	for _, productPtr := range products {
		productArray = append(productArray, *productPtr)
	}

	jsonData, e := json.Marshal(productArray)
	if e != nil {
		//http.Error(w, e.Error(), http.StatusInternalServerError)
		errorDispatch(w, r, repositories.SomethingWentWrong)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, e2 := w.Write(jsonData)
	if e2 != nil {
		//http.Error(w, e2.Error(), http.StatusInternalServerError)
		errorDispatch(w, r, repositories.SomethingWentWrong)
		return
	}
	// TODO log success
}
func (ph *ProductHandler) PostNewProduct(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & JSON body productform, method Post
	// product
	var product models.ProductForm
	if err1 := json.NewDecoder(r.Body).Decode(&product); err1 != nil {
		//http.Error(w, err1.Error(), http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	}
	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		//http.Error(w, "Missing token", http.StatusUnauthorized)
		errorDispatch(w, r, repositories.NoPermission)
		return
	}
	id, err2 := ph.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}
	savedProductId, err3 := ph.repo.SaveProduct(id, &product)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.Header().Set("id", strconv.Itoa(savedProductId))
	w.WriteHeader(http.StatusCreated)
	// TODO log success
}
func (ph *ProductHandler) DeleteOrSellProduct(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id, Method Delete
	id, err1 := ParseIntPathParam(r.URL.Path, "products/")
	if err1 != nil {
		//http.Error(w, err1.Error(), http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	}
	token := r.Header.Get("sessionid")
	if token == "" {
		//http.Error(w, "Missing token", http.StatusUnauthorized)
		errorDispatch(w, r, repositories.NoPermission)
		return
	}
	userId, err2 := ph.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	err3 := ph.repo.DeleteProduct(id, userId)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.WriteHeader(http.StatusNoContent) // TODO update docs to include all the error codes that were added later
	// TODO log success
}
