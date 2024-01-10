package schedule

import (
	"hotinfo/app/schedule/bilibili"
	"hotinfo/app/schedule/hyper"
	"hotinfo/app/schedule/lol"
)
func Run() {
	bilibili.Run()
	hyper.Run()
	lol.Run()
}
