package model

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var RedisClient *redis.Client

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
func Redis() {

	// 创建 Redis 客户端连接
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host"),
		Password: viper.GetString("redis.password"), // Redis 未设置密码时为空
		DB:       0,                                 // 使用默认数据库
	})
	// 测试连接是否成功
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v", err)
		return
	}
	fmt.Println("Connected to Redis")
}
func RedisClose() {
	err := RedisClient.Close()
	if err != nil {
		fmt.Printf("Failed to close Redis connection: %v", err)
		return
	}
	fmt.Println("Redis connection closed")
}
