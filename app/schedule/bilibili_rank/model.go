package bilibili_rank

import "time"

type Ret struct {
	Code    int
	Message string
	Ttl     int
	Data    Data
}
type Data struct {
	Note string  `json:"note"`
	List []*List `json:"list"`
}
type List struct {
	TName string `json:"tname"`
	Title string `json:"title"`
	Owner struct {
		Name string `json:"name"`
	} `json:"owner"`
	ShortLinkV2 string `json:"short_link_v2"`
}

//CREATE TABLE `bilibili_rank` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`update_ver` bigint DEFAULT NULL,
//`title` varchar(255) DEFAULT NULL,
//`tag` varchar(255) DEFAULT NULL,
//`author` varchar(255) DEFAULT NULL,
//`url` varchar(255) DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`updated_time` datetime DEFAULT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

type BRank struct {
	Id          int64     `json:"id" gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	UpdateVer   int64     `json:"update_ver" gorm:"column:update_ver;default:NULL"`
	Title       string    `json:"title" gorm:"column:title;default:NULL"`
	Tag         string    `json:"tag" gorm:"column:tag;default:NULL"`
	Author      string    `json:"author" gorm:"column:author;default:NULL"`
	Url         string    `json:"url" gorm:"column:url;default:NULL"`
	CreatedTime time.Time `json:"created_time" gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `json:"updated_time" gorm:"column:updated_time;default:NULL"`
}

func (b *BRank) TableName() string {
	return "brank"
}
