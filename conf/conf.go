package conf

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server struct {
		Mode          string // debug | release
		Port          string
		DbAutoMigrate bool   // 是否自动迁移数据库表结构
		DbLogMode     string // silent | error | warn | info
	}
	Log struct {
		Level     string // debug | info | warn | error
		Prefix    string
		Format    string // text | json
		Directory string
	}
	JWT struct {
		Secret string
		Expire int64 // hour
		Issuer string
	}
	Mysql struct {
		Host     string // 服务器地址
		Port     string // 端口
		Config   string // 高级配置
		Dbname   string // 数据库名
		Username string // 数据库用户名
		Password string // 数据库密码
	}
	Redis struct {
		DB       int    // 指定 Redis 数据库
		Addr     string // 服务器地址:端口
		Password string // 密码
	}
}

var Conf *Config

// ReadConfig 读取配置文件
func ReadConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		panic("配置文件读取失败：" + err.Error())
	}
	if err := v.Unmarshal(&Conf); err != nil {
		panic("配置文件反序列化失败：" + err.Error())
	}
	log.Println("配置文件加载成功")
	return Conf
}
