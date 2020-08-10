package log

import (
	"fmt"

	"tk.com/util/datetime"
)

var (
	LOGGER = map[string]int{
		"Exception": 1,
		"Error":     2,
		"Warning":   3,
		"Info":      4,
		"Debug":     5,
	}
)

func Println(level, logType string, arg ...interface{}) {
	l, ok := LOGGER[level]
	if !ok {
		l = 5
	}
	lt, ok := LOGGER[logType]
	if !ok {
		return
	}
	if lt <= l {
		ts, _ := datetime.Get("", "", "Africa/Harare")
		fmt.Print("[" + logType + "] : " + ts + " -> ")
		fmt.Print(arg)
		fmt.Print("\n")
	}
}
