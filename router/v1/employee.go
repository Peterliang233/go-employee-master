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
	c.JSON(http.StatusOK, gin.H{
		"404" : "not found",
	})
}

func AddEmployee(c *gin.Context){  //add the new information
	var newEmployee model.WorkMan
	err := c.BindJSON(&newEmployee)
	if err != nil{
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"employee" : newEmployee,
	})
	model.Worker=append(model.Worker,newEmployee)
	model.Num++
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
			"employee": newEmployee,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"404" : "not found",
	})
}

func DeleteEmployee(c *gin.Context){   //delete the information
	number := c.Query("number")
	for i := 0; i < model.Num; i++ {
		if model.Worker[i].Number == number {
			//fmt.Println(cap(employee.Worker),len(employee.Worker))
			c.JSON(http.StatusOK,gin.H{
				"employee": model.Worker[i],
			})
			if i < len(model.Worker) - 1 {
				model.Worker = append(model.Worker[:i], model.Worker[i+1:]...)
			}else{
				model.Worker= model.Worker[:len(model.Worker)-1]
			}
			model.Num--
			//hello.CachedFile()
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"404" : "not found",
	})
}