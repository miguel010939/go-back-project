package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type BidHandler struct {
	repo repositories.BidRepo
	auth repositories.AuthRepo
}

func NewBidHandler(db *sql.DB) *BidHandler {
	return &BidHandler{
		repo: *repositories.NewBidRepo(db),
		auth: *repositories.NewAuthRepo(db),
	}
}

func (bh *BidHandler) PostBid(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & non-optional query params product & amount, Method Post
}
func (bh *BidHandler) DeleteBid(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (product) , method Delete
}
