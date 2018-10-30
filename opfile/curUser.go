package opfile

import (
	"io/ioutil"
)

var filename = "curUser.txt"

func GetCurrentUser() (string, bool) {
	username, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(username), len(username) != 0
}

func SetCurrentUser(username string) {
	err := ioutil.WriteFile(filename, []byte(username), 0644)
	if err != nil {
		panic(err)
	}
}
