package service

import (
	"Agenda/entity"
	"time"
)

//GetTimeFromUnix : convert a unix(Int64) syntax time to a Time object
func GetTimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

/*
IsTimeOverlapForUser : check whether a duration is illegal for a user
	- has overlap : return overlappedMeetingTitle, true
	- no overlap  : return "", false
*/
func IsTimeOverlapForUser(username string, startTimeUnix, endTimeUnix int64) (string, bool) {
	for _, hostMeetingStr := range entity.GetUserHostMeetings(username) {
		meeting, _ := entity.GetMeetingInfo(hostMeetingStr)
		if IsTimeOverlap(startTimeUnix, endTimeUnix, meeting.StartTime, meeting.EndTime) {
			return hostMeetingStr, true
		}
	}
	for _, parMeetingStr := range entity.GetUserHostMeetings(username) {
		meeting, _ := entity.GetMeetingInfo(parMeetingStr)
		if IsTimeOverlap(startTimeUnix, endTimeUnix, meeting.StartTime, meeting.EndTime) {
			return parMeetingStr, true
		}
	}
	return "", false
}

/*
IsTimeOverlap : check whether two given time(represented by Unix time format - int64) overlap
*/
func IsTimeOverlap(s1, e1, s2, e2 int64) bool {
	return !(e1 <= s2 || e2 <= s1)
}
