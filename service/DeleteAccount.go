package service

import (
	"Agenda/entity"
)

/*
DeleteAccount : Delete an account
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func DeleteAccount() (bool, string) {
	username, loginned := entity.GetCurrentUser()
	if !loginned {
		return false, "未登录"
	}
	entity.UserDelete(username)
	entity.WriteLog("DeleteAccount: user(" + username + ")")
	entity.SetCurrentUser("")
	return true, ""
}
