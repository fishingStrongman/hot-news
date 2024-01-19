package hyper

import "time"

type Ret struct {
	ResultCode int
	ResultMsg  string
	SystemTime int64
	Data       Data
}
type Data struct {
	HotNews []*HotNews `json:"hotNews"`
}
type HotNews struct {
	ContId  string     `json:"contId"`
	Name    string     `json:"name"`
	TagList []*TagList `json:"tagList"`
}
type TagList struct {
	Tag string `json:"tag"`
}

//数据库建表语句

//CREATE TABLE `hyper` (
//`Id` bigint NOT NULL AUTO_INCREMENT,
//`update_ver` bigint DEFAULT NULL,
//`title` varchar(255) DEFAULT NULL,
//`key_word` varchar(255) DEFAULT NULL,
//`url` varchar(255) DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`updated_time` datetime DEFAULT NULL,
//PRIMARY KEY (`Id`),
//KEY `index` (`update_ver`) USING BTREE
//) ENGINE=InnoDB AUTO_INCREMENT=9583 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

type Hyper struct {
	ID          int64     `json:"id" gorm:"id"`
	UpdateVer   int64     `json:"update_ver" gorm:"column:update_ver"`
	Title       string    `json:"title" gorm:"title"`
	KeyWord     string    `json:"keyWord" gorm:"key_word"`
	Url         string    `json:"url" gorm:"url"`
	CreatedTime time.Time `json:"createdTime" gorm:"created_time"`
	UpdatedTime time.Time `json:"updatedTime" gorm:"updated_time"`
}

func (h *Hyper) TableName() string {
	return "hyper"
}
