package tengxun

import "time"

type AlternateWord struct {
	Word string `json:"word"`
	From string `json:"from"`
}

type TopWords struct {
	Fixed     []string        `json:"fixed"`
	Alternate []AlternateWord `json:"alternate"`
}

type HotlistItem struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Time        string `json:"time"`
	Abstract    string `json:"abstract"`
	Comments    int    `json:"comments"`
	ShareUrl    string `json:"shareUrl"`
	LikeInfo    int    `json:"likeInfo"`
	ReadCount   int    `json:"readCount"`
	Source      string `json:"source"`
	TlTitle     string `json:"tlTitle"`
	NewsSource  string `json:"NewsSource"`
	UserAddress string `json:"userAddress"`
}

type Response struct {
	Ret      int           `json:"ret"`
	TopWords TopWords      `json:"topWords"`
	TraceID  string        `json:"trace_id"`
	Type     int           `json:"type"`
	Hotlist  []HotlistItem `json:"hotlist"`
}

//数据库建表语句0

// CREATE TABLE `tengxun` (
// `id` bigint NOT NULL AUTO_INCREMENT,
// `update_ver` bigint DEFAULT NULL,
// `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `key_word` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `created_time` datetime DEFAULT NULL,
// `updated_time` datetime DEFAULT NULL,
// PRIMARY KEY (`id`),
// KEY `index` (`update_ver`) USING BTREE
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
type TengXun struct {
	ID          int64     `json:"id" gorm:"id"`
	UpdateVer   int64     `json:"update_ver" gorm:"update_ver"`
	Title       string    `json:"title" gorm:"title"`
	Time        string    `json:"time" gorm:"time"`
	Url         string    `json:"url" gorm:"url"`
	Source      string    `json:"source" gorm:"source"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (TengXun) TableName() string {
	return "tengxun"
}
