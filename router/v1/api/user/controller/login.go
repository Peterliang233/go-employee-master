package controller

import (
	"github.com/Peterliang233/Function/middlerware"
	"github.com/Peterliang233/Function/model"
	InternalModel "github.com/Peterliang233/Function/router/v1/api/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

var IdentifyAndUsername model.IdentifyAndUsername //用来存储登录者的用户信息

func Login(c *gin.Context) {
	var loginRequestBody InternalModel.LoginRequestBody
	err := c.ShouldBind(&loginRequestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "参数错误",
			"detail": err,
			"code":   1,
		})
		return
	}
	user := model.User{
		Username:     loginRequestBody.Username,
		PasswordHash: loginRequestBody.Password,
	}
	code, err := user.CheckoutPassword()
	if err != nil {
		if code == model.UserNotFound || code == model.UserCheckPasswordError {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   code,
				"detail": user.Username + "password error or not found",
				"msg":    "用户密码错误",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":   code,
				"msg":    "login fail",
				"detail": "登录失败",
			})
			return
		}
	}
	//获取用户角色
	err = user.GetUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 3,
			"msg":  "获取权限失败",
		})
		return
	}
	token, err := middlerware.GenerateToken(user.Username, user.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 3,
			"msg":  "token 生成失败",
		})
		return
	}
	IdentifyAndUsername.Username = user.Username
	IdentifyAndUsername.Identify = user.Roles[0]
	c.JSON(http.StatusOK, gin.H{
		"code":     0,
		"msg":      "登录成功",
		"detail":   "welcome",
		"username": user.Username,
		"roles":    user.Roles,
		"token":    token,
	})
}
