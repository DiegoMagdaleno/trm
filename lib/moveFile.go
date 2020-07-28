package lib

import (
	"log"
	"os"
)

func fileExist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}

}

func MoveFile(file string) {
	fileName := getLastItem(file)
	filePathInTrash := getTrash() + "/" + fileName
	if fileExist(getTrash() + "/" + fileName) {
		filePathInTrash = filePathInTrash + "-" + getCurrentDate()
	}
	err := os.Rename(file, filePathInTrash)
	if err != nil {
		log.Fatal(err)
	}

}
