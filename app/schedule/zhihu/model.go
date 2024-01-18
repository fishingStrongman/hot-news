package zhihu

import "time"

// 接受百度的热搜数据
type Hot struct {
	Data []Date `json:"data"`
}
type Date struct {
	List        ResDate `json:"target"`
	DeTail_Text string  `json:"detail_Text"`
}
type ResDate struct {
	Title string `json:"title"`
	Url   string `json:"url"`
	Type  string `json:"type"`
}

// 返回给前端的结构体
//
//	type ResDate struct {
//		Num  int    `json:"num" `
//		Name string `json:"name"`
//		Hot  string `json:"hot"`
//		Url  string `json:"url"`
//	}
type ZhiHu struct {
	ID int64 `json:"id" gorm:"id"`
	//更新的版本
	UpdateVer int64 `json:"update_ver" gorm:"update_ver"`
	//热搜标题
	Title string `json:"title" gorm:"title"`
	//关键字 or url
	Url string `json:"url" gorm:"url"`
	//热度
	Hot string `json:"hot" gorm:"hot"`
	//创建时间
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	//更新时间
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*ZhiHu) TableName() string {
	return "zhihu"
}
