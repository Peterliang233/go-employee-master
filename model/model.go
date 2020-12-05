package model

import (
	"github.com/dgrijalva/jwt-go"
)

type WorkMan struct {
	Number     string //工号
	Name       string //姓名
	Profession string //职业
	Task       string //该员工的任务
}

type UserInfo struct { //用户登录时候输入用户名和密码
	Username string `json:"username"`
	Password string `json:"password"`
}

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var MySecret = []byte("登录")

//const TokenExpireDuration =time.Hour * 2  //时间设置可以优化
