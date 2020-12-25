package api

import (
	"github.com/Peterliang233/Function/database"
	"github.com/Peterliang233/Function/model"
	"github.com/Peterliang233/Function/router/v1/api/user/controller"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
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
		"msg":  "ok",
		"code": 2000,
		"data": map[string]interface{}{
			"employee":    employee,
			"roles":       User.Roles,
			"departments": User.Departments,
			"username":    username,
		},
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
		"msg":  "ok",
		"code": 2000,
		"data": map[string]interface{}{
			"username": username,
		},
	})
}

type NewEmployee struct {
	model.Employee
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	Department string `json:"department"`
}

//添加员工信息，只有管理员才能有注册帐号的功能
func AddEmployee(c *gin.Context) {
	var NewEmployee NewEmployee
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
	//username := c.Query("username")
	//password := c.Query("password")
	//employeeId := int(NewEmployee.ID)
	//role := c.Query("role")
	//department := c.Query("department")
	//if err != nil {
	//	fmt.Println("字符串转化错误")
	//	return
	//}
	//判断两者的id是否对应
	//if employeeId != int(NewEmployee.ID) {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"msg":    "参数错误",
	//		"detail": "用户employee_id与employee的id不对应",
	//		"code":   1,
	//	})
	//	return
	//}
	if controller.IdentifyAndUsername.Identify != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 0,
			"msg":  "权限不足",
		})
		return
	}
	//if err := database.DB.Create(&NewEmployee).Error; err != nil {
	//	//可能出现员工的ID重复的情况
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"code": 2,
	//		"msg":  "数据库创建失败",
	//	})
	//	return
	//}
	//查询数据库里面是否存在相同的id
	//fmt.Println(NewEmployee.ID)
	if err := database.DB.Where("id = ?", NewEmployee.ID).First(&model.Employee{}).Error; err == nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 7,
			"msg":  "用户id重复",
			"data": map[string]interface{}{
				"id": NewEmployee.ID,
			},
		})
		return
	}
	//查询数据库里面是否有已经存在这个用户名,如果不存在，则返回record not found
	if err := database.DB.Where("username = ?", NewEmployee.Username).First(&model.User{}).Error; err != nil {
		//fmt.Println(err)
		//将用户名和密码导入数据库
		PasswordHash, err := bcrypt.GenerateFromPassword([]byte(NewEmployee.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":  "加密错误",
				"code": 5,
			})
			return
		}
		employeeId := int(NewEmployee.ID)
		NewEmployee.Password = string(PasswordHash)
		//利用原生的mysql语言新建一个employee表格
		if err := database.DB.Exec("insert into employee (id, real_name, nick_name, english_name, sex,"+
			" age, address, mobile_phone, id_card) values (?, ?, ?, ?, ?, ?, ?, ?, ?);",
			NewEmployee.ID,
			NewEmployee.RealName,
			NewEmployee.NickName,
			NewEmployee.EnglishName,
			NewEmployee.Sex,
			NewEmployee.Age,
			NewEmployee.Address,
			NewEmployee.MobilePhone,
			NewEmployee.IDCard).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 2,
				"msg":  "数据库创建失败",
			})
			return
		}
		//利用原生的mysql语言往表格插入数据，填充user表格
		if err := database.DB.Exec("insert into user (username, password_hash, employee_id) values (?, ?, ?);",
			NewEmployee.Username, NewEmployee.Password, employeeId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": 2,
				"msg":  "数据库创建失败",
			})
			return
		}
		//完善user_role表格和user_department表格
		var userId, roleId, departmentId []uint64
		if err := database.DB.Table("user").Where("username = ?", NewEmployee.Username).Pluck("id", &userId).
			Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 4,
				"msg":  "数据库查询失败",
			})
			return
		}
		if err := database.DB.Table("role").Where("role_name = ?", NewEmployee.Role).Pluck("id", &roleId).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code": 4,
				"msg":  "数据库查询失败",
			})
			return
		}
		if err := database.DB.Table("department").Where("department_name = ?", NewEmployee.Department).
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
			"msg": "创建用户成功",
			"data": map[string]interface{}{
				"id":           NewEmployee.ID,
				"real_name":    NewEmployee.RealName,
				"nick_name":    NewEmployee.NickName,
				"english_name": NewEmployee.EnglishName,
				"sex":          NewEmployee.Sex,
				"age":          NewEmployee.Age,
				"address":      NewEmployee.Address,
				"mobile_phone": NewEmployee.MobilePhone,
				"id_card":      NewEmployee.IDCard,
				"username":     NewEmployee.Username,
				"role":         NewEmployee.Role,
				"department":   NewEmployee.Department,
			},
			"code": 5,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "该用户已经存在",
			"code": 3,
		})
	}
}
