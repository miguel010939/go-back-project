package factories

import (
	"fmt"
	"main.go/models"
	"math/rand"
)

var (
	names = []string{"manuel", "jose", "luis", "adrian", "rafael", "ivan", "juan",
		"ana", "eva", "elena", "alicia", "mireia", "uxia", "celia"}
	passwords    = []string{"12", "123", "1234", "12345", "123456", "contrasenha", "password", "nomoreideas"}
	productNames = []string{"car", "airplane", "mockingbird", "window", "computer", "cookie", "stamp", "rock", "sleeve",
		"paper", "sofa", "boot", "door"}
	words = []string{"awesome", "goose", "is", "has", "a", "opening", "safe", "children", "snow", "sweet", "unhealthy",
		"cat", "runs", "in", "wall"}
	images = []string{"www.something1.com", "www.something2.com", "www.something3.com", "www.something4.com",
		"www.something5.com", "www.something6.com", "www.something7.com", "www.something8.com", "www.something9.com"}
)

func randomUser() *models.UserSignUpForm {
	username := names[rand.Intn(len(names))]
	email := fmt.Sprintf("%s%d@gmail.com", username, rand.Intn(69420)) // 2 users cant have the same email
	password := passwords[rand.Intn(len(passwords))]
	return &models.UserSignUpForm{
		Username: username,
		Email:    email,
		Password: password,
	}
}

func randomProduct() *models.ProductForm {
	name := productNames[rand.Intn(len(productNames))]
	var description string
	for i := 0; i < 9; i++ {
		description += words[rand.Intn(len(words))] + " "
	}
	image := images[rand.Intn(len(images))]
	return &models.ProductForm{
		Name:        name,
		Description: description,
		ImageUrl:    image,
	}
}
