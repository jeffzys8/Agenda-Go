package service

import (
	"Agenda/entity"
	"strings"
)

/*
AddParticToMeeting : Add a participator to a meeting host by current user
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func AddParticToMeeting(title, particName string) (bool, string) {
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

	if _, hasParc := entity.UserHasParcMeeting(particName, title); hasParc {
		return false, "该用户已是会议成员."
	}

	if _, overlap := IsTimeOverlapForUser(particName, meetingInfo.StartTime, meetingInfo.EndTime); overlap {
		return false, "该用户时间冲突."
	}
	entity.UserAddParticMeeting(particName, title)
	entity.MeetingAddPartic(title, particName)
	entity.WriteLog("AddParticipator: host(" + hostname + ") adds participator (" + particName + ") to meeting [" + title + "]")
	return true, ""
}
