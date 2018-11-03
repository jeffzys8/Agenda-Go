package service

import (
	"Agenda/entity"
)

/*
Logout : as it says
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func Logout() (bool, string) {
	if username, exist := entity.GetCurrentUser(); exist {
		entity.SetCurrentUser("")
		entity.WriteLog("Logout: (" + username + ") logout")
		return true, ""
	}
	return false, "未登陆，无法登出"
}
