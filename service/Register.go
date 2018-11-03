package service

import (
	"Agenda/entity"
	"fmt"
	"regexp"
)

const (
	phoneReg = "^(13[0-9]|14[57]|15[0-35-9]|18[07-9])\\d{8}$"
	emailReg = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

/*
Register : For registration
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func Register(username, password, phone, email string) (bool, string) {
	if _, loginned := entity.GetCurrentUser(); loginned {
		return false, "已登录."
	}

	if _, exist := entity.GetUserInfo(username); exist {
		return false, "该用户已注册，不可重复注册"
	}

	passCheck, passErr := checkPassword(password)
	if !passCheck {
		return false, "密码格式错误: " + passErr
	}

	if !checkEmail(email) {
		return false, "邮箱格式错误"
	}

	if !checkPhone(phone) {
		return false, "邮箱格式错误"
	}

	entity.UserCreate(username, password, phone, email)
	fmt.Println("成功注册!")
	entity.WriteLog("Register: Successful registration by (" + username + ")")
	return true, ""
}

func checkPassword(passwd string) (bool, string) {

	indNum := [4]int{0, 0, 0, 0}
	spCode := []byte{'!', '@', '#', '$', '%', '^', '&', '*', '_', '-'}

	if len(passwd) < 6 {
		return false, "password too short"
	}

	passwdByte := []byte(passwd)

	for _, i := range passwdByte {
		if i >= 'A' && i <= 'Z' {
			indNum[0] = 1
			continue
		}

		if i >= 'a' && i <= 'z' {
			indNum[1] = 1
			continue
		}

		if i >= '0' && i <= '9' {
			indNum[2] = 1
			continue
		}

		notEnd := 0
		for _, s := range spCode {
			if i == s {
				indNum[3] = 1
				notEnd = 1
				break
			}
		}

		if notEnd != 1 {
			return false, "Unsupport code"
		}
	}

	codeCount := 0

	for _, i := range indNum {
		codeCount += i
	}

	if codeCount < 3 {
		return false, "Too simple password"
	}

	return true, ""
}

func checkPhone(phone string) bool {
	reg := regexp.MustCompile(phoneReg)
	return reg.MatchString(phone)
}

func checkEmail(email string) bool {
	reg := regexp.MustCompile(emailReg)
	if !reg.MatchString(email) {
		return false
	}
	return true
}
