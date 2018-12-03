package main

import (
	"fmt"
	"gin-xorm/controller"
	_ "gin-xorm/db"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	appController := new(controller.MemberAppController)
	get := engine.GET("/app/:id", appController.GetById)
	fmt.Print(get)

	engine.Run(":9204")
}
