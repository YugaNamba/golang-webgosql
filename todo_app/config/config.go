package config

import (
	"log"
	"todo_app/utils"

	"gopkg.in/ini.v1"
)

type ConfigList struct {
	Port string
	SQLDriver string
	DBName string
	LogFile string
	Static string
}


var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Fatalln("Fail to read file: ", err)
	}

	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DBName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("log").Key("logFile").String(),
		Static: cfg.Section("web").Key("static").String(),
	}
}
