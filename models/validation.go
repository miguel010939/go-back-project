package models

import (
	"main.go/repositories"
)

func (usf *UserSignUpForm) isValid() bool {
	_, e1 := repositories.CleanUsername(usf.Username)
	_, e2 := repositories.CleanPassword(usf.Password)
	_, e3 := repositories.CleanEmail(usf.Email)
	if e1 != nil || e2 != nil || e3 != nil {
		return false
	}
	return true
}
func (ulf *UserLogInForm) isValid() bool {
	_, e1 := repositories.CleanUsername(ulf.Username)
	_, e2 := repositories.CleanPassword(ulf.Password)
	if e1 != nil || e2 != nil {
		return false
	}
	return true
}

func (pf *ProductForm) isValid() bool {
	_, e1 := repositories.CleanName(pf.Name)
	_, e2 := repositories.CleanDescription(pf.Description)
	_, e3 := repositories.CleanUrl(pf.ImageUrl)
	if e1 != nil || e2 != nil || e3 != nil {
		return false
	}
	return true
}
