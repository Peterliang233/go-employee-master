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

type Login struct {
	Username string
	Password string
}

var UserString = &Login{}
var cfg *ini.File

func LoadSettings() {
	var err error
	cfg, err = ini.Load("config/configs.ini")
	if err != nil {
		log.Fatalln("settings.LoadSettings, fail to parse")
	}
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseString)
	mapTo("login", UserString)
}

func mapTo(s string, i interface{}) {  //解析配置文件里面的信息
	err := cfg.Section(s).MapTo(i)
	if err != nil {
		log.Fatalln("Cfg.MapTo", s, "err", err)
	}
}