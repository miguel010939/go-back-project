package repositories

import (
	"database/sql"
	"errors"
)

type BidRepo struct {
	db *sql.DB
}

func NewBidRepo(db *sql.DB) *BidRepo {
	return &BidRepo{db}
}

func (r *BidRepo) MakeBid(userId int, productId int, amount float32) (int, error) {
	if amount <= 0 {
		return -1, InvalidInput
	}
	var testId int
	getProductQuery := "SELECT id FROM products WHERE id = $1"
	err := r.db.QueryRow(getProductQuery, productId).Scan(&testId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, NotFound
		}
		return -1, SomethingWentWrong
	}
	// i dont need to check userId because it comes from the token
	insertQuery := `INSERT INTO bids (userx, product, amount) 
					VALUES ($1, $2, $3) RETURNING id`
	var bidId int
	err2 := r.db.QueryRow(insertQuery, userId, productId, amount).Scan(&bidId)
	if err2 != nil {
		return -1, SomethingWentWrong
	}
	return bidId, nil
}
func (r *BidRepo) DeleteBid(userId int, bidId int) error {
	var testBidId int
	getBidQuery := "SELECT id FROM bids WHERE id = $1"
	err := r.db.QueryRow(getBidQuery, bidId).Scan(&testBidId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return NotFound
		}
		return SomethingWentWrong
	}
	// again, i dont need to check userId because it comes from the token
	deleteQuery := `DELETE FROM bids 
					WHERE userx=$1 AND id=$2`
	result, err1 := r.db.Exec(deleteQuery, userId, bidId)
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
