package schedule

import (
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

var Tasks []func()

func Run() {
	Tasks = []func(){
		baidu.Run,
		bilibili.Run,
		bilibili_rank.Run,
		douyin.Run,
		hyper.Run,
		juejin.Run,
		lol.Run,
		tengxun.Run,
		toutiao.Run,
		wangyi.Run,
		weibo.Run,
		zhihu.Run,
	}

	for _, task := range Tasks {
		task()
	}
}
