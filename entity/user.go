package entity

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

/*
UserInfo : struct for Users infos
*/
type UserInfo struct {
	Password     string
	Email        string
	Phone        string
	HostMeetings []string // Meetings as host
	ParMeetings  []string // Meetings as participators
}

/*
users : Record the data for all users, unaccessable to outside
	- exclusive, not accessable to other packages
	- username(string) as key
*/
var users = make(map[string]*UserInfo)

/*
userfilename : the file path to store users.json
*/
var userfilename = "entity/users.json"

/*
LoadUsers : Load the data of users from 'users.json'
*/
func LoadUsers() {
	userJSON, err := ioutil.ReadFile(userfilename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(userJSON, &users)
	if err != nil {
		panic(err)
	}
}

/*
SaveUsers : Save the data of users to 'users.json'
*/
func SaveUsers() {
	userJSON, _ := json.Marshal(users)
	err := ioutil.WriteFile(userfilename, userJSON, 0644)
	if err != nil {
		panic(err)
	}
}

/*
GetUserInfo : Return a copy of the userinfo of the given username
	- return a copy, therefore outside package won't be able to modify data
*/
func GetUserInfo(username string) (UserInfo, bool) {
	user, exist := users[username]
	copy := *user
	return copy, exist
}

/*
UserCreate : Create a new user
*/
func UserCreate(name, pass, phone, email string) {
	/*
		_pass, err := DES.TripleDesEncrypt([]byte("pass"), KEY)
		if err != nil {
			panic(err)
		}
		pass = string(_pass[:])
	*/
	if _, exist := users[name]; exist {
		panic(exist)
	}
	users[name] = &UserInfo{Password: pass, Phone: phone, Email: email}
}

/*
UserAddParticMeeting : Add a meeting to a users' participated meeting list
*/
func UserAddParticMeeting(username, title string) {
	if _, exist := users[username]; exist {
		panic(exist)
	}
	users[username].ParMeetings = append(users[username].ParMeetings, title)
}

/*
UserAddHostMeeting : Add a meeting to a uers' host meeting list
*/
func UserAddHostMeeting(username, title string) {
	if _, exist := users[username]; exist {
		panic(exist)
	}
	users[username].HostMeetings = append(users[username].HostMeetings, title)
}

/*
UserRemoveParticMeeting : Remove a meeting from a user's participated meeting list
*/
func UserRemoveParticMeeting(username, title string) {
	userInfo, exist := users[username]
	if !exist {
		panic(exist)
	}
	for i, t := range userInfo.ParMeetings {
		if strings.EqualFold(title, t) {
			userInfo.ParMeetings = append(userInfo.ParMeetings[0:i], userInfo.ParMeetings[i+1:]...)
			return
		}
	}
}

/*
UserRemoveHostMeeting : Remove a meeting from a user's host meeting list
*/
func UserRemoveHostMeeting(username, title string) {
	userInfo, exist := users[username]
	if !exist {
		panic(exist)
	}
	for i, t := range userInfo.HostMeetings {
		if strings.EqualFold(title, t) {
			userInfo.HostMeetings = append(userInfo.HostMeetings[0:i], userInfo.HostMeetings[i+1:]...)
			break
		}
	}
}

/*
UserDelete : Delete a user, together with all its meetings
*/
func UserDelete(username string) {
	userInfo, exist := GetUserInfo(username)
	if !exist {
		panic(exist)
	}
	for _, title := range userInfo.HostMeetings {
		MeetingDelete(title)
	}

	for _, title := range userInfo.ParMeetings {
		MeetingRemovePartic(title, username)
	}
	delete(users, username)
}

/*
GetUserHostMeetings : Get a list of meetings held by user
*/
func GetUserHostMeetings(username string) []string {
	return users[username].HostMeetings
}

/*
GetUserParMeetings : Get a list of meetings participated by user
*/
func GetUserParMeetings(username string) []string {
	return users[username].ParMeetings
}

/*
UserHasParcMeeting : Check if user has participated a meeting, if does, returns the index
*/
func UserHasParcMeeting(username, title string) (int, bool) {
	for index, mName := range users[username].ParMeetings {
		if strings.EqualFold(mName, title) {
			return index, true
		}
	}
	return -1, false
}
