package lib

import (
	"log"
	"os/user"
	"runtime"
)

func getUserHome() (string, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir, err
}

func getTrash() string {
	homeDir, err := getUserHome()
	if err != nil {
		log.Fatal(err)
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		return homeDir + "/.Trash"
	case "linux", "bsd":
		return homeDir + "/.local/share/Trash/files"
	default:
		log.Fatal("Not a suported operating system")
	}

	return "unknow"

}
