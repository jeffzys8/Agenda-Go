package service

import (
	"Agenda/entity"
	"strings"
)

/*
RemoveParticipator : Remove a participator from a meeting
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func RemoveParticipator(title, particName string) (bool, string) {
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

	_, parcExist := entity.GetUserInfo(particName)
	if !parcExist {
		return false, "该用户不存在."
	}

	_, hasPar := entity.UserHasParcMeeting(particName, title)
	if !hasPar {
		return false, "该用户不是会议参与者."
	}
	entity.MeetingRemovePartic(title, particName)
	entity.WriteLog("RemoveParticipator: host(" + hostname + ") removes participator(" + particName + ") from meeting[" + title + "]")

	return true, ""
}
