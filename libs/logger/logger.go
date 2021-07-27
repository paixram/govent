package Logger

import (
	"errors"
	"fmt"
)

var Reset = "\033[0m"
var Red = "\033[31m"
var Green = "\033[32m"
var Yellow = "\033[33m"
var Blue = "\033[34m"
var Purple = "\033[35m"
var Cyan = "\033[36m"
var Gray = "\033[37m"
var White = "\033[97m"

type LOGINFO string

const (
	OK    LOGINFO = "OK"
	WARN  LOGINFO = "WARN"
	ERROR LOGINFO = "ERROR"
)

type LogInfo struct {
	Title string
	Log   map[string]LOGINFO
}

func (a *LogInfo) Logging() error {
	if a.Log == nil {
		return errors.New("The map of loger is not initialized")
	}

	fmt.Printf(Green+"******************* %s logs *******************"+Reset+"\n", a.Title)
	for i, d := range a.Log {

		switch d {
		case OK:
			lpInfo := fmt.Sprintf(Green+"[%v]: "+Reset+"[%s]"+"\n", d, i)
			fmt.Println(lpInfo)

		case WARN:
			lpInfo := fmt.Sprintf(Yellow+"[%v]: "+Reset+"[%s]"+"\n", d, i)
			fmt.Println(lpInfo)

		case ERROR:
			lpInfo := fmt.Sprintf(Red+"[%v]: "+Reset+"[%s]"+"\n", d, i)
			fmt.Println(lpInfo)

		default:
			lpInfo := fmt.Sprintf("Error in Log")
			fmt.Println(lpInfo)
		}

	}

	return nil
}
