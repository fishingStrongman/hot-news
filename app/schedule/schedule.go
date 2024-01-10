package schedule

import (
	"hotinfo/app/schedule/bilibili"
	"hotinfo/app/schedule/juejin"
	"hotinfo/app/schedule/tengxun"
	"hotinfo/app/schedule/wangyi"
	"hotinfo/app/schedule/weibo"
)

func Run() {
	bilibili.Run()
	juejin.Run()
	tengxun.Run()
	wangyi.Run()
	weibo.Run()
}
