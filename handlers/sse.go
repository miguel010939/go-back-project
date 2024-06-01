package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main.go/repositories"
	"net/http"
	"strconv"
	"sync"
)

const (
	AUCTION_CAPACITY   = 200
	eventMessageFormat = "event:%s\ndata:%s\n\n"
)

type AuctionHandler struct {
	auth     repositories.AuthRepo
	auctions map[int]*Auction // TODO Create and Delete Auctions
}

func NewAuctionHandler(db *sql.DB) *AuctionHandler {
	return &AuctionHandler{
		auth:     *repositories.NewAuthRepo(db),
		auctions: make(map[int]*Auction), // maps the product id to its associated auction
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

type BidMessage struct {
	Amount float32 `json:"amount"`
}

func NewBidMessage(amount float32) *BidMessage {
	return &BidMessage{amount}
}

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
func (a *Auction) NotifyAllObservers(bm *BidMessage) {
	for _, obs := range a.NotList {
		go obs.notify(bm, &a.NotificationWG)
	}
	a.NotificationWG.Wait()
}
func (obs *AuctionObserver) notify(bm *BidMessage, wg *sync.WaitGroup) {
	wg.Add(1)
	if obs == nil {
		return
	}
	obs.channel <- *bm // bm is not modified in the goroutine, as it is, so no risk of race conditions

	wg.Done()
}

type cliente struct {
	ID          string
	sendMessage chan EventMessage
}
type EventMessage struct {
	EventName string
	Data      any
}

func NewCliente(id string) *cliente {
	return &cliente{
		ID: id,
	}
}

type HandlerEvent struct {
	m        sync.Mutex
	clientes map[string]*cliente
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clientes: make(map[string]*cliente),
	}
}
func (h *HandlerEvent) Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	id := r.URL.Query().Get("id")

	flusher, ok := w.(http.Flusher)
	if !ok {
		//Error 400
		return
	}
	c := NewCliente(id)

	h.register(c)

	c.Online(r.Context(), w, flusher)

	h.remove(c)
}

func (h *HandlerEvent) register(cliente *cliente) {
	h.m.Lock()
	defer h.m.Unlock()
	h.clientes[cliente.ID] = cliente
}

func (h *HandlerEvent) remove(cliente *cliente) {
	h.m.Lock()
	delete(h.clientes, cliente.ID)
	defer h.m.Unlock()

}
func (h *HandlerEvent) Broadcast(m EventMessage) {
	h.m.Lock()
	defer h.m.Unlock()
	for _, cliente := range h.clientes {
		cliente.sendMessage <- m
	}
}

func (c *cliente) Online(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case m := <-c.sendMessage:
			data, err := json.Marshal(m.Data)
			if err != nil {
				log.Println(err)
			}
			const format = "event:%s\ndata:%s\n\n"
			fmt.Fprintf(w, format, m.EventName, string(data))
			flusher.Flush()

		case <-ctx.Done():
			return
		}
	}
}
