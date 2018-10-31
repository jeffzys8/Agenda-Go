package entity

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"time"
)

//MeetingInfo :
type MeetingInfo struct {
	StartTime int64 //recorded with UNIX time
	EndTime   int64 //recorded with UNIX time
	Host      string
	Partics   []string
}

var meetings = make(map[string]*MeetingInfo)
var meetingfilename = "entity/meetings.json"

// TimeFormat : output for time specification
var TimeFormat = "2006-1-2 15:04"

//LoadMeetings : load the data of meetings
func LoadMeetings() {
	meetingJSON, err := ioutil.ReadFile(meetingfilename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(meetingJSON, &meetings)
	if err != nil {
		panic(err)
	}
}

//SaveMeetings : save the data of meetings
func SaveMeetings() {

	meetingJSON, jsonerr := json.Marshal(meetings)
	if jsonerr != nil {
		panic(jsonerr)
	}
	err := ioutil.WriteFile(meetingfilename, meetingJSON, 0644)
	if err != nil {
		panic(err)
	}
}

//GetTimeFromUnix : convert a unix(Int64) syntax time to a Time object
func GetTimeFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0)
}

//GetMeetingInfo : get the info of a specific meeting
func GetMeetingInfo(title string) (*MeetingInfo, bool) {
	info, ok := meetings[title]
	return info, ok
}

//CreateMeeting : create a meeting
func CreateMeeting(title string, startTime, endTime int64, host, partic string) {
	meetings[title] = &MeetingInfo{StartTime: startTime, EndTime: endTime, Host: host, Partics: []string{partic}}
}

//IsTimeOverlapForUser : check whether a duration is illegal for a user
func IsTimeOverlapForUser(username string, startTimeUnix, endTimeUnix int64) (string, bool) {
	for _, hostMeetingStr := range GetUserHostMeetings(username) {
		meeting, _ := GetMeetingInfo(hostMeetingStr)
		if IsTimeOverlap(startTimeUnix, endTimeUnix, meeting.StartTime, meeting.EndTime) {
			return hostMeetingStr, true
		}
	}
	for _, parMeetingStr := range GetUserHostMeetings(username) {
		meeting, _ := GetMeetingInfo(parMeetingStr)
		if IsTimeOverlap(startTimeUnix, endTimeUnix, meeting.StartTime, meeting.EndTime) {
			return parMeetingStr, true
		}
	}
	return "", false
}

//IsTimeOverlap : check whether two given time overlap
func IsTimeOverlap(s1, e1, s2, e2 int64) bool {
	return !(e1 <= s2 || e2 <= s1)
}

//RemoveParticFromMeeting : as it says
func RemoveParticFromMeeting(title, username string) {

	meetingInfo, _ := GetMeetingInfo(title)
	for index, partName := range meetingInfo.Partics {
		if strings.EqualFold(username, partName) {
			tempsilce := meetingInfo.Partics[index+1:]
			meetingInfo.Partics = append([]string{}, meetingInfo.Partics[0:index]...)
			meetingInfo.Partics = append(meetingInfo.Partics, tempsilce...)
			RemoveMeetingIfEmpty(title)
			break
		}
	}

}

//RemoveMeetingIfEmpty : remove the meeting both from 'meetings' or from the host list of users if it has no participators
func RemoveMeetingIfEmpty(title string) {
	meetingInfo, _ := GetMeetingInfo(title)
	if len(meetingInfo.Partics) == 0 {
		RemoveHostMeetingFromUser(meetingInfo.Host, title)
		delete(meetings, title)
	}
}

//DeleteMeeting :
func DeleteMeeting(title string) {
	meetingInfo, _ := GetMeetingInfo(title)
	for _, pName := range meetingInfo.Partics {
		RemovePartMeetingFromUser(pName, title)
	}
	RemoveHostMeetingFromUser(meetingInfo.Host, title)
	delete(meetings, title)
}
