package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
)

type BidHandler struct {
	repo repositories.BidRepo
}

func NewBidHandler(db *sql.DB) *BidHandler {
	return &BidHandler{
		repo: *repositories.NewBidRepo(db),
	}
}

func (bh *BidHandler) PostBid(w http.ResponseWriter, r *http.Request) {

}
func (bh *BidHandler) DeleteBid(w http.ResponseWriter, r *http.Request) {

}
