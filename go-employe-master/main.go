package main

import (
	"github.com/Peterliang233/Function/database"
	routers "github.com/Peterliang233/Function/router"
)

func main() {
	database.ReadFile() //read the history information
	r := routers.InitRouters()
	r.Run(":9090")
}