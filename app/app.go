package app

import (
	"hotinfo/app/model"
	"hotinfo/app/schedule"
)

func Start() {
	model.NewMysql()
	TaskRun()
}

func TaskRun() {
	schedule.Run()
}
