package weibo

import "time"

type Response struct {
	Ok   int `json:"ok"`
	Data struct {
		HotGov   HotGov     `json:"hotgov"`
		Realtime []Realtime `json:"realtime"`
	} `json:"data"`
}
type HotGov struct {
	// 定义 hotgov 结构体的字段
}
type Realtime struct {
	Flag     int    `json:"flag"`
	RealPos  int    `json:"realpos"`
	RawHot   int    `json:"raw_hot"`
	Note     string `json:"note"`
	IconDesc string `json:"icon_desc"`
	Category string `json:"category"`
}

// CREATE TABLE `weibo` (
// `id` bigint NOT NULL AUTO_INCREMENT,
// `update_ver` bigint DEFAULT NULL,
// `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `icon_desc` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `category` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `created_time` datetime DEFAULT NULL,
// `updated_time` datetime DEFAULT NULL,
// `hot` bigint DEFAULT NULL,
// PRIMARY KEY (`id`),
// KEY `index` (`update_ver`) USING BTREE
// ) ENGINE=InnoDB AUTO_INCREMENT=14997 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
type WeiBo struct {
	ID          int64     `json:"id" gorm:"id"`
	UpdateVer   int64     `json:"update_ver" gorm:"update_ver"`
	Title       string    `json:"title" gorm:"title"`
	Hot         int       `json:"hot"`
	Url         string    `json:"url" gorm:"url"`
	IconDesc    string    `json:"icon_desc" gorm:"icon_desc"`
	Category    string    `json:"category" gorm:"category"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (WeiBo) TableName() string {
	return "weibo"
}
