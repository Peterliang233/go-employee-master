package database

import (
	"bufio"
	"github.com/Peterliang233/Function/model"
	"github.com/Peterliang233/Function/settings"
	"log"
	"os"
)

func ReadFile(){
	model.Num = 0
	//file := "database/information.txt"
	file := settings.DatabaseString.FilePath
	inputString, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer func(){
		if err = inputString.Close(); err != nil{
			log.Fatal(err)
		}
	}()
	scan := bufio.NewScanner(inputString)
	i := 0
	var newEmployee model.WorkMan
	for scan.Scan() {
		if i == 0{
			newEmployee.Number=scan.Text()
		}
		if i == 1{
			newEmployee.Name=scan.Text()
		}
		if i == 2 {
			newEmployee.Profession=scan.Text()
		}
		if i == 3 {
			newEmployee.Task=scan.Text()
			i=-1
			model.Worker=append(model.Worker, newEmployee)
			model.Num++
		}
		i++
	}
	err = scan.Err()
	if err != nil {
		log.Fatal(err)
	}
}