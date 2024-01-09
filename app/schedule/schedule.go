package schedule

import (
	"hotinfo/app/schedule/baidu"
	"hotinfo/app/schedule/douyin"
)

func Run() {
	baidu.Run()
	douyin.Run()
	//toutiao.Run()
	//zhihu.Run()
	//bilibili.Run()
}
