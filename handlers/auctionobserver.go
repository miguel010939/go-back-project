package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const (
	eventMessageFormat = "event:%s\ndata:%s\n\n"
)

type AuctionObserver struct {
	UserId  int
	index   int
	channel chan BidMessage
}

func NewAuctionObserver(userId int) *AuctionObserver {
	return &AuctionObserver{
		UserId:  userId,
		channel: make(chan BidMessage),
	}
}

func (obs *AuctionObserver) notify(bm *BidMessage, wg *sync.WaitGroup) {
	if obs == nil {
		return
	}
	wg.Add(1)

	obs.channel <- *bm // bm is not modified in the goroutine, as it is, so no risk of race conditions

	wg.Done()
}
func (obs *AuctionObserver) Listen(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case bm := <-obs.channel:
			message, err := json.Marshal(bm)
			if err != nil {
				log.Println(err)
			}
			_, err = fmt.Fprintf(w, eventMessageFormat, "bid", string(message))
			if err != nil {
				log.Println(err)
			}
			flusher.Flush()

		case <-ctx.Done():
			return
		}
	}
}

func firstNotNilIndex(slice []*AuctionObserver) (int, error) {
	for i, whatever := range slice {
		if whatever == nil {
			return i, nil
		}
	}
	return -1, fmt.Errorf("auction max capacity exceeded")
}
