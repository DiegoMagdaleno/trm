package lib

import (
	"log"
	"os/user"
)

func getUserHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir, err
}

func GetTrash() string {
	homeDir, err := getUserHome()
	if err != nil {
		log.Fatal(err)
	}

	return homeDir + "/.local/share/Trash/files"
}
