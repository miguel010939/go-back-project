package handlers

import (
	"sync"
)

const (
	AUCTION_CAPACITY = 200
)

type Auction struct {
	ProdId         int                // Product id
	Max            float32            // Current price to beat
	Subs           map[int]int        // Maps the ids of the observers/subscribers to the auction with their index in the notification list
	NotList        []*AuctionObserver // Notification list for all current observers
	NotificationWG sync.WaitGroup
}

func NewAuction(prodId int) *Auction {
	return &Auction{
		ProdId:  prodId,
		Max:     0.0,
		Subs:    make(map[int]int),
		NotList: make([]*AuctionObserver, AUCTION_CAPACITY),
	}
}

func (a *Auction) Bid(amount float32) {
	if amount <= a.Max {
		return
	}
	a.Max = amount
	bm := NewBidMessage(amount)
	a.NotifyAllObservers(bm)
}

func (a *Auction) subscribe(observer *AuctionObserver) {
	index, err := firstNotNilIndex(a.NotList)
	if err != nil {
		return
	}
	a.Subs[observer.UserId] = index
	a.NotList[index] = observer
	observer.index = index
}
func (a *Auction) unsubscribe(observer *AuctionObserver) {
	index := a.Subs[observer.UserId]
	delete(a.Subs, observer.UserId)
	a.NotList[index] = nil
}
func (a *Auction) NotifyAllObservers(bm *BidMessage) {
	for _, obs := range a.NotList {
		go obs.notify(bm, &a.NotificationWG)
	}
	a.NotificationWG.Wait()
}
