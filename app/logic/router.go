package logic

import (
	"github.com/gin-gonic/gin"
	"hotinfo/app/schedule/baidu"
	"hotinfo/app/schedule/bilibili"
	"hotinfo/app/schedule/bilibili_rank"
	"hotinfo/app/schedule/douyin"
	"hotinfo/app/schedule/hyper"
	"hotinfo/app/schedule/juejin"
	"hotinfo/app/schedule/lol"
	"hotinfo/app/schedule/tengxun"
	"hotinfo/app/schedule/toutiao"
	"hotinfo/app/schedule/wangyi"
	"hotinfo/app/schedule/weibo"
	"hotinfo/app/schedule/zhihu"
)

func Router() {
	r := gin.Default()
	manual := r.Group("/manual")
	{
		manual.GET("/bilibili", bilibili.Do)
		manual.GET("/baidu", baidu.Do)
		manual.GET("/bilibili_rank", bilibili_rank.Do)
		manual.GET("/douyin", douyin.Do)
		manual.GET("/hyper", hyper.Do)
		manual.GET("/juejin", juejin.Do)
		manual.GET("/lol", lol.Do)
		manual.GET("/tengxun", tengxun.Do)
		manual.GET("/toutiao", toutiao.Do)
		manual.GET("/wangyi", wangyi.Do)
		manual.GET("/weibo", weibo.Do)
		manual.GET("/zhihu", zhihu.Do)
	}

	if err := r.Run(":8080"); err != nil {
		panic("gin 启动失败！")
	}
}
