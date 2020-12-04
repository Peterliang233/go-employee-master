package v1

import (
	"github.com/Peterliang233/Function/database"
	"github.com/Peterliang233/Function/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindEmployee(c *gin.Context){  //查找员工信息
	number := c.Query("number")
	var newEmployee model.WorkMan
	//newEmployee.Number = number
	if err:=database.Db.Where("number = ?", number).First(&newEmployee).Error;err != nil{
		//未查找到
		c.JSON(http.StatusNotFound, gin.H{
			"code" : "E00001",
			"data": map[string]interface{}{},
			"message": "employee not found",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code" : "999999",
			"employee" : newEmployee,
		})
	}
}

func AddEmployee(c *gin.Context){  //添加一个员工的信息
	var newEmployee model.WorkMan
	err := c.BindJSON(&newEmployee)
	if err != nil{
		log.Fatal(err)
	}
	//对新添加的信息进行导入数据库
	//database.Db.AutoMigrate(&model.WorkMan{})
	if err := database.Db.Create(&newEmployee).Error;err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : "E000001",
			"data": map[string]interface{}{},
			"message": "not find employee",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": "999999",
			"employee" : newEmployee,
			"message": "ok",
		})
	}
	//hello.CachedFile()
}

func UpdateEmployee(c *gin.Context){   //更新员工的信息
	var newEmployee model.WorkMan
	err := c.BindJSON(&newEmployee)   //通过post请求获取修改的信息
	if err != nil{
		log.Fatal(err)
	}
	//对数据库里面的对应的信息进行修改
	//database.Db.AutoMigrate(&model.WorkMan{})
	if err := database.Db.Model(&newEmployee).Where("Number=?",newEmployee.Number).Update(map[string]interface{}{
		"Name" : newEmployee.Name,
		"Profession" : newEmployee.Profession,
		"Task" : newEmployee.Task,
	}).Error;err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code" : "E000001",
			"data": map[string]interface{}{},
			"message": "not find employee",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"code": "999999",
			"employee": newEmployee,
			"message": "ok",
		})
	}
	//for i := 0; i < model.Num; i++ {
	//	if model.Worker[i].Number == newEmployee.Number {
	//		model.Worker[i] = newEmployee
	//	}
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": "999999",
	//		"employee": newEmployee,
	//		"message": "ok",
	//	})
	//	return
	//}
}

func DeleteEmployee(c *gin.Context){   //delete the information
	number := c.Query("number")  //获取员工的工号
	//for i := 0; i < model.Num; i++ {
	//	if model.Worker[i].Number == number {
	//		c.JSON(http.StatusOK,gin.H{
	//			"code": "999999",
	//			"employee": model.Worker[i],
	//			"message": "ok",
	//		})
	//		model.DeleteEmployee(i)
	//		//hello.CachedFile()
	//		return
	//	}
	//}
	//通过员工的工号进行删除数据
	//database.Db.AutoMigrate(&model.WorkMan{})
	var newEmployee model.WorkMan
	database.Db.Where("Number = ?", number).First(&newEmployee)
	if err := database.Db.Where("Number=?", number).Delete(&model.WorkMan{}).Error;err != nil {
		//database.Db.Delete(&newEmployee)  //删除对应的员工
		c.JSON(http.StatusNotFound, gin.H{
			"code" : "E000001",
			"data": map[string]interface{}{},
			"message": "not find employee",
		})
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code": "999999",
			"employee": newEmployee,
			"message": "ok",
		})
	}
}