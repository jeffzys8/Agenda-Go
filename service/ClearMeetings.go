package service

import "Agenda/entity"

/*
ClearMeetings : Clear all the meetings host by current user
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func ClearMeetings() (bool, string) {

	hostname, loginned := entity.GetCurrentUser()
	if !loginned {
		return false, "未登录"
	}

	userInfo, exist := entity.GetUserInfo(hostname)
	if !exist {
		panic(exist)
	}

	// 调用entity执行删除
	for _, title := range userInfo.HostMeetings {
		entity.MeetingDelete(title)
	}

	entity.WriteLog("ClearMeetings: user(" + hostname + ")")
	return true, ""
}
