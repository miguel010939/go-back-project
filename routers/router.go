package routers

import (
	"database/sql"
	"main.go/handlers"
	"net/http"
)

// Routes makes the connection between the different routes and their corresponding handler functions
func Routes(r *http.ServeMux, db *sql.DB) {
	uh := handlers.NewUserHandler(db)
	r.HandleFunc("POST /users/new", uh.SignUpHandler)
	r.HandleFunc("POST /users/login", uh.LogInHandler)
	r.HandleFunc("DELETE /users/logout", uh.LogOutHandler)

	ph := handlers.NewProductHandler(db)
	r.HandleFunc("GET /products/", ph.GetProduct)
	r.HandleFunc("GET /products", ph.GetListProducts)
	r.HandleFunc("POST /products/new", ph.PostNewProduct)
	r.HandleFunc("DELETE /products/", ph.DeleteOrSellProduct)

	auh := handlers.NewAuctionHandler(db)
	r.HandleFunc("POST /auctions/watch", auh.ObserveAuction)
	r.HandleFunc("POST /auctions", auh.PostAuction)
	r.HandleFunc("DELETE /auctions", auh.DeleteAuction)

	bh := handlers.NewBidHandler(db, auh)
	r.HandleFunc("POST /bids", bh.PostBid)
	r.HandleFunc("DELETE /bids/", bh.DeleteBid)

	foh := handlers.NewFollowerHandler(db)
	r.HandleFunc("GET /followers/follow", foh.GetUsersImFollowing)
	r.HandleFunc("POST /followers/follow/", foh.FollowUser)
	r.HandleFunc("DELETE /followers/follow/", foh.UnfollowUser)

	fah := handlers.NewFavoriteHandler(db)
	r.HandleFunc("GET /favorites", fah.GetFavorites)
	r.HandleFunc("POST /favorites/", fah.SaveFavorite)
	r.HandleFunc("DELETE /favorites/", fah.DeleteFavorite)

	// TODO maybe its better to sub my toy path param parser with the official recent addition to the SL
	r.Handle("/{$}", http.FileServer(http.Dir("./static/")))
}
