package model

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func init() {
	// 设置配置文件名和路径
	viper.SetConfigName("config.yaml") // 配置文件名（不含扩展名）
	viper.SetConfigType("yaml")        // 配置文件类型
	viper.AddConfigPath(".")           // 配置文件所在路径

	err := viper.ReadInConfig()
	if err != nil {
		panic("配置文件读取失败")
	}
}

func NewMySql() {
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.database"),
	)
	conn, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(
			err)
	}
	Conn = conn
}
func Close() {
	db, _ := Conn.DB()
	_ = db.Close()
	return
}
