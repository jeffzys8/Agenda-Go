package opfile

import(
	"io/ioutil"
)

var filename string = "curUser.txt"

func GetCurrentUser() (string, bool){
	username, err := ioutil.ReadFile(filename)
	if(err != nil){
		panic(err)
	}
	return string(username), len(username) != 0
} 