package lib

import (
	"strings"
	"time"
)

func getLastItem(path string) (string, string) {
	splitString := strings.Split(path, "/")
	finalElementInPath := splitString[len(splitString)-1]
	if strings.Contains(finalElementInPath, ".") {
		fileWithExtension := strings.Split(finalElementInPath, ".")
		return fileWithExtension[0], fileWithExtension[1]
	}
	return finalElementInPath, ""
}

func getCurrentDate() string {
	dn := time.Now()
	return dn.Format("01-02-2006 15:04:05")
}
