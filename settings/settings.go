package settings

import (
	"fmt"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

// Conf 定义全局Conf
var Conf = new(Config)

type Config struct {
	AppConfig   *AppConfig   `mapstructure:"app"`
	LogConfig   *LogConfig   `mapstructure:"log"`
	MySQLConfig *MySQLConfig `mapstructure:"mysql"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
}

// AppConfig 应用配置
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Mode    string `mapstructure:"mode"`
	Port    int    `mapstructure:"port"`
	Version string `mapstructure:"version"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

// MySQLConfig MySQL 数据库配置
type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int    `mapstructure:"pool_size"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yml")
	// 设置配置文件名称（不带扩展名）
	//viper.SetConfigName("config")
	// 设置配置文件路径（可以是相对路径或绝对路径）
	viper.AddConfigPath(".")
	// 设置配置文件类型
	//viper.SetConfigType("yaml")

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("Fatal error config file: %w \n", err)
		return err
	}

	// 将配置映射到结构体
	err = viper.Unmarshal(&Conf)
	if err != nil {
		fmt.Println("failed to unmarshal config: %w", err)
		return err
	}

	// 动态加载配置信息——修改后自动生效（可注释关闭）
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		err := viper.Unmarshal(&Conf) // 更新全局 Conf
		if err != nil {
			fmt.Printf("failed to reload config: %v\n", err)
		}
	})

	return nil
}
