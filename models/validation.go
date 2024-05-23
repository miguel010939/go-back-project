package models

import (
	"crypto/sha256"
	"encoding/hex"
	"main.go/repositories"
)

func (usf *UserSignUpForm) IsValiD() bool {
	_, e1 := repositories.CleanUsername(usf.Username)
	_, e2 := repositories.CleanPassword(usf.Password)
	_, e3 := repositories.CleanEmail(usf.Email)
	if e1 != nil || e2 != nil || e3 != nil {
		return false
	}
	return true
}
func (ulf *UserLogInForm) IsValiD() bool {
	_, e1 := repositories.CleanUsername(ulf.Username)
	_, e2 := repositories.CleanPassword(ulf.Password)
	if e1 != nil || e2 != nil {
		return false
	}
	return true
}

func (pf *ProductForm) IsValiD() bool {
	_, e1 := repositories.CleanName(pf.Name)
	_, e2 := repositories.CleanDescription(pf.Description)
	_, e3 := repositories.CleanUrl(pf.ImageUrl)
	if e1 != nil || e2 != nil || e3 != nil {
		return false
	}
	return true
}

func (ulf *UserLogInForm) HashPwd() {
	ulf.Password = hashPassword(ulf.Password)
}
func (usf *UserSignUpForm) HashPwd() {
	usf.Password = hashPassword(usf.Password)
}
func hashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	hashedBytes := hasher.Sum(nil)               //32 byte output
	hashedPwd := hex.EncodeToString(hashedBytes) //64 chars in hex
	return hashedPwd
}
