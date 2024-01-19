package douyin

import "time"

// 接受百度的热搜数据
type Hot struct {
	Data Date `json:"data"`
}
type Date struct {
	List []ResDate `json:"word_list"`
}
type ResDate struct {
	Id        string `json:"id"`
	Title     string `json:"word"`
	Pic       string `json:"pic"`
	Hot       int    `json:"hot_value"`
	MobileUrl string `json:"mobileUrl"`
}

// CREATE TABLE `douyin` (
// `id` bigint NOT NULL AUTO_INCREMENT,
// `update_ver` bigint DEFAULT NULL,
// `title` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
// `url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
// `hot` int DEFAULT NULL,
// `created_time` datetime DEFAULT NULL,
// `updated_time` datetime DEFAULT NULL,
// PRIMARY KEY (`id`),
// KEY `index` (`update_ver`) USING BTREE /*!80000 INVISIBLE */
// ) ENGINE=InnoDB AUTO_INCREMENT=10966 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
type DouYin struct {
	ID int64 `json:"id" gorm:"id"`
	//更新的版本
	UpdateVer int64 `json:"update_ver" gorm:"update_ver"`
	//热搜标题
	Title string `json:"title" gorm:"title"`
	//关键字 or url
	Url string `json:"url" gorm:"url"`
	//热度
	Hot int `json:"hot" gorm:"hot"`
	//创建时间
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	//更新时间
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*DouYin) TableName() string {
	return "douyin"
}
