package models

import (
	"crypto/sha256"
	"encoding/hex"
)

func (usf *UserSignUpForm) IsValiD() bool {
	_, e1 := CleanUsername(usf.Username)
	//_, e2 := repositories.CleanPassword(usf.Password)
	_, e3 := CleanEmail(usf.Email)
	if e1 != nil || e3 != nil {
		return false
	}
	return true
}
func (ulf *UserLogInForm) IsValiD() bool {
	_, e1 := CleanUsername(ulf.Username)
	//_, e2 := repositories.CleanPassword(ulf.Password)
	if e1 != nil {
		return false
	}
	return true
}

func (pf *ProductForm) IsValiD() bool {
	_, e1 := CleanName(pf.Name)
	_, e2 := CleanDescription(pf.Description)
	//_, e3 := CleanUrl(pf.ImageUrl) Im going to skip validating the url for now... fewer headaches during dev phase
	if e1 != nil || e2 != nil /*|| e3 != nil*/ {
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
