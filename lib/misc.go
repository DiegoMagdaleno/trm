package lib

import "strings"

func GetLastItem(path string) string {
	splitString := strings.Split(path, "/")
	return splitString[len(splitString)-1]
}
