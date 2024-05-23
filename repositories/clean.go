package repositories

import (
	"errors"
	"regexp"
)

type RedFlag string

// these are up for debate tbf
const (
	usernameRedFlag    RedFlag = "[^a-z0-9]]"
	passwordRedFlag    RedFlag = "[^A-Za-z0-9%!]"
	emailRedFlag       RedFlag = "[^a-z0-9.-@]"
	nameRedFlag        RedFlag = "[^A-Za-z ]"
	descriptionRedFlag RedFlag = "[^A-Za-z,. ]"
	urlRedFlag         RedFlag = "[^a-z0-9.:/]"
	tokenRedFlag       RedFlag = "[^a-f0-9]"
)

func CleanUsername(username string) (string, error) {
	str, e := clean(usernameRedFlag, username)
	if e != nil {
		return "", e
	}
	return str, nil
}
func CleanPassword(password string) (string, error) {
	str, e := clean(passwordRedFlag, password)
	if e != nil {
		return "", e
	}
	return str, nil
}
func CleanEmail(email string) (string, error) {
	str, e := clean(emailRedFlag, email)
	if e != nil {
		return "", e
	}
	return str, nil
}
func CleanName(name string) (string, error) {
	str, e := clean(nameRedFlag, name)
	if e != nil {
		return "", e
	}
	return str, nil
}
func CleanDescription(description string) (string, error) {
	str, e := clean(descriptionRedFlag, description)
	if e != nil {
		return "", e
	}
	return str, nil
}
func CleanUrl(url string) (string, error) {
	str, e := clean(urlRedFlag, url)
	if e != nil {
		return "", e
	}
	return str, nil
}
func CleanToken(token string) (string, error) {
	str, e := clean(tokenRedFlag, token)
	if e != nil {
		return "", e
	}
	return str, nil
}

func clean(redFlag RedFlag, dirtyString string) (string, error) {
	subStrRegEx := regexp.MustCompile(string(redFlag))
	if subStrRegEx.MatchString(dirtyString) {
		return "", errors.New("invalid input, untrustworthy string")
	}
	return dirtyString, nil
}
