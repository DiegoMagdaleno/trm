package lib

import (
	"fmt"
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
	fileName, extension := getLastItem(file)
	filePathInTrash := fmt.Sprintf("%s/%s", getTrash(), fileName)
	if extension != "" {
		filePathInTrash = fmt.Sprintf("%s/%s.%s", getTrash(), fileName, extension)
	}
	if fileExist(filePathInTrash) {
		filePathInTrash = fmt.Sprintf("%s/%s-%s", getTrash(), fileName, getCurrentDate())
		if extension != "" {
			filePathInTrash = fmt.Sprintf("%s/%s-%s.%s", getTrash(), fileName, getCurrentDate(), extension)
		}
	}
	err := os.Rename(file, filePathInTrash)
	if err != nil {
		log.Fatal(err)
	}

}
