package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type Server struct {
	RunMode string
	HttpMode int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct{
	Type string
	User string
	Password string
	Host string
	Dbname string
}

var DatabaseString = &Database{}

var cfg *ini.File

func LoadSettings() {
	var err error
	cfg, err = ini.Load("config/configs.ini")
	if err != nil {
		log.Fatalln("settings.LoadSettings, fail to parse")
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseString)
}

func mapTo(s string, i interface{}) {
	err := cfg.Section(s).MapTo(i)
	if err != nil {
		log.Fatalln("Cfg.MapTo", s, "err", err)
	}
}