package service

import (
	"Agenda/entity"
	"time"
)

/*
CreateMeating : Create a meeting
	- if success, returns true, ""
	- otherwise, returns false, errorMsg
*/
func CreateMeating(title string, startTimeUnix, endTimeUnix int64, participatorStr string) (bool, string) {
	hostname, loginned := entity.GetCurrentUser()
	if !loginned {
		return false, "未登录."
	}
	if _, exist := entity.GetMeetingInfo(title); exist {
		return false, "该会议已存在."
	}

	if endTimeUnix <= startTimeUnix || startTimeUnix < time.Now().Unix() {
		return false, "不合法的时间."
	}

	// 检查是否和host时间重合
	if meetingName, overlap := IsTimeOverlapForUser(hostname, startTimeUnix, endTimeUnix); overlap {
		return false, "该时间与您的会议[" + meetingName + "]时间冲突"
	}

	// 检查part是否存在
	_, parExist := entity.GetUserInfo(participatorStr)
	if !parExist {
		return false, "输入的用户不存在"
	}
	// 检查是否和part时间重合
	if meetingName, overlap := IsTimeOverlapForUser(participatorStr, startTimeUnix, endTimeUnix); overlap {
		return false, "该时间与参与者(" + participatorStr + ")的会议[" + meetingName + "]时间冲突"
	}

	// 调用entity接口创建会议
	entity.MeetingCreate(title, startTimeUnix, endTimeUnix, hostname, participatorStr)
	entity.WriteLog("CreateMeeting: Meeting created [" + title + "] by (" + hostname + ")" + "with initial participator (" + participatorStr + ")")
	return true, ""
}
