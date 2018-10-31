package opfile

import (
	"bytes"
	"io/ioutil"
	"time"
)

var LogFileName string = "log.txt"

func WriteLog(str string) error {

	// add current time
	curTime := time.Now().Unix()
	str = "[" + time.Unix(curTime, 0).String() + "] " + str

	// load the origin file
	originLog, err := ioutil.ReadFile(LogFileName)
	newLog := bytes.Join([][]byte{originLog, []byte(str)}, []byte{'\n'})

	// add the string
	err = ioutil.WriteFile(LogFileName, newLog, 0644)
	return err
}
