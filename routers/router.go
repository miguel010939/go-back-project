package routers

import (
	"database/sql"
	"main.go/handlers"
	"net/http"
)

// Routes makes the connection between the different routes and their corresponding handler functions
func Routes(r *http.ServeMux, db *sql.DB) {
	uh := handlers.NewUserHandler(db)
	r.HandleFunc("/users/new", uh.SignUpHandler)
	r.HandleFunc("/users/login", uh.LogInHandler)
	r.HandleFunc("/users/logout", uh.LogOutHandler)

	ph := handlers.NewProductHandler(db)
	r.HandleFunc("products/", ph.GetProduct)
	r.HandleFunc("products", ph.GetListProducts)
	r.HandleFunc("products/new", ph.PostNewProduct)
	r.HandleFunc("products/", ph.DeleteOrSellProduct)

	bh := handlers.NewBidHandler(db)
	r.HandleFunc("bids", bh.PostBid)
	r.HandleFunc("bids/", bh.DeleteBid)

	foh := handlers.NewFollowerHandler(db)
	r.HandleFunc("followers/follow", foh.GetUsersImFollowing)
	r.HandleFunc("followers/follow/", foh.FollowUser)
	r.HandleFunc("followers/follow/", foh.UnfollowUser)

	fah := handlers.NewFavoriteHandler(db)
	r.HandleFunc("/favorites", fah.GetFavorites)
	r.HandleFunc("/favorites/", fah.SaveFavorite)
	r.HandleFunc("/favorites/", fah.DeleteFavorite)

	r.Handle("/", http.FileServer(http.Dir("./static/")))
}
