package entity

import (
	"regexp"
)

const (
	regular = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
)

func validate(mobileNum string) bool {
	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}

func check_phone (num string) bool{
	if validate(num) {
		//println("是手机号")
		return true
	}
	return false
}