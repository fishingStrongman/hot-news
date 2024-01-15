package lol

import "time"

type Ret struct {
	Data Data
}
type Data struct {
	Result []*Result `json:"result"`
}
type Result struct {
	SIdxTime string `json:"sIdxTime"`
	STitle   string `json:"sTitle"`
	IDocId   string `json:"iDocId"`
}

//建表语句
//CREATE TABLE `lol` (
//`update_ver` bigint DEFAULT NULL,
//`id` bigint NOT NULL AUTO_INCREMENT,
//`title` varchar(255) DEFAULT NULL,
//`url` varchar(255) DEFAULT NULL,
//`time` varchar(255) DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`updated_time` datetime DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

type Lol struct {
	UpdateVer   int64     `json:"update_ver" gorm:"column:update_ver;default:NULL"`
	Id          int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title       string    `json:"title" gorm:"column:title;default:NULL"`
	Url         string    `json:"url" gorm:"column:url;default:NULL"`
	Time        string    `json:"time" gorm:"column:time;default:NULL"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;default:NULL"`
}

func (l *Lol) TableName() string {
	return "lol"
}
