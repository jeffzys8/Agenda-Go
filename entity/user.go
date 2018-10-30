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

// Users : Record the data for all users
var Users = make(map[string]UserInfo)

var filename = "entity/users.txt"

//LoadUsers : Load the data of users
func LoadUsers() {
	userJSON, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(userJSON, &Users)
	if err != nil {
		panic(err)
	}
}

//SaveUsers : Save the data of users
func SaveUsers() {
	userJSON, _ := json.Marshal(Users)
	err := ioutil.WriteFile(filename, userJSON, 0644)
	if err != nil {
		panic(err)
	}
}
