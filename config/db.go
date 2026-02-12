package config

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-web/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func InitDb() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Cfg.MySQL.User,
		Cfg.MySQL.Pass,
		Cfg.MySQL.Host,
		Cfg.MySQL.Port,
		Cfg.MySQL.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}
	sqlDB, err := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了可以重新使用连接的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	//将db赋给全局global.DB
	global.DB = db
	fmt.Println("✅ Mysql 数据库连接成功！")
}

func InitRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", Cfg.Redis.Host, Cfg.Redis.Port),
		Password: Cfg.Redis.Pass, // 空字符串 = 无密码
		DB:       Cfg.Redis.DB,
	})
	global.RDB = rdb
	fmt.Println("✅ Redis 连接成功！")
}
