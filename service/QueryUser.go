package service

import (
	"Agenda/entity"
)

/*
QueryUser: Query a user
	- if success, returns true, "", userInfoString
	- otherwise, returns false, errorMsg, ""
*/
func QueryUser(username string) (bool, string, string) {
	_, haslogin := entity.GetCurrentUser()
	if !haslogin {
		return false, "未登录", ""
	}
	userinfo, exist := entity.GetUserInfo(username)
	if !exist {
		return false, "无此用户", ""
	}

	tempString := "Name: " + username +
		"\nPhone: " + userinfo.Phone +
		"\nEmail: " + userinfo.Email
	return true, "", tempString
}
