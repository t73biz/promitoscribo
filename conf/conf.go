package conf

import (
	"log"
	"os"

	"github.com/Unknwon/goconfig"
)

func GetDevRoot() (path string) {
	if os.Args[1] == "" {
		log.Fatal("You must provide the working directory as an argument.")
	}
	workdir := os.Args[1]

	cfg, err := goconfig.LoadConfigFile(workdir + "conf/app.ini")
	if err != nil {
		log.Fatal(4, "Fail to parse 'config/app.ini': %v", err)
	}

	devroot, err := cfg.GetValue("app", "DEV_ROOT")
	if err != nil {
		log.Fatal("Could not find the value for DEV_ROOT")
	}

	return devroot
}
