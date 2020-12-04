package database

import (
	"fmt"
	"github.com/Peterliang233/Function/model"
	_ "github.com/Peterliang233/Function/model"
	"github.com/Peterliang233/Function/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "gorm.io/gorm"
)

var Db *gorm.DB

func ConnectMysql(){  //连接数据库
	var err error
	//fmt.Printf("%s:%s@tcp(%s)/%s?" +
	//	"charset=utf8mb4&PraseTime=True&loc=Local",
	//	settings.DatabaseString.User,
	//	settings.DatabaseString.Password,
	//	settings.DatabaseString.Host,
	//	settings.DatabaseString.Dbname,
	//)
	Db, err = gorm.Open(settings.DatabaseString.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.DatabaseString.User,
		settings.DatabaseString.Password,
		settings.DatabaseString.Host,
		settings.DatabaseString.Dbname,
		))
	//Db, err = gorm.Open("mysql", "root:mysqlpassword@(localhost)/data?charset=utf8mb4&parseTime=True&loc=Local")
	Db.SingularTable(true)
	if err != nil {
		panic(err)
	}
	//defer Db.Close()
	Db.AutoMigrate(&model.WorkMan{})
}