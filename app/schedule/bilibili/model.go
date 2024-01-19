package bilibili

import "time"

type Ret struct {
	Code    int
	Message string
	Ttl     int
	Data    Data
}

type Data struct {
	Trending *Trend `json:"trending"`
}

type Trend struct {
	Title   string  `json:"title"`
	TrackId int     `json:"track_id"`
	List    []*List `json:"list"`
}

type List struct {
	Keyword  string `json:"keyword"`
	ShowName string `json:"show_name"`
	Icon     string `json:"icon"`
	Uri      string `json:"uri"`
	Goto     string `json:"goto"`
}

//数据库建表语句
//
//CREATE TABLE `bilibili` (
//`id` bigint NOT NULL AUTO_INCREMENT,
//`update_ver` bigint DEFAULT NULL,
//`title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
//`icon` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
//`key_word` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
//`created_time` datetime DEFAULT NULL,
//`updated_time` datetime DEFAULT NULL,
//`hot` bigint DEFAULT NULL,
//PRIMARY KEY (`id`),
//KEY `index` (`update_ver`) USING BTREE
//) ENGINE=InnoDB AUTO_INCREMENT=5521 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

type Bilibili struct {
	ID          int64     `json:"id" gorm:"id"`
	UpdateVer   int64     `json:"update_ver" gorm:"update_ver"`
	Title       string    `json:"title" gorm:"title"`
	Icon        string    `json:"icon" gorm:"icon"`
	KeyWord     string    `json:"key_word" gorm:"key_word"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (*Bilibili) TableName() string {
	return "bilibili"
}
