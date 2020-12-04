package controller

import (
	"errors"
	"fmt"
	"github.com/Peterliang233/Function/model"
	"github.com/Peterliang233/Function/settings"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AuthHandler(c *gin.Context) {
	// 用户发送用户名和密码过来
	var user model.UserInfo
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	// 校验用户名和密码是否正确
	//fmt.Println(user.Username,user.Password)
	if user.Username == settings.UserString.Username && user.Password == settings.UserString.Password {
		// 生成Token
		tokenString, _ := GenToken(user.Username)
		fmt.Println(tokenString)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return
}

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := model.MyClaims{
		Username: username, // 自定义字段
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(model.TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "my-project",                               // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(model.MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return model.MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}