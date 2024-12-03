package handlers

import (
	"fmt"
	"regexp"
)

const (
	ExpEmail  = `^[a-zA-Z0-9._+-=]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	ExpPasswd = `^[A-Za-z0-9]{8,20}$`
	ExpName   = `^[a-zA-Z0-9]{3,20}$`
)

func VerifyData(info *SignUpData) bool {
	isValidEmail, _ := regexp.MatchString(ExpEmail, info.Email)
	isValidName, _ := regexp.MatchString(ExpName, info.Username)
	isValidPsswd, _ := regexp.MatchString(ExpPasswd, info.Passwd)
	fmt.Printf("email=%v %v,pqsswd=%v %v, user=%v %v\n", info.Email, isValidEmail, info.Passwd, isValidPsswd, info.Username, isValidName)
	if !isValidEmail || !isValidName || !isValidPsswd {
		return false
	}
	return true
	// email :=
}
