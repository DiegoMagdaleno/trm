package lib

import (
	"log"
	"os"
)

func MoveFile(file string) {
	fileName := GetLastItem(file)
	err := os.Rename(file, GetTrash() + "/" + fileName) 
	if err != nil{
		log.Fatal(err)
	}

}
