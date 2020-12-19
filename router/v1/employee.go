package api

import (
	"fmt"
	"github.com/Peterliang233/Function/database"
	"github.com/Peterliang233/Function/model"
	"github.com/Peterliang233/Function/router/v1/api/user/controller"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
)

func FindEmployee(c *gin.Context) {
	username := c.Query("username")
	//fmt.Println(username)
	//根据登录信息判断是否具有权限
	if controller.IdentifyAndUsername.Identify == "employee" && controller.IdentifyAndUsername.Username != username {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "权限不足",
			"code": 2001,
		})
		return
	}
	var User model.User
	if err := database.DB.Where("username = ?", username).First(&User).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "数据库查询失败",
			"code": 2002,
		})
		return
	}
	//fmt.Println(User)
	var employee model.Employee
	if err := database.DB.Where("id = ?", User.EmployeeID).First(&employee).Error; err != nil {
		//fmt.Println(employee)
		//fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "数据库查询失败",
			"code": 2002,
		})
		return
	}
	err := User.GetUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取角色错误",
			"code": 2003,
		})
		return
	}
	err = User.GetUserDepartments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "获取部门错误",
			"code": 2004,
		})
		return
	}
	//经理不能查看老板的信息
	if User.Roles[0] == "boss" && controller.IdentifyAndUsername.Identify == "manager" {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "权限不足",
			"code": 2001,
		})
		return
	}
	//员工不能查看经理和老板的信息
	if controller.IdentifyAndUsername.Identify == "employee" &&
		(User.Roles[0] == "boss" || User.Roles[0] == "manager") {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "权限不足",
			"code": 2001,
		})
		return
	}
	//如果是同一级的身份，但是不是查询自身的信息
	if username != controller.IdentifyAndUsername.Username && User.Roles[0] == controller.IdentifyAndUsername.Identify {
		c.JSON(http.StatusNotFound, gin.H{
			"msg":  "权限不足",
			"code": 2001,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":         "ok",
		"code":        2000,
		"employee":    employee,
		"roles":       User.Roles,
		"departments": User.Departments,
	})
}

func DeleteEmployee(c *gin.Context) {
	username := c.Query("username")
	if controller.IdentifyAndUsername.Identify != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "权限不足",
			"code": 2001,
		})
		return
	}
	//在数据库里面执行删除操作
	//admin可以删除非管理员的信息
	var user model.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		//fmt.Println(user)
		//fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "用户名不存在",
			"code": 2002,
		})
		return
	}
	err := user.GetUserRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":    "权限获取失败",
			"detail": "get access error",
			"code":   2003,
		})
		return
	}
	if user.Roles[0] == "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "权限不足",
			"code": 2004,
		})
		return
	}
	//正式删除该用户信息
	//获取用户id
	var id []uint64
	if err := database.DB.Table("user").Where("username = ?", username).Pluck("id", &id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "查询失败",
			"code": 2005,
		})
		return
	}
	//删除user_role表中的关联，删除user_department表中的关联
	if err := database.DB.Exec("delete from user_role where user_id = ?;", id[0]).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "查询失败",
			"code": 2005,
		})
		return
	}
	if err := database.DB.Exec("delete from user_department where user_id = ?;", id[0]).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "查询失败",
			"code": 2005,
		})
		return
	}
	if err := database.DB.Exec("delete from user where username = ?;", username).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "查询失败",
			"code": 2005,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":      "ok",
		"username": username,
		"code":     2000,
	})
}

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
				"code":     5,
				"msg":      "ok",
				"employee": employee,
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
			if err := database.DB.Model(&employee).Where("id = ?", employee.ID).Save(&employee).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code": 4,
					"msg":  "数据库操作失败",
				})
				return
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code":     5,
					"msg":      "ok",
					"employee": employee,
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
					"code":     5,
					"msg":      "ok",
					"employee": employee,
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
					"code":     5,
					"msg":      "ok",
					"employee": employee,
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

//添加员工信息，类似与注册帐号的功能
func AddEmployee(c *gin.Context) {
	var NewEmployee model.Employee
	err := c.ShouldBind(&NewEmployee)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":    "参数错误",
			"detail": "err",
			"code":   1,
		})
		return
	}
	//注册一个新用户，但是要求的是这个新用户的帐号，密码,员工ID
	username := c.Query("username")
	password := c.Query("password")
	employeeId, err := strconv.Atoi(c.Query("employee_id"))
	role := c.Query("role")
	department := c.Query("department")
	if err != nil {
		fmt.Println("字符串转化错误")
		return
	}
	if controller.IdentifyAndUsername.Identify != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "权限不足",
		})
		return
	}
	if err := database.DB.Create(&NewEmployee).Error; err != nil {
		//可能出现员工的ID重复的情况
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 2,
			"msg":  "数据库创建失败",
		})
		return
	}
	//查询数据库里面是否有已经存在这个用户名
	if err := database.DB.Where("username = ?", username).First(&model.User{}).Error; err != nil {
		//将用户名和密码导入数据库
		PasswordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "加密错误",
				"code": 5,
			})
			return
		}
		employeeId = int(NewEmployee.ID)
		password = string(PasswordHash)
		//利用原生的mysql语言往表格插入数据，填充user表格
		if err := database.DB.Exec("insert into user (username, password_hash, employee_id) values (?, ?, ?);",
			username, password, employeeId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 2,
				"msg":  "数据库创建失败",
			})
			return
		}
		//完善user_role表格和user_department表格
		var userId, roleId, departmentId []uint64
		if err := database.DB.Table("user").Where("username = ?", username).Pluck("id", &userId).
			Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 4,
				"msg":  "数据库查询失败",
			})
			return
		}
		if err := database.DB.Table("role").Where("role_name = ?", role).Pluck("id", &roleId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 4,
				"msg":  "数据库查询失败",
			})
			return
		}
		if err := database.DB.Table("department").Where("department_name = ?", department).
			Pluck("id", &departmentId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 4,
				"msg":  "数据库查询失败",
			})
			return
		}
		err = database.DB.Exec("insert into user_role (user_id, role_id) values (?,?);", userId[0], roleId[0]).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 2,
				"msg":  "数据库创建失败",
			})
			return
		}
		err = database.DB.Exec("insert into user_department (user_id, department_id) values (?,?);", userId[0], departmentId[0]).Error
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 2,
				"msg":  "数据库创建失败",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":      "创建用户成功",
			"username": username,
			"code":     5,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "该用户已经存在",
			"code": 3,
		})
	}
}
