package service

import (
	"Agenda/entity"
	"strings"
)

/*
ExitMeeting : Exit a participated meeting
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func ExitMeeting(title string) (bool, string) {
	username, loginned := entity.GetCurrentUser()
	if !loginned {
		return false, "未登录"
	}

	meetingInfo, meetingExist := entity.GetMeetingInfo(title)
	if !meetingExist {
		return false, "该会议不存在."
	}

	if strings.EqualFold(meetingInfo.Host, username) {
		return false, "你是会议发起人，应使用取消会议"
	}

	_, isPart := entity.UserHasParcMeeting(username, title)
	if !isPart {
		return false, "你不在会议中"
	}
	entity.MeetingRemovePartic(title, username)
	entity.WriteLog("ExitMeeting: (" + username + ") exit meeting [" + title + "]")
	return true, ""
}
