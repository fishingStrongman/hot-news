package wangyi

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type T struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		List []struct {
			SkipType   string      `json:"skipType"`
			Title      string      `json:"title"`
			Ptime      string      `json:"ptime"`
			Source     string      `json:"source"`
			Score      float64     `json:"score,omitempty"`
			Docid      string      `json:"docid"`
			SkipID     string      `json:"skipID"`
			Keyword    string      `json:"_keyword,omitempty"`
			Origin     string      `json:"_origin,omitempty"`
			Coreword   interface{} `json:"_coreword"`
			Type       string      `json:"_type,omitempty"`
			PortOrigin string      `json:"portOrigin"`
			Url        string      `json:"url"`
			Style      string      `json:"style"`
			PicInfo    []struct {
				Url string `json:"url,omitempty"`
			} `json:"picInfo"`
			PublishTime string `json:"publishTime,omitempty"`
			CreateTime  string `json:"createTime,omitempty"`
			Imgsrc      string `json:"imgsrc,omitempty"`
			RecImgsrc   string `json:"recImgsrc,omitempty"`
			Tag         string `json:"tag,omitempty"`
		} `json:"list"`
	} `json:"data"`
}
type WangYi struct {
	ID          int64     `json:"id" gorm:"id"`
	UpdateVer   int64     `json:"update_ver" gorm:"update_ver"`
	Title       string    `json:"title" gorm:"title"`
	KeyWord     string    `json:"key_word" gorm:"key_word"`
	Url         string    `json:"url" gorm:"url"`
	CreatedTime time.Time `json:"created_time" gorm:"created_time"`
	UpdatedTime time.Time `json:"updated_time" gorm:"updated_time"`
}

// TableName 表名称
func (WangYi) TableName() string {
	return "wangyi"
}

func Createtable() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "123456", "localhost:3306", "hotinfo")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// AutoMigrate 会自动创建表，如果表已经存在则会更新字段
	err = db.AutoMigrate(&WangYi{})
	if err != nil {
		panic("Failed to migrate table")
	}
}
