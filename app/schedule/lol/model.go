package lol

import "time"

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
	UpdateVer   int64     `gorm:"column:update_ver;default:NULL"`
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title       string    `gorm:"column:title;default:NULL"`
	Url         string    `gorm:"column:url;default:NULL"`
	Time        string    `gorm:"column:time;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (l *Lol) TableName() string {
	return "lol"
}
