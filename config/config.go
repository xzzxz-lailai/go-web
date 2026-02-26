package config

import (
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	MySQL  MySQLConfig  `mapstructure:"mysql"`
	Redis  RedisConfig  `mapstructure:"redis"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	SMTP   SMTPConfig   `mapstructure:"smtp"`
	Api    ApiConfig    `mapstructure:"api"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
}

type MySQLConfig struct {
	Host   string `mapstructure:"host"`
	Port   int    `mapstructure:"port"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"password"`
	DBName string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Pass string `mapstructure:"password"`
	DB   int    `mapstructure:"db"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire string `mapstructure:"expire"`
}

type SMTPConfig struct {
	Host      string
	Port      int
	User      string
	Pass      string
	From      string
	From_Name string
}

type ApiConfig struct {
	Weather_key string
}

var Cfg Config

func InitConfig() {
	content := InitNacos() // 从 Nacos 拉配置（返回 yaml 字符串）

	// 配置文件类型（扩展名）
	viper.SetConfigType("yaml")

	// 读取配置内容
	err := viper.ReadConfig(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	// 将读取到的配置解码到 Config 结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		panic(err)
	}
	InitDb()    // 初始化mysql
	InitRedis() // 初始化redis
}
