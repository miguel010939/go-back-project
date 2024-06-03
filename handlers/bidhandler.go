package handlers

import (
	"database/sql"
	"fmt"
	"main.go/repositories"
	"net/http"
	"strconv"
)

type BidHandler struct {
	repo           repositories.BidRepo
	auth           repositories.AuthRepo
	auctionHandler *AuctionHandler
}

func NewBidHandler(db *sql.DB, auctionHandler *AuctionHandler) *BidHandler {
	return &BidHandler{
		repo:           *repositories.NewBidRepo(db),
		auth:           *repositories.NewAuthRepo(db),
		auctionHandler: auctionHandler,
	}
}

func (bh *BidHandler) PostBid(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & non-optional query params product & amount, Method Post
	productIdStr := r.URL.Query().Get("product")
	amountStr := r.URL.Query().Get("amount")
	if productIdStr == "" || amountStr == "" {
		//http.Error(w, "missing query params", http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	}
	productId, e1 := strconv.Atoi(productIdStr)
	amount, e2 := strconv.ParseFloat(amountStr, 64)
	if e1 != nil || e2 != nil {
		//http.Error(w, "invalid query params", http.StatusBadRequest)
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
	userId, err2 := bh.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}
	if !bh.auctionHandler.AuctionExists(productId) {
		//http.Error(w, "product not in auction", http.StatusNotFound)
		errorDispatch(w, r, repositories.NotFound)
		return
	}
	bidId, err3 := bh.repo.MakeBid(userId, productId, float32(amount))
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	bh.auctionHandler.BidForProduct(userId, productId, float32(amount))
	w.Header().Set("bidid", fmt.Sprintf("%d", bidId))
	w.WriteHeader(http.StatusCreated)
	// TODO log success
}
func (bh *BidHandler) DeleteBid(w http.ResponseWriter, r *http.Request) {
	// takes header "sessionid" token & path param id (product) , method Delete
	id, err := ParseIntPathParam(r.URL.Path, "bids/")
	if err != nil {
		//http.Error(w, err.Error(), http.StatusBadRequest)
		errorDispatch(w, r, repositories.InvalidInput)
		return
	} // where does this id come from? i may have to refactor the post method so it returns the id of the bid just made, so the user can undo it

	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		//http.Error(w, "Missing token", http.StatusUnauthorized)
		errorDispatch(w, r, repositories.NoPermission)
		return
	}
	userId, err2 := bh.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	err3 := bh.repo.DeleteBid(userId, id)
	if err3 != nil {
		errorDispatch(w, r, err3)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	// TODO log success
}
