package schedule

import "hotinfo/app/schedule/bilibili"

var Tasks []func()

func Run() {
	Tasks = []func(){
		bilibili.Run,
		//juejin.Run,
		//tengxun.Run,
		//wangyi.Run,
		//weibo.Run,
	}

	for _, task := range Tasks {
		task()
	}
}
