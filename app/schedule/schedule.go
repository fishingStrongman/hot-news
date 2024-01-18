package schedule

import (
	"fmt"
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

	// 设置 WaitGroup 的计数器为并发次数

	for i := 0; i < len(Tasks); i++ {
		go func(index int) {
			Tasks[index]()
			// 执行并发任务
			fmt.Printf("并发任务 %d 执行\n", index)
			// 模拟任务执行时间
			// 这里可以替换为你的实际任务逻辑
			// 例如，调用一个函数或执行一段代码
			// time.Sleep(time.Second)
			// ...
			// 完成并发任务，减少 WaitGroup 的计数器
		}(i)
	}

	// 等待所有并发任务完成

	//
	fmt.Println("所有并发任务已完成")
}
