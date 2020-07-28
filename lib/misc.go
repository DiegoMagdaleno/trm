package lib

import (
	"strings"
	"time"
)

func getLastItem(path string) string {
	splitString := strings.Split(path, "/")
	return splitString[len(splitString)-1]
}

func getCurrentDate() string {
	dn := time.Now()
	return dn.Format("01-02-2006 15:04:05")
}
