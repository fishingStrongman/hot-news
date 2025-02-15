package toutiao

import "time"

// 接受百度的热搜数据
type Hot struct {
	List []ResDate `json:"data"`
}
type ResDate struct {
	//名称
	Title string `json:"Title"`
	//冷热数据
	Label string `json:"Label"`
	//热量
	Hot string `json:"HotValue"`
	//地址
	Url string `json:"Url"`
	//	类型
	LabelDesc string `json:"LabelDesc"`
}

//CREATE TABLE `toutiao` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`update_ver` bigint DEFAULT NULL,
//`title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
//`icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
//`url` varchar(8000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
//`hot` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
//`label_desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`updated_time` datetime DEFAULT NULL,
//PRIMARY KEY (`id`),
//KEY `index` (`update_ver`) USING BTREE
//) ENGINE=InnoDB AUTO_INCREMENT=13801 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

type TouTiao struct {
	ID int64 `json:"id" gorm:"id"`
	//更新的版本
	UpdateVer int64 `json:"update_ver" gorm:"update_ver"`
	//热搜标题
	Title string `json:"title" gorm:"title"`
	//冷热新闻
	Icon string `json:"icon" gorm:"icon"`
	//关键字 or url
	Url string `json:"url" gorm:"url"`
	//热度
	Hot string `json:"hot" gorm:"hot"`
	//新闻类型
	LabelDesc string `json:"LabelDesc" gorm:"label_desc"`
	//创建时间
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	//更新时间
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*TouTiao) TableName() string {
	return "toutiao"
}
