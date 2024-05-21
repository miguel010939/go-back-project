package repositories

import "database/sql"

type BidRepo struct {
	db *sql.DB
}

func NewBidRepo(db *sql.DB) *BidRepo {
	return &BidRepo{db}
}

func (r *BidRepo) MakeBid(userId int, productId int, amount float32) error {

	return nil
}
func (r *BidRepo) DeleteBid(userId int, bidId int) error {

	return nil
}
