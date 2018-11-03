package service

import (
	"Agenda/entity"
	"strings"
)

/*
CancelMeeting : Cancel a meeting host by current user
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func CancelMeeting(title string) (bool, string) {
	hostname, loginned := entity.GetCurrentUser()
	if !loginned {
		return false, "未登录"
	}
	meetingInfo, meetingExist := entity.GetMeetingInfo(title)
	if !meetingExist {
		return false, "该会议不存在."
	}
	if !strings.EqualFold(meetingInfo.Host, hostname) {
		return false, "您无该会议的操作权."
	}

	// 进行数据操作函数调用
	entity.MeetingDelete(title)
	entity.WriteLog("CancelMeeting: Meeting [" + title + "] cancelled by (" + hostname + ")")
	return true, ""
}
