package service

import (
	"Agenda/entity"
	"time"
)

/*
QueryMeeting : query the meetings of the current user via a given duration
	- if success, returns true, "", hostMeetingInfoList, particMeetingInfoList
	- otherwise, returns false, errorMsg, [], []
*/
func QueryMeeting(startTime, endTime int64) (bool, string, []string, []string) {
	username, loginned := entity.GetCurrentUser()
	if !loginned {
		return false, "未登录", nil, nil
	}

	userInfo, userExist := entity.GetUserInfo(username)
	if !userExist {
		panic(userExist)
	}

	var hosts, partics []string
	for _, title := range userInfo.HostMeetings {
		meeting, _ := entity.GetMeetingInfo(title)
		if meeting.StartTime >= startTime && meeting.EndTime <= endTime {
			tempString :=
				"	标题: " + title +
					"\n	开始时间" + time.Unix(meeting.StartTime, 0).String() +
					"\n	结束时间" + time.Unix(meeting.StartTime, 0).String() +
					"\n	会议参与者:"
			for _, v := range meeting.Partics {
				tempString += "			" + v
			}
			hosts = append(hosts, tempString)
		}
	}

	for _, title := range userInfo.ParMeetings {
		meeting, _ := entity.GetMeetingInfo(title)
		if meeting.StartTime >= startTime && meeting.EndTime <= endTime {
			tempString :=
				"	标题: " + title +
					"\n	开始时间" + time.Unix(meeting.StartTime, 0).String() +
					"\n	结束时间" + time.Unix(meeting.StartTime, 0).String() +
					"\n	会议参与者:"
			for _, v := range meeting.Partics {
				tempString += "			" + v
			}
			partics = append(partics, tempString)
		}
	}
	return true, "", hosts, partics
}
