package hyper

import "time"

//数据库建表语句

//CREATE TABLE `hyper` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
//`key_word` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`updated_time` datetime DEFAULT NULL,
//PRIMARY KEY (`id`),
//KEY `index` (`update_ver`) USING BTREE
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

type Hyper struct {
	UpdateVer   int64     `gorm:"column:update_ver"`
	Title       string    `json:"title"`
	KeyWord     string    `json:"keyWord"`
	Url         string    `json:"url"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}

func (h *Hyper) TableName() string {
	return "hyper"
}
