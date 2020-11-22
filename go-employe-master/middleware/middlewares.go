package middlewares

import (
	"fmt"
	"github.com/Peterliang233/Function/model"
	_ "github.com/Peterliang233/Function/model"
	"github.com/gin-gonic/gin"
	"os"
)

func Adapter(c *gin.Context){
	inputString, err := os.OpenFile(
		"database/information.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		fmt.Println("Open error")
	}
	defer inputString.Close()
	for i := 0; i < model.Num; i++ {
		fmt.Fprintln(inputString, model.Worker[i].Number)
		fmt.Fprintln(inputString, model.Worker[i].Name)
		fmt.Fprintln(inputString, model.Worker[i].Profession)
		fmt.Fprintln(inputString, model.Worker[i].Task)
	}
	return
}
