package juejin

import "time"

type Content struct {
	ContentID  string   `json:"content_id"`
	ItemType   int      `json:"item_type"`
	Format     string   `json:"format"`
	AuthorID   string   `json:"author_id"`
	Title      string   `json:"title"`
	Brief      string   `json:"brief"`
	Status     int      `json:"status"`
	CTime      int64    `json:"ctime"`
	MTime      int64    `json:"mtime"`
	CategoryID string   `json:"category_id"`
	TagIDs     []string `json:"tag_ids"`
}

type ContentCounter struct {
	View          int   `json:"view"`
	Like          int   `json:"like"`
	Collect       int   `json:"collect"`
	HotRank       int64 `json:"hot_rank"`
	CommentCount  int   `json:"comment_count"`
	InteractCount int   `json:"interact_count"`
}

type Author struct {
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	IsFollowed bool   `json:"is_followed"`
}

type AuthorCounter struct {
	Level    int `json:"level"`
	Power    int `json:"power"`
	Follower int `json:"follower"`
	Followee int `json:"followee"`
	Publish  int `json:"publish"`
	View     int `json:"view"`
	Like     int `json:"like"`
	HotRank  int `json:"hot_rank"`
}

type UserInteract struct {
	IsUserLike    bool `json:"is_user_like"`
	IsUserCollect bool `json:"is_user_collect"`
	IsFollow      bool `json:"is_follow"`
}

type Response struct {
	ErrNo  int           `json:"err_no"`
	ErrMsg string        `json:"err_msg"`
	Data   []HotListItem `json:"data"`
}

type HotListItem struct {
	Content        Content        `json:"content"`
	ContentCounter ContentCounter `json:"content_counter"`
	Author         Author         `json:"author"`
	AuthorCounter  AuthorCounter  `json:"author_counter"`
	UserInteract   UserInteract   `json:"user_interact"`
}

// 建表语句
// CREATE TABLE `juejin` (
// `id` bigint NOT NULL AUTO_INCREMENT,
// `hot` bigint DEFAULT NULL,
// `update_ver` bigint DEFAULT NULL,
// `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `author_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `author_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
// `created_time` datetime DEFAULT NULL,
// `updated_time` datetime DEFAULT NULL,
// PRIMARY KEY (`id`),
// KEY `index` (`update_ver`) USING BTREE
// ) ENGINE=InnoDB AUTO_INCREMENT=11102 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
type Juejin struct {
	ID          int64     `json:"id" gorm:"id"`
	Hot         int64     `json:"hot" gorm:"gorm"`
	UpdateVer   int64     `json:"update_ver" gorm:"update_ver"`
	Title       string    `json:"title" gorm:"title"`
	Url         string    `json:"url" gorm:"url"`
	AuthorId    string    `json:"author_id" gorm:"author_id"`
	AuthorName  string    `json:"author_name" gorm:"author_name"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (Juejin) TableName() string {
	return "juejin"
}
