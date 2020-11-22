package v1

import (
	"github.com/Peterliang233/Function/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func FindEmployee(c *gin.Context){  //find the information
	number := c.Query("number")
	for i := 0; i < model.Num; i++ {
		if model.Worker[i].Number == number {
			c.JSON(http.StatusOK, gin.H{
				"employee": model.Worker[i],
			})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"code" : "E00001",
		"data": map[string]interface{}{},
		"message": "employee not found",
	})
}

func AddEmployee(c *gin.Context){  //add the new information
	var newEmployee model.WorkMan
	err := c.BindJSON(&newEmployee)
	if err != nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "999999",
		"employee" : newEmployee,
		"message": "ok",
	})
	defer newEmployee.AddEmployee()
	//hello.CachedFile()
}

func UpdateEmployee(c *gin.Context){   //update the information
	var newEmployee model.WorkMan
	err := c.BindJSON(&newEmployee)
	if err != nil{
		log.Fatal(err)
	}
	for i := 0; i < model.Num; i++ {
		if model.Worker[i].Number == newEmployee.Number {
			model.Worker[i] = newEmployee
		}
		c.JSON(http.StatusOK, gin.H{
			"code": "999999",
			"employee": newEmployee,
			"message": "ok",
		})
		return
	}
	c.JSON(http.StatusNotFound, gin.H{
		"code" : "E000001",
		"data": map[string]interface{}{},
		"message": "not find employee",
	})
}

func DeleteEmployee(c *gin.Context){   //delete the information
	number := c.Query("number")
	for i := 0; i < model.Num; i++ {
		if model.Worker[i].Number == number {
			c.JSON(http.StatusOK,gin.H{
				"code": "999999",
				"employee": model.Worker[i],
				"message": "ok",
			})
			model.DeleteEmployee(i)
			//hello.CachedFile()
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{
		"code" : "E000001",
		"data": map[string]interface{}{},
		"message": "not find employee",
	})
}