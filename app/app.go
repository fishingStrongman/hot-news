package app

import (
	"hotinfo/app/logic"
	"hotinfo/app/model"
	"hotinfo/app/schedule"
	"hotinfo/app/tools"
)

func Start() {
	tools.InitFile("app/log/", "")
	model.NewMySql()
	model.Redis()
	//爬虫定时器启动
	taskRun()

	//服务器必须最后启动
	logic.Router()
	defer func() {
		model.RedisClose()
		model.Close()
	}()
}

func taskRun() {
	schedule.Run()
}
