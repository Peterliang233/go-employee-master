package model

import (
	"errors"
	"fmt"
	"github.com/Peterliang233/Function/database"
	"github.com/Peterliang233/Function/settings"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

const (
	UserNotFound           = 10001
	UserFailedFind         = 10002
	UserCheckPasswordError = 10003
)

//检查登录时候的密码是否正确
func (user *User) CheckoutPassword() (int, error) {
	var passwordHashPair []string
	if err := database.DB.Table("user").Where("username = ?", &user.Username).
		Pluck("password_hash", &passwordHashPair).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return UserNotFound, err
		}
		return UserFailedFind, err
	}
	if len(passwordHashPair) == 0 {
		return UserNotFound, errors.New("false")
	}
	//对密码进行解密操作
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashPair[0]), []byte(user.PasswordHash))
	if err != nil {
		return UserCheckPasswordError, err
	}
	return 0, nil
}

//获取用户的角色
func (user *User) GetUserRoles() error {
	err := database.DB.Table("user u").
		Joins("LEFT JOIN user_role ur ON u.id = ur.user_id").
		Joins("LEFT JOIN role r ON r.id = ur.role_id").
		Where("u.username = ?", &user.Username).
		Pluck("r.role_name", &user.Roles).Error
	if err != nil {
		return err
	}
	return nil
}

//获取用户的身份
func (user *User) GetUserDepartments() error {
	err := database.DB.Table("user u").
		Joins("LEFT JOIN user_department ud ON u.id = ud.user_id").
		Joins("LEFT JOIN department d ON d.id = ud.department_id").
		Where("u.username = ?", &user.Username).
		Pluck("d.department_name", &user.Departments).Error
	if err != nil {
		return err
	}
	return nil
}

func ConnectMysql() { //连接数据库
	var err error
	database.DB, err = gorm.Open(settings.DatabaseString.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		settings.DatabaseString.User,
		settings.DatabaseString.Password,
		settings.DatabaseString.Host,
		settings.DatabaseString.Dbname,
	))
	database.DB.SingularTable(true)
	if err != nil {
		panic(err)
	}
}

func CloseDatabase() {
	database.DB.Close()
}
