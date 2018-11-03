package entity

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

//MeetingInfo :
type MeetingInfo struct {
	StartTime int64 //recorded with UNIX time
	EndTime   int64 //recorded with UNIX time
	Host      string
	Partics   []string
}

/*
users : Record the data for all meetings, unaccessable to outside
	- exclusive, not accessable to other packages
	- username(string) as key
*/
var meetings = make(map[string]*MeetingInfo)

/*
meetingfilename : the file path to store meetings.json
*/
var meetingfilename = "entity/meetings.json"

/*
LoadMeetings : Load meetings from 'meetings.json'
*/
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

/*
SaveMeetings : Save meetings to 'meetings.json'
*/
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

/*
GetMeetingInfo : Get the info of a specific meeting
	- return a copy, therefore outside package won't be able to modify data
*/
func GetMeetingInfo(title string) (MeetingInfo, bool) {
	info, ok := meetings[title]
	copy := *info
	return copy, ok
}

/*
MeetingAddPartic : Add a participator to a meeting
*/
func MeetingAddPartic(title, particName string) {
	meetingInfo, mExist := GetMeetingInfo(title)
	if !mExist {
		panic(mExist)
	}
	if _, uExist := GetUserInfo(particName); !uExist {
		panic(uExist)
	}
	meetingInfo.Partics = append(meetingInfo.Partics, particName)
}

/*
MeetingRemovePartic : Remove a participator from a meeting
*/
func MeetingRemovePartic(title, particName string) {
	meetingInfo, mExist := GetMeetingInfo(title)
	if !mExist {
		panic(mExist)
	}
	for index, tempName := range meetingInfo.Partics {
		if strings.EqualFold(particName, tempName) {
			meetingInfo.Partics = append(meetingInfo.Partics[0:index], meetingInfo.Partics[index+1:]...)
			UserRemoveParticMeeting(tempName, title)
			if len(meetingInfo.Partics) == 0 {
				MeetingDelete(title)
			}
			break
		}
	}
}

/*
MeetingDelete : Delete a meeting
*/
func MeetingDelete(title string) {

	meetingInfo, exist := meetings[title]
	if !exist {
		panic(exist)
	}
	// 删除 users 中对该 Meeting的引用
	for _, pName := range meetingInfo.Partics {
		UserRemoveParticMeeting(pName, title)
	}
	UserRemoveHostMeeting(meetingInfo.Host, title)
	delete(meetings, title)
}

/*
MeetingCreate : Create a meeting
*/
func MeetingCreate(title string, startTime, endTime int64, hostName, particName string) {
	if _, exist := meetings[title]; exist {
		panic(exist)
	}
	meetings[title] = &MeetingInfo{StartTime: startTime, EndTime: endTime, Host: hostName, Partics: []string{particName}}
	UserAddHostMeeting(hostName, title)
	UserAddParticMeeting(particName, title)
}
