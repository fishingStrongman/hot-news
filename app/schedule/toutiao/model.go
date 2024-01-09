package toutiao

// 接受百度的热搜数据
type Hot struct {
	Data Date `json:"data"`
}
type Date struct {
	List []ResDate `json:"list"`
}
type ResDate struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Pic       string `json:"pic"`
	Hot       string `json:"hot"`
	Url       string `json:"url"`
	MobileUrl string `json:"mobileUrl"`
}

type TouTiao struct {
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
	CreatedTime string `json:"created_time" gorm:"created_time"`
	//更新时间
	UpdatedTime string `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*TouTiao) TableName() string {
	return "toutiao"
}
