package app

import (
	"hotinfo/app/logic"
	"hotinfo/app/model"
	"hotinfo/app/schedule"
)

func Start() {

	model.NewMySql()
	//TaskRun()

	//baidu.Do()
	//爬虫定时器启动
	//	taskRun()

	//服务器必须最后启动
	logic.Router()
}

func taskRun() {
	schedule.Run()
}
