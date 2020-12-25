package api

import (
	"github.com/Peterliang233/Function/database"
	"github.com/Peterliang233/Function/model"
	"github.com/Peterliang233/Function/router/v1/api/user/controller"
	InternalModel "github.com/Peterliang233/Function/router/v1/api/user/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateEmployee(c *gin.Context) {
	//自己可以修改自己的信息
	//老板可以修改经理的信息，经理可以修改员工的信息
	var employee model.Employee
	//获取想要修改的员工的用户名
	username := c.Query("username")
	err := c.ShouldBind(&employee)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code ": 1,
			"msg":   "参数传递错误",
		})
		return
	}
	if err := database.DB.Where("username = ?", username).First(&model.User{}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "参数传递错误",
		})
		return
	}
	var user model.User
	user.Username = username
	err = user.GetUserRoles() //获取被更新者的信息
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "获取权限失败",
		})
		return
	}
	//修改的是登录用户自身的信息
	if username == controller.IdentifyAndUsername.Username {
		if err := database.DB.Model(&employee).Where("id = ?", employee.ID).Save(&employee).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 4,
				"msg":  "数据库操作失败",
			})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 5,
				"msg":  "ok",
				"data": map[string]interface{}{
					"employee": employee,
				},
			})
			return
		}
	}
	if controller.IdentifyAndUsername.Identify == "employee" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 3,
			"msg":  "权限不足",
		})
		return
	} else if controller.IdentifyAndUsername.Identify == "manager" {
		if user.Roles[0] == "employee" {
			//对数据库进行更新操作
			if err := database.DB.Model(&employee).Where("id = ?", employee.ID).Save(&employee).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 4,
					"msg":  "数据库操作失败",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 5,
					"msg":  "ok",
					"data": map[string]interface{}{
						"employee": employee,
					},
				})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 3,
				"msg":  "权限不足",
			})
			return
		}
	} else if controller.IdentifyAndUsername.Identify == "boss" {
		if user.Roles[0] == "manager" || user.Roles[0] == "employee" {
			if err := database.DB.Model(&employee).Where("id = ?", employee.ID).Save(&employee).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 4,
					"msg":  "数据库操作失败",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 5,
					"msg":  "ok",
					"data": map[string]interface{}{
						"employee": employee,
					},
				})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 3,
				"msg":  "权限不足",
			})
			return
		}
	} else if controller.IdentifyAndUsername.Identify == "admin" {
		if user.Roles[0] == "admin" {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 3,
				"msg":  "权限不足",
			})
			return
		} else {
			if err := database.DB.Model(&employee).Where("id = ?", employee.ID).Save(&employee).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 4,
					"msg":  "数据库操作失败",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 5,
					"msg":  "ok",
					"data": map[string]interface{}{
						"employee": employee,
					},
				})
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 6,
			"msg":  "用户身份错误",
		})
		return
	}
}

func UpdatePassword(c *gin.Context) {
	var loginRequestBody InternalModel.LoginRequestBody
	err := c.ShouldBind(&loginRequestBody)
	newPassword := c.Query("new_password")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "参数错误",
			"detail": err,
			"code":   1,
		})
		return
	}
	//判断是否拥有修改密码的权限
	if controller.IdentifyAndUsername.Identify != "admin" &&
		controller.IdentifyAndUsername.Username != loginRequestBody.Username {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "权限不足",
		})
		return
	}
	//检查对应的用户名和密码是否对应
	user := model.User{
		Username:     loginRequestBody.Username,
		PasswordHash: loginRequestBody.Password,
	}
	code, err := user.CheckoutPassword()
	if err != nil {
		if code == model.UserNotFound || code == model.UserCheckPasswordError {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   3,
				"detail": user.Username + "password error or not found",
				"msg":    "用户密码错误",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":   4,
				"msg":    "login fail",
				"detail": "检查失败",
			})
			return
		}
	}
	err = user.GetUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 3,
			"msg":  "获取权限失败",
		})
		return
	}
	if controller.IdentifyAndUsername.Username == loginRequestBody.Username {
		//可以修改自己的密码
		code, err := model.ChangePassword(loginRequestBody.Username, newPassword)
		if err != nil {
			if code == model.GeneratePasswordError {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 6,
					"msg":  "加密错误",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 6,
					"msg":  "数据库操作错误",
				})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg": "修改成功",
			"data": map[string]interface{}{
				"username": loginRequestBody.Username,
			},
			"code": 0,
		})
		return
	} else {
		//管理员则可以修改自身及非管理员的密码
		if controller.IdentifyAndUsername.Identify == "admin" {
			//自己可以修改自己的密码
			if user.Roles[0] == "admin" {
				c.JSON(http.StatusBadRequest, gin.H{
					"code":   7,
					"msg":    "权限不足",
					"detail": "不能修改非自身的管理员的密码",
				})
				return
			}
			code, err := model.ChangePassword(loginRequestBody.Username, newPassword)
			if err != nil {
				if code == model.GeneratePasswordError {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 6,
						"msg":  "加密错误",
					})
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{
						"code": 6,
						"msg":  "数据库操作错误",
					})
				}
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"msg": "修改成功",
				"data": map[string]interface{}{
					"username": loginRequestBody.Username,
				},
				"code": 0,
			})
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg":  "权限不足",
		"code": 4,
	})
	return
}

func UpdateRole(c *gin.Context) {
	username := c.Query("username")
	newRole := c.Query("new_role")
	if controller.IdentifyAndUsername.Identify != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "权限不足",
			"code": 1,
		})
		return
	}
	var userId []int
	if err := database.DB.Table("user").Where("username = ?", username).Pluck("id", &userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "数据库操作错误",
		})
		return
	}
	//查找对应的角色
	user := model.User{
		Username: username,
	}
	err := user.GetUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 3,
			"msg":  "获取角色失败",
		})
		return
	}
	//寻找对应的用户的角色的id
	var roleId []int
	if err := database.DB.Table("role").Where("role_name = ?", user.Roles[0]).Pluck("id", &roleId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "数据库操作错误",
		})
		return
	}
	var newRoleId []int
	if err := database.DB.Table("role").Where("role_name = ?", newRole).Pluck("id", &newRoleId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "数据库操作错误",
		})
		return
	}
	//对数据库进行更新操作
	if err := database.DB.Table("user_role").Where("user_id = ?", userId[0]).Where("role_id = ?", roleId[0]).
		Updates(map[string]interface{}{
			"user_id": userId[0],
			"role_id": newRoleId[0],
		}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "更新数据库操作失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
		"data": map[string]interface{}{
			"username": username,
			"old_role": user.Roles[0],
			"new_role": newRole,
		},
		"code": 0,
	})
}

func UpdateDepartment(c *gin.Context) {
	username := c.Query("username")
	newDepartment := c.Query("new_department")
	if controller.IdentifyAndUsername.Identify != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "权限不足",
			"code": 1,
		})
		return
	}
	var userId []int
	if err := database.DB.Table("user").Where("username = ?", username).Pluck("id", &userId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "数据库操作错误",
		})
		return
	}
	//查找对应的角色
	user := model.User{
		Username: username,
	}
	err := user.GetUserDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 3,
			"msg":  "获取部门失败",
		})
		return
	}
	//寻找对应的用户的角色的id
	var departmentId []int
	if err := database.DB.Table("department").Where("department_name = ?", user.Departments[0]).Pluck("id", &departmentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "数据库操作错误",
		})
		return
	}
	var newDepartmentId []int
	if err := database.DB.Table("department").Where("department_name = ?", newDepartment).Pluck("id", &newDepartmentId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "数据库操作错误",
		})
		return
	}
	//对数据库进行更新操作
	//fmt.Println(userId[0], departmentId[0])
	if err := database.DB.Table("user_department").Where("user_id = ?", userId[0]).
		Where("department_id = ?", departmentId[0]).Updates(map[string]interface{}{
		"user_id":       userId[0],
		"department_id": newDepartmentId[0],
	}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 2,
			"msg":  "更新数据库操作失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "修改成功",
		"data": map[string]interface{}{
			"username": username,
			"old_role": user.Departments[0],
			"new_role": newDepartment,
		},
		"code": 0,
	})
}
