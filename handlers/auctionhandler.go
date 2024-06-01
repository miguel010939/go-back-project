package handlers

import (
	"database/sql"
	"main.go/repositories"
	"net/http"
	"strconv"
)

type BidMessage struct {
	Amount float32 `json:"amount"`
}

func NewBidMessage(amount float32) *BidMessage {
	return &BidMessage{amount}
}

type AuctionHandler struct {
	auth     repositories.AuthRepo
	prod     repositories.ProductRepo
	auctions map[int]*Auction // TODO Create and Delete Auctions
}

func NewAuctionHandler(db *sql.DB) *AuctionHandler {
	return &AuctionHandler{
		auth:     *repositories.NewAuthRepo(db),
		prod:     *repositories.NewProductRepo(db),
		auctions: make(map[int]*Auction), // maps the product id to its associated auction
	}
}

func (auh *AuctionHandler) PostAuction(w http.ResponseWriter, r *http.Request) {
	// TODO
}
func (auh *AuctionHandler) DeleteAuction(w http.ResponseWriter, r *http.Request) {
	// TODO
}
func (auh *AuctionHandler) ObserveAuctionHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	// product
	prodIdStr := r.URL.Query().Get("product")
	if prodIdStr == "" {
		http.Error(w, "missing product id", http.StatusBadRequest)
		return
	}
	prodId, err1 := strconv.Atoi(prodIdStr)
	if err1 != nil {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}
	// user
	token := r.Header.Get("sessionid")
	if token == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}
	userId, err2 := auh.auth.GetID(token)
	if err2 != nil {
		errorDispatch(w, r, err2)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "invalid product id", http.StatusBadRequest)
		return
	}

	obs := NewAuctionObserver(userId)
	auction, exists := auh.auctions[prodId]
	if !exists {
		http.Error(w, "product not in auction", http.StatusNotFound)
		return
	}

	auction.subscribe(obs)
	obs.Listen(r.Context(), w, flusher)
	auction.unsubscribe(obs)
}
