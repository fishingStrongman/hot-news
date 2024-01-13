package app

import (
	"hotinfo/app/logic"
	"hotinfo/app/model"
	"hotinfo/app/schedule"
)

func Start() {
<<<<<<<<< Temporary merge branch 1
	model.NewMySql()
	TaskRun()
=========
	model.NewMysql()
	//baidu.Do()
	//爬虫定时器启动
	//	taskRun()

	//服务器必须最后启动
	logic.Router()
>>>>>>>>> Temporary merge branch 2
}

func taskRun() {
	schedule.Run()
}
