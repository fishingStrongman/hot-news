package logic

import (
	"github.com/gin-gonic/gin"
	"hotinfo/app/middleware"
)

func Router() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/refresh", Select)
	if err := r.Run(":8080"); err != nil {
		panic("gin 启动失败！")
	}
}
