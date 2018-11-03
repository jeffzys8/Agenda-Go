package service

import (
	"Agenda/entity"
	"strings"
)

/*
Login : Service for user login
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func Login(username, password string) (bool, string) {
	// pass, error := DES.TripleDesDecrypt([]byte(_password), []byte("sfe023f_sefiel#fi32lf3e!"))
	// if error != nil {
	// 	panic(error)
	// }
	// desPassword := string(pass[:])
	// read the current file
	// password, err := DES.TripleDesDecrypt([]byte(desPassword), []byte("sfe023f_sefiel#fi32lf3e!"))
	// if err != nil {
	// 	panic(err)
	// }
	_, exist := entity.GetCurrentUser()
	if exist {
		return false, "已经登陆，无需重复登陆"
	}
	user, exist := entity.GetUserInfo(username)
	if !exist {
		return false, "账户不存在，请核对"
	}
	if strings.EqualFold(user.Password, string(password)) == false {
		return false, "密码错误，请核对"
	}
	entity.SetCurrentUser(username)
	entity.WriteLog("Login: Successful login by (" + username + ")")
	return true, ""
}
