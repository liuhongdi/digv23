package main

import (
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv23/global"
	"github.com/liuhongdi/digv23/router"
	"io"
	"log"
	"os"
)

func init() {
	//mysql
	err := global.SetupDBLink()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
	//redis
	err = global.SetupRedisDb()
	if err != nil {
		log.Fatalf("init.SetupRedisDb err: %v", err)
	}
}

func main() {
	gin.SetMode("release")
	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	//引入路由
	r := router.Router()
	//run
	r.Run(":8080")
}




