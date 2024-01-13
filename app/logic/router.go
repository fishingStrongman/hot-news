package logic

import (
	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()
	r.GET("/refresh", Select)
	if err := r.Run(":8080"); err != nil {
		panic("gin 启动失败！")
	}
}
