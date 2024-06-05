package factories

import (
	"database/sql"
	"main.go/repositories"
	"math/rand"
)

func UralFactories(db *sql.DB, nUsers int, nProducts int, nFavorites int, nFollowers int) {
	if nUsers <= 0 || nProducts <= 0 || nFavorites <= 0 || nFollowers <= 0 {
		panic("invalid config")
	}
	userFactory(db, nUsers)
	productFactory(db, nUsers, nProducts)
	//favoriteFactory(db, nUsers, nProducts, nFavorites)
	followerFactory(db, nUsers, nFollowers)
}

func userFactory(db *sql.DB, nUsers int) {
	userRepo := repositories.NewUserRepo(db)
	authRepo := repositories.NewAuthRepo(db)
	for i := 0; i < nUsers; i++ {
		user := randomUser()
		_, err := userRepo.UserSignUp(user, authRepo)
		// control for 409, just in case
		if err != nil {
			i--
		}
	}
}
func productFactory(db *sql.DB, nUsers int, nProducts int) {
	productRepo := repositories.NewProductRepo(db)
	for i := 0; i < nProducts; i++ {
		product := randomProduct()
		user := rand.Intn(nUsers) + 1
		productRepo.SaveProduct(user, product)
	}
}

/*
	func favoriteFactory(db *sql.DB, nUsers int, nProducts int, nFavorites int) {
		favoriteRepo := repositories.NewFavoriteRepo(db)
		for i := 0; i < nFavorites; i++ {
			err := favoriteRepo.SaveFavorite(rand.Intn(nUsers)+1, rand.Intn(nProducts)+1)
			// control for 409, specially
			if err != nil {
				i--
			}
		}
	}
*/
func followerFactory(db *sql.DB, nUsers int, nFollowers int) {
	followerRepo := repositories.NewFollowerRepo(db)
	for i := 0; i < nFollowers; i++ {
		// maybe a user could follow himself, we dont care
		err := followerRepo.FollowSomeone(rand.Intn(nUsers)+1, rand.Intn(nUsers)+1)
		// control for 409, specially
		if err != nil {
			i--
		}
	}
}
