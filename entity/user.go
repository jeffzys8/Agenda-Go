package entity

import (
	"encoding/json"
	"io/ioutil"
)

// UserInfo : struct for Users infos
type UserInfo struct {
	Password     string
	Email        string
	Phone        string
	HostMeetings []string // Meetings as host
	ParMeetings  []string // Meetings as participators
}

// users : Record the data for all users, unaccessable to outside
var users = make(map[string]*UserInfo)

var userfilename = "entity/users.txt"

//LoadUsers : Load the data of users
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

//SaveUsers : Save the data of users
func SaveUsers() {
	userJSON, _ := json.Marshal(users)
	err := ioutil.WriteFile(userfilename, userJSON, 0644)
	if err != nil {
		panic(err)
	}
}

//GetUserInfo : Return a copy of the userinfo of the given username
func GetUserInfo(username string) (*UserInfo, bool) {
	user, exist := users[username]
	return user, exist
}

//CreateUser : Create a new user via register
func CreateUser(name, pass, phone, email string) {
	users[name] = &UserInfo{Password: pass, Phone: phone, Email: email}
}

//GetUserHostMeetings : Get a list of meetings held by user
func GetUserHostMeetings(username string) []string {
	return users[username].HostMeetings
}

//GetUserParMeetings : Get a list of meetings participated by user
func GetUserParMeetings(username string) []string {
	return users[username].ParMeetings
}

//AddUserMeetingHost :
func AddUserMeetingHost(username, title string) {
	users[username].HostMeetings = append(users[username].HostMeetings, title)

}

//AddUserMeetingParc :
func AddUserMeetingParc(username, title string) {
	users[username].ParMeetings = append(users[username].ParMeetings, title)
}
