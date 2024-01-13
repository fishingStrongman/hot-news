package app

import (
	"hotinfo/app/model"
	"hotinfo/app/schedule"
)

func Start() {
	model.NewMySql()
	TaskRun()
}

func TaskRun() {
	schedule.Run()
}
